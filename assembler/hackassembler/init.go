package hackassembler

import (
	"bufio"
	"fmt"
	"nand2Tetris/assembler/generator"
	"nand2Tetris/assembler/parser"
	"nand2Tetris/assembler/symbol"
	"os"
	"strings"
)

/*
Assembler is the engine for the hack assebmler. It processes a .asm file and converts it
to machine code using the hack computer instruction set for the hack computer to process
*/
type Assembler struct {
	Parser      *parser.Parser
	Generator   *generator.Generator
	SymbolTable map[string]string
}

/*
NewAssembler opens a .asm file, and gets it ready to process. It also constructs a symbol table
Adds all the predefined symbol to it
*/
func NewAssembler(filePath string) *Assembler {

	asm := new(Assembler)

	asm.Parser = parser.NewParser(filePath)
	asm.Generator = generator.NewGenerator()
	asm.SymbolTable = symbol.NewSymbolTable()

	return asm
}

func (a *Assembler) ParseLabels(sType parser.SymbolType) error {

	lineNumber := 0
	stackNumber := 16

	scanFile := new(os.File)
	*scanFile = *a.Parser.File
	fileScanner := bufio.NewScanner(scanFile)
	for fileScanner.Scan() {

		fileLine := fileScanner.Text()
		fileLine = strings.Trim(fileLine, " ")
		fileLine = strings.Trim(fileLine, "\n")
		fileLine = strings.Trim(fileLine, "\t")
		if strings.HasPrefix(fileLine, "//") || len(fileLine) == 0 {
			continue
		}

		switch sType {
		case parser.Labels:

			if strings.HasPrefix(fileLine, "(") && strings.HasSuffix(fileLine, ")") {
				fileLine = strings.Trim(fileLine, "(")
				fileLine = strings.Trim(fileLine, ")")
				a.SymbolTable[fileLine] = symbol.DecToBin(lineNumber + 1)
			} else {
				lineNumber++
			}

		case parser.Variables:

			fmt.Println("dealing with variables")

			if strings.HasPrefix(fileLine, "@") {
				fmt.Println("dealing with variables")
				fileLine = strings.Trim(fileLine, "@")
				// check if it already exists in map
				if _, found := a.SymbolTable[fileLine]; !found {
					a.SymbolTable[fileLine] = symbol.DecToBin(stackNumber)
					stackNumber++
				}

				lineNumber++
			}

		default:
			fmt.Println("inside default")
			continue
		}
	}
	return nil
}
