package parser

type Parser interface {
	Parse() (command []string, err error)
}
