package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/arjaygallentes/calcuCo/mathcalc"
	"golang.org/x/crypto/ssh/terminal"
)

var tests = []string{
	"(1+3)*7", // 28, example from task description.
	"1+3*7",   // 22, shows operator precedence.
	"7",       // 7, a single literal is a valid expression.c
	"7/3",     // eval only does integer math.
	"7.3",     // this parses, but we disallow it in eval.
	"7^3",     // parses, but disallowed in eval.
	"go",      // a valid keyword, not valid in an expression.
	"3@7",     // error message is "illegal character."
	"",        // EOF seems a reasonable error message.
}

// func main() {
// 	for _, exp := range tests {
// 		if r, err := mathcalc.ParseAndEval(exp); err == nil {
// 			fmt.Println(exp, "=", r)
// 		} else {
// 			fmt.Printf("%s: %v\n", exp, err)
// 		}
// 	}
// }

// Stores the state of the terminal before making it raw
var regularState *terminal.State

func main() {
	if len(os.Args) > 1 {
		input := strings.Replace(strings.Join(os.Args[1:], ""), " ", "", -1)
		res, err := mathcalc.ParseAndEval(input)
		if err != nil {
			return
		}
		fmt.Print("\n", res)
		return
	}

	var err error
	regularState, err = terminal.MakeRaw(0)
	if err != nil {
		panic(err)
	}
	defer terminal.Restore(0, regularState)

	term := terminal.NewTerminal(os.Stdin, "> ")
	term.AutoCompleteCallback = handleKey
	for {
		text, err := term.ReadLine()
		if err != nil {
			if err == io.EOF {
				// Quit without error on Ctrl^D
				exit()
			}
			panic(err)
		}

		text = strings.Replace(text, " ", "", -1)
		if text == "exit" || text == "quit" {
			break
		}

		res, err := mathcalc.ParseAndEval(text)
		if err != nil {
			term.Write([]byte(fmt.Sprintln("Error: " + err.Error())))
			continue
		}
		term.Write([]byte(fmt.Sprintln(res)))
	}
}

func handleKey(line string, pos int, key rune) (newLine string, newPos int, ok bool) {
	if key == '\x03' {
		// Quit without error on Ctrl^C
		exit()
	}
	return "", 0, false
}

func exit() {
	terminal.Restore(0, regularState)
	fmt.Println()
	os.Exit(0)
}
