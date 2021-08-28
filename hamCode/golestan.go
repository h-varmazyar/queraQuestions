//question url: https://quera.ir/contest/assignments/32393/problems/108518
package hamCode

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var persons map[string]person
var classes map[string]Class

type person struct {
	Role       string
	Id         string
	Name       string
	Year       string
	Field      string
	ClassNames []string
	Grades     map[string]string
	PGrades    []int
}

type Class struct {
	Name     string
	Id       string
	Field    string
	Students []person
	Prof     *person
	Grades   map[string]string
}

func main() {
	classes = make(map[string]Class)
	persons = make(map[string]person)
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "end" {
			break
		}
		arr := strings.Split(input, " ")
		command(arr[0], arr[1:])
	}
}

func command(command string, args []string) {
	switch strings.TrimSpace(command) {
	case "register_student":
		registerStudent(args[0], args[1], args[2], args[3])
	case "register_professor":
		registerProf(args[0], args[1], args[2])
	case "make_class":
		makeClass(args[0], args[1], args[2])
	case "student_status":
		studentStatus(args[0])
	case "professor_status":
		profStatus(args[0])
	case "class_status":
		classStatus(args[0])
	case "add_student":
		addStudent(args[0], args[1])
	case "add_professor":
		addProf(args[0], args[1])
	case "set_final_mark":
		setFinalMark(args[0], args[1], args[2], args[3])
	case "mark_student":
		markStudent(args[0], args[1])
	case "mark_list":
		markList(args[0])
	case "average_mark_professor":
		average_mark_professor(args[0])
	case "average_mark_student":
		fmt.Println(average_mark_student(args[0]))
	case "top_student":
		top_student(args[0], args[1])
	case "top_mark":
		top_mark(args[0])
	}
}

func top_mark(classId string) {
	max := float64(0)
	found := false
	class, ok := classes[classId]
	if !ok {
		fmt.Println("invalid class")
	} else {
		for _, g := range class.Grades {
			grade, _ := strconv.ParseFloat(strings.TrimSpace(g), 64)
			if grade > max {
				max = grade
				found = true
			}
		}

		if !found {
			fmt.Println("None")
		} else {
			fmt.Println(max)
		}
	}
}

func top_student(field, year string) {
	max := float64(0)
	found := false
	for _, st := range persons {
		if st.Role == "student" && st.Field == field && st.Year == year {
			found = true
			avg := average_mark_student(st.Id)
			if avg != "None" {
				grade, _ := strconv.ParseFloat(strings.TrimSpace(avg), 64)
				if grade > max {
					max = grade
				}
			}
		}
	}
	if !found {
		fmt.Println("None")
	} else {
		fmt.Println(max)
	}
}

func average_mark_student(stId string) string {
	student, ok := persons[stId]
	if !ok {
		return "invalid student"
	} else {
		if student.Grades == nil || len(student.Grades) == 0 {
			return "None"
		} else {
			total := 0
			for _, grade := range student.Grades {
				t, _ := strconv.Atoi(strings.TrimSpace(grade))
				total += t
			}
			return fmt.Sprintf("%.2f\n", float64(total/len(student.Grades)))
		}
	}
}

func average_mark_professor(pId string) {
	prof, ok := persons[pId]
	if !ok {
		fmt.Println("invalid professor")
	} else {
		if prof.PGrades == nil || len(prof.PGrades) == 0 {
			fmt.Println("None")
		} else {
			total := 0
			for _, grade := range prof.PGrades {
				total += grade
			}
			fmt.Printf("%.2f\n", float64(total/len(prof.PGrades)))
		}
	}
}

func markList(classId string) {
	err := false
	class, ok3 := classes[classId]
	if !ok3 {
		fmt.Println("invalid class")
		err = true
	}
	if class.Prof == nil {
		fmt.Println("no professor")
		err = true
	}
	if class.Students == nil || len(class.Students) == 0 {
		fmt.Println("no student")
		err = true
	}
	if !err {
		g := ""
		for _, student := range class.Students {
			mark, okM := class.Grades[student.Id]
			if okM {
				g += strings.TrimSpace(mark) + " "
			} else {
				g += "None "
			}
		}
		fmt.Println(strings.TrimSpace(g))
	}
}

func markStudent(stId, classId string) {
	err := false
	student, ok2 := persons[stId]
	if !ok2 {
		fmt.Println("invalid student")
		err = true
	}
	class, ok3 := classes[classId]
	if !ok3 {
		fmt.Println("invalid class")
		err = true
	}
	found := false
	for _, s := range class.Students {
		if s.Id == stId {
			found = true
		}
	}
	if !found {
		fmt.Println("student did not registered")
		err = true
	}
	if !err {
		val, ok := student.Grades[classId]
		if ok {
			fmt.Println(strings.TrimSpace(val))
		} else {
			fmt.Println("None")
		}
	}
}

