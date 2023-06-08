package interpreter

import (
	"fmt"
	"strconv"
)

type Read struct {
}

func init() {
	RegisterInstruction("read", func(fields []string) Instruction {
		return Read{}
	})
}

func (r Read) execute(ctx *Context) {
	var value int64 = 0

	for {
		fmt.Print("waiting on input: ")
		// read a word
		var text string
		var _, err = fmt.Scan(&text)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		// convert it to int
		var v int64
		v, err = strconv.ParseInt(text, 10, 64)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		// we've got a value, take it and break
		value = v
		break
	}

	// push value to stack
	ctx.ExecutionStack.Push(value)
}
