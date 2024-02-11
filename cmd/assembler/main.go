package main

import (
	"fmt"
	"github.com/isongjosiah/hack/nand-tetris/assembler/parser"
)

func main() {

	fmt.Println("welcome to hack assembler")
	fmt.Println(parser.FileReader("/Users/konig/go/src/github.com/isongjosiah/hack/nand-tetris/assembler/test.asm"))
	fmt.Println(parser.FileReader("/Users/konig/go/src/github.com/isongjosiah/hack/nand-tetris/assembler/testasm.txt"))

}
