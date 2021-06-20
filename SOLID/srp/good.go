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

func CreateStudent(name string, class int) *Student {
	return &Student{
		Name:  name,
		Class: class,
	}
}

func CreateTeacher(name string, classes []int) *Teacher {
	return &Teacher{
		Name:  name,
		Class: classes,
	}
}

func CreateClass() *Class {
	teacher := CreateTeacher("colin", []int{1, 2})
	student := CreateStudent("lily", 1)
	return &Class{
		Teacher: teacher,
		Student: student,
	}
}
