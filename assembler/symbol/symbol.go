package symbol

import (
	"fmt"
)

/*
NewSymbolTable creates a new symbole table with the predefined symbols
setup
*/
func NewSymbolTable() map[string]string {

	symbol := make(map[string]string)

	for i := 0; i < 16; i++ {
		symbol[fmt.Sprintf("R%v", i)] = DecToBin(i)
	}

	symbol["SCREEN"] = DecToBin(16384)
	symbol["KBD"] = DecToBin(24576)
	symbol["SP"] = DecToBin(0)
	symbol["LCL"] = DecToBin(1)
	symbol["ARG"] = DecToBin(2)
	symbol["THIS"] = DecToBin(3)
	symbol["THAT"] = DecToBin(4)

	return symbol

}
