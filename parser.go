package golangmdtty

import (
	"bufio"
	"fmt"
)

// Parser is the object we use to parse the file and output it
type Parser struct {
	scanner bufio.Scanner
}

func (p *Parser) getNextLine() {
	for p.scanner.Scan() {
		parseLine(p.scanner.Text())
	}
}

func parseLine(line string) {
	// TODO fix this, currently just testing
	fmt.Println(line)
}
