package main

import (
	"fmt"
	"os"
	
	"michaelcanudas.dough/commands"
)

func main() {
	args := os.Args[1:]
	
	if len(args) == 0 {
		fmt.Println("no command")
		return
	}
	
	switch args[0] {
	case "build":
		build(args)
	case "run":
		run(args)
	}
}

func build(args []string) {
	_, err := commands.Build(args[1:])
	if err != nil {
		panic(err)
	}
	
	fmt.Println("built!")
}

func run(args []string) {
	err := commands.Run(args[1:])
	if err != nil {
		panic(err)
	}
}
