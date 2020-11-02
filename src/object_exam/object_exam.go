package object_exam

import "fmt"

func (b Person) in() string {
	text := fmt.Sprintf("Username: %s Age: %d Height: %f", b.name, b.age, b.height)
	return text
}

func MainObjectExam() {
	people := People{
		{
			name:   "Nam",
			age:    19,
			height: 1.71,
		},
		{
			name:   "Ngọc",
			age:    20,
			height: 1.81,
		},
	}
	for i, s := range people {
		fmt.Println(i, s.in())
	}
	return
}
