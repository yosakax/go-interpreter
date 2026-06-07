package runner

import (
	"io"
	"os"

	"yosaka/go-interpreter/evaluator"
	"yosaka/go-interpreter/lexer"
	"yosaka/go-interpreter/object"
	"yosaka/go-interpreter/parser"
)

func Run(filename string, out io.Writer) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	data := make([]byte, 1024)
	_, err = f.Read(data)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	input := string(data)
	env := object.NewEnvironment()
	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")
	}
}