func setFinalMark(profId, stId, classId, mark string) {
	err := false
	prof, ok := persons[profId]
	if !ok {
		fmt.Println("invalid professor")
		err = true
	}
	student, ok2 := persons[stId]
	if !ok2 {
		fmt.Println("invalid student")
		err = true
	}
	class, ok3 := classes[classId]
	if !ok3 {
		fmt.Println("invalid class")
		err = true
	}
	if class.Prof != nil && class.Prof.Id != profId {
		fmt.Println("professor class is not match")
		err = true
	}
	found := false
	for _, s := range class.Students {
		if s.Id == stId {
			found = true
		}
	}
	if !found {
		fmt.Println("student did not registered")
		err = true
	}
	if !err {
		g, _ := strconv.Atoi(strings.TrimSpace(mark))
		prof.PGrades = append(prof.PGrades, g)
		class.Grades[stId] = mark
		classes[classId] = class
		student.Grades[classId] = mark
		persons[stId] = student
		fmt.Println("student final mark added or changed")
	}
}

func registerStudent(name, id, year, field string) {
	_, ok := persons[id]
	if ok {
		fmt.Println("this identical number previously registered")
	} else {
		persons[id] = person{
			Role:  "student",
			Id:    id,
			Name:  name,
			Year:  year,
			Field: field,
		}
		fmt.Println("welcome to golestan")
	}
}

func registerProf(name, id, field string) {
	_, ok := persons[id]
	if ok {
		fmt.Println("this identical number previously registered")
	} else {
		persons[id] = person{
			Role:  "prof",
			Id:    id,
			Name:  name,
			Field: field,
		}
		fmt.Println("welcome to golestan")
	}
}

func makeClass(name, id, field string) {
	_, ok := classes[id]
	if ok {
		fmt.Println("this class id previously used")
	} else {
		classes[id] = Class{
			Id:    id,
			Name:  name,
			Field: field,
		}
		fmt.Println("class added successfully")
	}
}

func addStudent(stId, classId string) {
	student, ok := persons[stId]
	if !ok {
		fmt.Println("invalid student")
		return
	}
	class, ok2 := classes[classId]
	if !ok2 {
		fmt.Println("invalid class")
		return
	}
	if student.Field != class.Field {
		fmt.Println("student field is not match")
		return
	}
	found := false
	for _, s := range class.Students {
		if s.Id == stId {
			found = true
			fmt.Println("student is already registered")
			return
		}
	}
	if !found {
		student.ClassNames = append(student.ClassNames, class.Name)
		persons[stId] = student
		class.Students = append(class.Students, student)
		classes[classId] = class
		fmt.Println("student added successfully to the class")
	}
}

func addProf(pId, classId string) {
	prof, ok := persons[pId]
	if !ok {
		fmt.Println("invalid professor")
		return
	}
	class, ok2 := classes[classId]
	if !ok2 {
		fmt.Println("invalid class")
		return
	}
	if prof.Field != class.Field {
		fmt.Println("professor field is not match")
		return
	}
	if class.Prof != nil {
		fmt.Println("this class has a professor")
		return
	} else {
		prof.ClassNames = append(prof.ClassNames, class.Name)
		persons[pId] = prof
		class.Prof = &prof
		fmt.Println("professor added successfully to the class")
	}
}

func studentStatus(id string) {
	student, ok := persons[id]
	if !ok {
		fmt.Println("invalid student")
	} else {
		cs := strings.TrimSuffix(strings.TrimPrefix(fmt.Sprint(student.ClassNames), "["), "]")
		fmt.Printf("%s %s %s %s\n", student.Name, student.Year, student.Field, cs)
	}
}

func profStatus(id string) {
	prof, ok := persons[id]
	if !ok {
		fmt.Println("invalid professor")
	} else {
		cs := strings.TrimSuffix(strings.TrimPrefix(fmt.Sprint(prof.ClassNames), "["), "]")
		fmt.Printf("%s %s %s\n", prof.Name, prof.Field, cs)
	}
}

func classStatus(id string) {
	class, ok := classes[id]
	if !ok {
		fmt.Println("invalid class")
	} else {
		profName := "None"
		if class.Prof != nil {
			profName = class.Prof.Name
		}
		st := strings.TrimSuffix(strings.TrimPrefix(fmt.Sprint(class.Students), "["), "]")
		fmt.Printf("%s %s\n", profName, st)
	}
}
