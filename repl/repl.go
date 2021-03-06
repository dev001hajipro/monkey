package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/dev001hajipro/monkey/evaluator"
	"github.com/dev001hajipro/monkey/lexer"
	"github.com/dev001hajipro/monkey/object"
	"github.com/dev001hajipro/monkey/parser"
)

const MONKEY_FACE = `
           __,__
   .--. .-"     "-.  .--.
  / ..\/  .-. .-.  \/ .. \
 | | '|  /   Y   \  |'  | |
 | \  \  \ ^ | ^ /  /   / |
 \ '- ,\.-"""""""-./, -' /
  ''-' /_   ^ ^   _\ '-''
      |  \._   _./  |
      \   \ '~' /   /
       '._ '-=-' _.'
          '-----'
`

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	macroEnv := object.NewEnvironment()

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}

		// debug print
		io.WriteString(out, "[DEBUG] " + program.String())
		io.WriteString(out, "\n")

		// I COULD SUPPORT MACRO!!
		// search macros and register macros.
		evaluator.DefineMacro(program, macroEnv)
		expanded := evaluator.ExpandMacro(program, macroEnv)

		evaluated := evaluator.Eval(expanded, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParseErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser error:\n")

	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
