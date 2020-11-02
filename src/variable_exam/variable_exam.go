package variableexam

import (
	"fmt"
	"reflect"
	"strconv"
)

func incressCounter() int {
	counter := 1
	return counter
}

func arrayTest() [3]int {
	var numbers [3]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	return numbers
}

func slidesTest() []int {
	numbers := make([]int, 3)
	numbers[0] = 4
	return numbers
}

// MainVariableExam excute print variable lession
func MainVariableExam() {
	var hello string = "Xin chào hihi...! "
	counter := 1
	customArray := arrayTest()
	customArray2 := slidesTest()
	fmt.Println(hello + strconv.Itoa(counter) + " - " + strconv.Itoa(incressCounter()) + " Array: " + reflect.ValueOf(customArray).String())
	fmt.Println(customArray2)
	for index, num := range customArray2 {
		text := fmt.Sprintf("Index %d and Number %d", index, num)
		fmt.Println(text)
	}
	return

}
