package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/yosugi/md2tt/lib/freeze"
	"github.com/yosugi/md2tt/lib/functional"
	"github.com/yosugi/md2tt/lib/readline"
)

const defaultIndent = 4

// int -> (Freeze -> Freeze)
func md2tt(indent int) func(string) {
	composed := functional.Compose(
		freeze.Then(heading),
		freeze.Then(bulletList(indent)),
		freeze.Then(numberedList(indent)),
	)

	return func(line string) {
		freeze := freeze.Wrap(line)
		freeze = composed(freeze)
		fmt.Println(freeze.Unwrap())
	}
}

// TODO mod README
// add test monad law
func main() {
	indent := defaultIndent
	if len(os.Args) >= 2 {
		argIndent, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
		indent = argIndent
	}

	err := readline.On(md2tt(indent))
	if err != nil {
		panic(err)
	}
}
