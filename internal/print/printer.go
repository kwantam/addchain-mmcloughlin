package print

import (
	"fmt"
	"io"
	"strings"
)

const DefaultIndent = "\t"

type Printer struct {
	out     io.Writer
	level   int    // current indentation level
	indent  string // indentation string
	pending bool   // if there's a pending indentation
	err     error  // saved error from printing
}

func New(w io.Writer) Printer {
	return Printer{
		out:    w,
		indent: DefaultIndent,
	}
}

func (p *Printer) SetIndentString(indent string) {
	p.indent = indent
}

func (p *Printer) Indent() {
	p.level++
}

func (p *Printer) Dedent() {
	p.level--
}

func (p *Printer) Linef(format string, args ...interface{}) {
	p.Printf(format, args...)
	p.NL()
}

func (p *Printer) NL() {
	p.Printf("\n")
	p.pending = true
}

func (p *Printer) Printf(format string, args ...interface{}) {
	if p.err != nil {
		return
	}
	if p.pending {
		indent := strings.Repeat(p.indent, p.level)
		format = indent + format
		p.pending = false
	}
	_, err := fmt.Fprintf(p.out, format, args...)
	p.SetError(err)
}

func (p *Printer) Error() error {
	return p.err
}

func (p *Printer) SetError(err error) {
	if p.err == nil {
		p.err = err
	}
}