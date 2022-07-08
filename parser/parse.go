package parser

type Parser interface {
	Decode(interface{}) error
}

func Parse(parser Parser, v interface{}) error {
	return parser.Decode(v)
}
