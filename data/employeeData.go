package data

type Employee struct {
	Name string
	Age  int
}

type Student struct {
	Name string
	Age  int
}

var names = []string{"Mayank", "Anshul", "Ankit", "Anupam", "Meha"}
var studentName = []string{"Student: Mayank", "Student: Anshul", "Student: Ankit", "Student: Anupam", "Student: Meha"}
var empList = []Employee{}
var studList = []Student{}

func GetEmployeeData() []Employee {
	if len(empList) > 0 {
		return empList
	}
	var emp Employee
	for i := 0; i < len(names); i++ {
		emp = Employee{names[i], (30 + i)}
		empList = append(empList, emp)
	}
	return empList
}

func GetStudentData() []Student {
	if len(studList) > 0 {
		return studList
	}
	var stud Student
	for i := 0; i < len(studentName); i++ {
		stud = Student{studentName[i], (30 + i)}
		studList = append(studList, stud)
	}
	return studList
}
