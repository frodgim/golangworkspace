package language

import (
	"fmt"
	"slices"
)

func TestLanguage() {
	// exampleIfElse()
	// exampleSwitch()
	// exampleLoops()
	// exampleArrays()
	// exampleSlides()
	// exampleMaps()
	// exampleFunctions()

	exampleVariadic()
}

func exampleIfElse() {
	a := 1
	b := 6

	fmt.Printf("a is %v and b is %v \n", a, b)
	if a < b/3 {
		fmt.Println("a < b/3 is true")
	} else {
		fmt.Println("a < b/3 is false")

	}
}

func exampleSwitch() {
	var a int = 2
	fmt.Printf("a is %v \n", a)
	switch a {
	case 1, 2:
		fmt.Println("found case 1,2")
	case 3:
		fmt.Println("found case 3")
	default:
		fmt.Println("NOT found case")
	}

	switch {
	case a <= 5:
		fmt.Println("found case <= 5")
	default:
		fmt.Println("found case > 5")
	}

}

func exampleLoops() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("range", i)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}

func exampleArrays() {

	var arr []int = []int{0, 0, 0, 0, 0}

	var arr2 []int
	arr2 = []int{1, 2, 3, 4, 5}

	fmt.Println("Array values:", arr)
	fmt.Println("Length Array:", len(arr))
	fmt.Println("Cap Array:", cap(arr))

	fmt.Println("Array2 values:", arr2)
	fmt.Println("Length Array2:", len(arr2))
	fmt.Println("Cap Array2:", cap(arr2))

	// testDisplayArr(arr, "-arr-")
	// testDisplayArr(arr2, "-arr2-")

	for i, j := 0, len(arr)-1; i < len(arr); i, j = i+1, j-1 {
		// z:= arr[i]
		arr[i] = arr2[j]
	}

	fmt.Println("Array values:", arr)
	fmt.Println("Array2 values:", arr2)

}

func testDisplayArr(arr []int, name string) {
	for i := range arr {
		fmt.Printf("[arr = %v] Value index-%v => %v \n", name, i, arr[i])
	}
}

func exampleSlides() {
	var s []string
	fmt.Println("uninit:", s, " is nil ", s == nil, " len == 0", len(s) == 0)

	s = make([]string, 3)
	fmt.Println("Array (s):", s, "len:", len(s), "cap:", cap(s))
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	s = append(s, "d")
	s = append(s, "e", "f")

	fmt.Printf("After append d,e,f")
	fmt.Println("Array (s):", s, "len:", len(s), "cap:", cap(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:4]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[3:]
	fmt.Println("sl3:", l)

	var l2 []string = make([]string, 10)
	fmt.Println("sl4:", l2, "len(l2)=", len(l2), "cap(l2)=", cap(l2))
	l2 = s[1:3:4]
	fmt.Println("sl4:", l2, "len(l2)=", len(l2), "cap(l2)=", cap(l2))

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	fmt.Println("dcl2:", t2)
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	} else {
		fmt.Println("t != t2")
	}

	twoD := make([][]int, 10)
	for i := 0; i < 10; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}

func exampleMaps() {

	mymap := make(map[string]string)

	mymap["uno"] = "one"
	mymap["dos"] = "two"

	fmt.Println("My map len: ", len(mymap))

	for key, val := range mymap {
		fmt.Println("My map key:", key, " and val:", val)
	}
	fmt.Println("-----")

	mymap["tres"] = "threw"
	fmt.Println("My map len: ", len(mymap))
	for key := range mymap {
		fmt.Println("My map key:", key, " and val:", mymap[key])
	}

}

func exampleFunctions() {
	fmt.Println("Calling fn newfuncwith2vars(2,4)", newfuncwith2vars(2, 4))
	fmt.Println("Calling fn newfuncwith3vars(2,4,6)", newfuncwith3vars(2, 4, 6))
	fmt.Println("Calling fn newfuncwithvariadicinput(2,4,6)", newfuncwithvariadicinput(2, 4, 6))
	sum, mul := newfuncreturningmultiple(2, 4)
	fmt.Println("Calling fn newfuncreturningmultiple(2,4)", " sum:", sum, ", mul:", mul)

	a, b := 5, 10

	fmt.Printf("Orig: a=%d,b=%d \n", a, b)
	swap(&a, &b)
	fmt.Printf("Swapped: a=%d,b=%d \n", a, b)

}

func newfuncwith2vars(a int, b int) int {
	return a * b
}

func newfuncwith3vars(a, b, c int) int {
	return (a + b) * c
}

func newfuncwithvariadicinput(a int, b ...int) int {

	fmt.Println("a's value: ", a, "b's value", b)

	var ret int = a

	for _, val := range b {
		ret += val
	}

	return ret
}

func newfuncreturningmultiple(a, b int) (int, int) {
	return a + b, a * b
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func exampleVariadic() {
	fmt.Println("sum(1,2)", sum(1, 2))
	fmt.Println("sum(1,2,9)", sum(1, 2, 9))

	arrInts := []int{4, 0, 1, 87}
	fmt.Println("sum(arrInts)", sum(arrInts...))
}

func sum(a ...int) int {
	fmt.Print(a)
	sum := 0
	for _, num := range a {
		sum += num
	}

	return sum
}
