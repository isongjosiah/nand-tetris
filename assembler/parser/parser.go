package parser

import (
	"errors"
	"fmt"
	"os"
	"path"
)

/*
FileReader parses an assembly file (ends in .asm) and generates a symbol table.
- It accepts a file path, ensures it is a valid file and confirms that it is an asm file
- Constructs a symbol table (labels and variables) and adds it to the global symbol table containing predefined symbols
*/
func FileReader(filePath string) error {

	// validate this file exists or is valid
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	// confirm that it is an asm file -> check that the extension is asm
	baseName := fileInfo.Name()
	if path.Ext(baseName) != ".asm" {
		return errors.New(fmt.Sprintf("invalid file %s. only asm files are supported", baseName))
	}

	// read the file

	return nil

}
