package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/shaneu/monkey/pkg/evaluator"
	"github.com/shaneu/monkey/pkg/object"
	"github.com/shaneu/monkey/pkg/parser"

	"github.com/shaneu/monkey/pkg/lexer"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		result := evaluator.Eval(program, env)
		if result != nil {

			io.WriteString(out, result.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
