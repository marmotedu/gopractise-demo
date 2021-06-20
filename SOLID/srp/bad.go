package srp

type Class struct {
	Teacher *Teacher
	Student *Student
}

type Teacher struct {
	Name  string
	Class int
}

type Student struct {
	Name  string
	Class int
}

func createClass(teacherName, studentName string, class int) (*Teacher, *Student) {
	teacher := &Teacher{
		Name:  teacherName,
		Class: class,
	}
	student := &Student{
		Name:  studentName,
		Class: class,
	}

	return teacher, student
}

func CreateClass() *Class {
	teacher, student := createClass("colin", "lily", 1)
	return &Class{
		Teacher: teacher,
		Student: student,
	}
}
