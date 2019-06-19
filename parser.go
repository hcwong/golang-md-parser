package golangmdtty

import (
	"bufio"
	"fmt"
	"strings"
)

type renderer interface {
}

// Parser is the object we use to parse the file and output it
type parser struct {
	scanner bufio.Scanner
}

// We define a markdown block as having two EOF characters in a row (one line spacing)
func splitFileFunction(data []byte, atEOF bool) (advanceBy int, block []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if atEOF {
		return len(data), data, nil
	}

	index := strings.Index(string(data), "\n\n")
	if index != -1 {
		return index + 1, data[0:index], nil
	}
	return
}

func (p *parser) setScannerSplit() {
	p.scanner.Split(splitFileFunction)
}

func (p *parser) getNextLine() {
	for p.scanner.Scan() {
		parseLine(p.scanner.Text())
	}
}

func parseLine(line string) {
	// TODO fix this, currently just testing
	fmt.Println(line)
}
