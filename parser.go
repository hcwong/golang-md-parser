package golangmdtty

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	blackfriday "gopkg.in/russross/blackfriday.v2"
)

// TtyRenderer implements the BlackFriday renderer
type TtyRenderer struct {
	// TODO: define some config options here from the config file
	current          blackfriday.NodeType // current provides child nodes with context of non-immediate parent nodes
	indentationLevel int
	tabSize          int
	orderedListNum   []int
}

func createRenderer(tabSize int) TtyRenderer {
	return TtyRenderer{tabSize: tabSize, indentationLevel: 0, orderedListNum: []int{}}
}

func (r *TtyRenderer) RenderNode(w io.Writer, node *blackfriday.Node, entering bool) (status blackfriday.WalkStatus) {
	switch node.Type {
	case blackfriday.Document:
	case blackfriday.BlockQuote:
	case blackfriday.List:
		if entering {
			r.indentationLevel++
			r.orderedListNum = r.orderedListNum[:r.indentationLevel]
			r.orderedListNum[r.indentationLevel-1] = 0
		} else {
			r.indentationLevel--
			w.Write([]byte("\n"))
		}
	case blackfriday.Item:
		if entering {
			r.orderedListNum[r.indentationLevel]++
			i := 0
			for i < (r.indentationLevel * r.tabSize) {
				w.Write([]byte(" "))
				i++
			}
			if node.ListData.ListFlags == 0 {
				w.Write([]byte(strconv.Itoa(r.orderedListNum[r.indentationLevel]) + string(node.ListData.BulletChar) + ". "))
			} else {
				w.Write([]byte("* "))
			}
		}
	case blackfriday.Paragraph:
		if entering {
			w.Write([]byte("\n\n"))
		}
	case blackfriday.Heading:
		if entering {
			headingLevel := node.HeadingData.Level
			for headingLevel > 0 {
				w.Write([]byte("#"))
				headingLevel--
			}
		} else {
			w.Write([]byte("\n"))
		}
	case blackfriday.HorizontalRule:
	case blackfriday.Emph:
	case blackfriday.Strong:
	case blackfriday.Del:
	case blackfriday.Link:
		if entering {
			w.Write([]byte("["))
		} else {
			w.Write([]byte("]"))
			w.Write([]byte("("))
			w.Write(node.LinkData.Destination)
			w.Write([]byte(")"))
		}
	case blackfriday.Image:
	// Text should always be a leaf node
	case blackfriday.Text:
		if entering {
			w.Write(node.Literal)
		}
	case blackfriday.HTMLBlock:
	case blackfriday.CodeBlock:
		if entering {
			w.Write([]byte("````\n"))
			r.indentationLevel++
			r.current = blackfriday.CodeBlock
		} else {
			r.indentationLevel--
			w.Write([]byte("````\n"))
		}
	case blackfriday.Softbreak:
	case blackfriday.Hardbreak:
	case blackfriday.Code:
		if entering {
			w.Write(node.Literal)
		}
	case blackfriday.HTMLSpan:
	case blackfriday.Table:
		if entering {
			r.current = blackfriday.Table
		}
	case blackfriday.TableCell:
	case blackfriday.TableHead:
	case blackfriday.TableBody:
	case blackfriday.TableRow:
	}

	return blackfriday.GoToNext
}

func (r *TtyRenderer) RenderHeader(w io.Writer, ast *blackfriday.Node) {

}

func (r *TtyRenderer) RenderFooter(w io.Writer, ast *blackfriday.Node) {

}

// Output the indentation of a node automatically
func (r *TtyRenderer) outputIndentation(w io.Writer) {
	i := 0
	for i < (r.indentationLevel * r.tabSize) {
		w.Write([]byte(" "))
		i++
	}
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
	// TOFIX
	fmt.Println(line)
}
