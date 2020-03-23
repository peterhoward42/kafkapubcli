package internal

import (
	"bufio"
	"os"
)

type Repl struct {
	publisher Publisher
}

func (r *Repl) RunForever() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		r.publisher.Publish(scanner.Bytes())
	}
}

func NewRepl(publisher Publisher) *Repl {
	return &Repl{publisher: publisher}
}
