package slices

import (
	"fmt"
)

func TestSlices() {
	fmt.Println("This is a slice of strings")

	var aString []string = []string{"uno", "dos", "tres"}

	for i, v := range aString {
		fmt.Printf("index = %v, value = %v \n", i, v)
	}

	var bString []string = make([]string, 3, 10)

	bString[2] = "ittttiiis"

	bString = append(bString, "hshshhshsh")

	bString = append(bString, "wwwww", "paooaoaoao", "asdasdsadsd")
	for i, v := range bString {
		fmt.Printf("index = %v, value = %v \n", i, v)
	}

	fmt.Printf("Type bString = %T \n", bString)

	fmt.Printf("Type bString slice = %T \n", bString[4:6])
	fmt.Println(bString[4:6])

	a := [5]int{1, 2, 3, 4, 5}
	t := a[1:3:5]

	fmt.Println(a)
	fmt.Println(t)
	fmt.Printf("Cap t = %v \n", cap(a))
	fmt.Printf("Cap t = %v \n", cap(t))
	t = append(t, 6, 7, 8)
	fmt.Println(t)

}
