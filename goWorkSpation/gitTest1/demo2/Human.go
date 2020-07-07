package demo2

type human struct {
	username string
	password string
	age      int
}

func (stu *human) setHuman(username, password string, age int) {

	stu.username = username
	stu.password = password
	stu.age = age
}
