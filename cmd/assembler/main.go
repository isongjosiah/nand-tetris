package main

import (
	"fmt"
	"log"
	"nand2Tetris/assembler/hackassembler"
	"nand2Tetris/assembler/parser"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	go time.Sleep(time.Second * 2)
	fmt.Println("welcome to hack assembler")

	engine := hackassembler.NewAssembler("/Users/konig/go/src/github.com/isongjosiah/hack/nand-tetris/assembler/test.asm")
	err := engine.Parser.FileReader()
	if err != nil {
		log.Fatal(err)
	}

	err = engine.ParseLabels(parser.Labels)
	if err != nil {
		log.Fatal(err)
	}

	err = engine.ParseLabels(parser.Variables)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(engine.SymbolTable)

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-shutdown

	fmt.Println("shutting dow hack assembler")

}
