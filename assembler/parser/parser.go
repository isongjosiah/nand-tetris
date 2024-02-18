package parser

import (
	"errors"
	"fmt"
	"os"
	"path"
)

type Parser struct {
	FilePath string
	File     *os.File
}

/*
FileReader parses an assembly file (ends in .asm) .
- It accepts a file path, ensures it is a valid file and confirms that it is an asm file
*/
func (p *Parser) FileReader() error {

	// validate this file exists or is valid
	fileInfo, err := os.Stat(p.FilePath)
	if err != nil {
		return err
	}

	// confirm that it is an asm file -> check that the extension is asm
	baseName := fileInfo.Name()
	if path.Ext(baseName) != ".asm" {
		return errors.New(fmt.Sprintf("invalid file %s. only asm files are supported", baseName))
	}

	// read the file
	file, err := os.Open(p.FilePath)
	if err != nil {
		return errors.New(fmt.Sprintf("unable to open file %s. %v", baseName, err.Error()))
	}
	p.File = file

	return nil
}

/*
NewParser creates a new instance of the parser with the provided file path
*/
func NewParser(filePath string) *Parser {
	return &Parser{
		FilePath: filePath,
	}
}
