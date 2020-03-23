package internal

import (
	"bufio"
	"fmt"
	"os"
)

type Repl struct{}

func (r *Repl) RunForever() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func NewRepl(publisher Publisher) *Repl {
	return &Repl{}
}
