package parser

import (
	"bufio"
	"fmt"
	"strings"
)

// Parser is VM parser.
type Parser struct {
	scanner  *bufio.Scanner
	text     string
	commands []string
}

// NewParser create ner Parser struct.
func NewParser(scanner *bufio.Scanner) *Parser {
	return &Parser{scanner, "", make([]string, 0)}
}

// HasMoreCommands return to be possible to scan.
func (p *Parser) HasMoreCommands() bool {
	return p.scanner.Scan()
}

func (p *Parser) ScannerError() error {
	return p.scanner.Err()
}

func (p *Parser) Text() string {
	return p.text
}

// Advance read current row text.
func (p *Parser) Advance() {
	p.text = p.scanner.Text()
	p.commands = strings.Split(p.text, " ")
}

func (p *Parser) CommandType() string {
	return p.text
}

func (p *Parser) Arg1() (string, error) {
	if len(p.commands) < 1 {
		return "", fmt.Errorf("This commands has no first argument. commands: %s", p.text)
	}
	t := p.commands[0]

	return t, nil
}

func (p *Parser) Arg2() (string, error) {
	if len(p.commands) < 2 {
		return "", fmt.Errorf("This commands has no second argument. commands: %s", p.text)
	}
	t := p.commands[0]

	return t, nil
}
