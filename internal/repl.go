package internal

import (
	"bufio"
	"fmt"
	"os"
)

type Repl struct {
	publisher Publisher
}

func (r *Repl) RunForever() {

	fmt.Printf("Making new stdin scanner\n")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Starting loop over input lines\n")
	for scanner.Scan() {
		fmt.Printf("About to publish input text: <%v>\n", scanner.Text())
		r.publisher.Publish(scanner.Bytes())
		fmt.Printf("Publish succeeded\n")
	}
	fmt.Printf("Loop over input lines terminated\n")
}

func NewRepl(publisher Publisher) *Repl {
	return &Repl{publisher: publisher}
}
