package main

import (
	"fmt"

	"github.com/suifei/xsd2code/pkg/types"
)

func main() {
	fmt.Printf("testElement -> %s\n", types.ToPascalCase("testElement"))
	fmt.Printf("test_element -> %s\n", types.ToPascalCase("test_element"))
	fmt.Printf("TestElement -> %s\n", types.ToPascalCase("TestElement"))
}
