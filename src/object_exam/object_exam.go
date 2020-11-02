package objectexam

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func (b Person) in() string {
	text := fmt.Sprintf("Username: %s Age: %d Height: %f", b.name, b.age, b.height)
	return text
}
func (n People) convertListPersionNameToString() string {
	all := make([]string, len(n))
	for i, p := range n {
		all[i] = p.name
	}
	return strings.Join(all, ",")
}
func (n People) saveFile(name string) error {
	text := n.convertListPersionNameToString()
	data := []byte(text)
	return ioutil.WriteFile(name, data, 0666)
}
func readFile(name string) ([]string, error) {
	data, error := ioutil.ReadFile(name)
	if error != nil {
		log.Printf("Cannot read file %v", name)
		os.Exit(1)
		// return nil, error
	}
	text := string(data)
	return strings.Split(text, ","), error
}
func (n People) getRandomPerson() Person {
	rand.Seed(time.Now().UnixNano())
	max := len(n)
	min := 0
	randNumber := rand.Intn(max-min) + min
	return n[randNumber]
}

// MainObjectExam print all exam of object
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
		{
			name:   "Ảnh",
			age:    22,
			height: 1.81,
		},
		{
			name:   "Việt",
			age:    19,
			height: 1.81,
		},
	}
	for i, s := range people {
		fmt.Println(i, s.in())
	}
	// fmt.Println(people[:2])
	// people.saveFile("user")
	// name, _ := readFile("user")
	// fmt.Println("Text from file", name)
	randUser := people.getRandomPerson()
	fmt.Printf("Random user with name: %v", randUser.name)

	return
}
