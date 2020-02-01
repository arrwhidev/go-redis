package parser

// General Parser interface
type Parser interface {
	Parse() (command string)
}

