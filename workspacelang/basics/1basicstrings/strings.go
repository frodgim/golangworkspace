package basicstrings

import (
	"fmt"
	"strings"
)

func ContainsStr(str string, substr string) bool {
	return strings.Contains(str, substr)
}

func Misc() {
	fmt.Printf("%v %v, %v \n", strings.Repeat("Paco", 3), "que mi Paco", strings.Repeat("Paco", 3))

	var arrayPacos []string = []string{"paco", "paco", "paco"}
	fmt.Println(arrayPacos)
	fmt.Printf("%v %v, %v", strings.Repeat("Paco", 3), "que mi Paco", strings.Join(arrayPacos, "-"))
}
