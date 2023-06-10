package interpreter

import (
	"strings"
)

func init() {
	RegisterInstruction("rsrc", func(fields []string) Instruction {
		var content ResourceContent

		if strings.HasPrefix(fields[2], "\"") {
			// string content
			content = StringResourceContent{
				Value: unescape(fields[2]),
			}
		} else {
			panic("error parsing resource contents")
		}

		return Rsrc{
			Name:    fields[1],
			Content: content,
		}
	})
}

func unescape(s string) string {
	r := s[1 : len(s)-1]
	r = strings.Replace(r, `\n`, "\n", -1)
	r = strings.Replace(r, `\"`, "\"", -1)
	return r
}

type Rsrc struct {
	Name    string
	Content ResourceContent
}

func (r Rsrc) execute(ctx *Context) {
	// nop
}

func (r Rsrc) getName() string {
	return r.Name
}

type ResourceContent interface {
	getData() []byte
}

type StringResourceContent struct {
	Value string
}

func (s StringResourceContent) getData() []byte {
	return []byte(s.Value)
}
