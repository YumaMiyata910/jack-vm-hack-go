package parser

const (
	CARITHMETIC string = "C_ARITHMETIC"
	CPUSH       string = "C_PUSH"
	CPOP        string = "C_POP"
	CLABEL      string = "C_LABEL"
	CGOTO       string = "C_GOTO"
	CIF         string = "C_IF"
	CFUNCTION   string = "C_FUNCTON"
	CRETURN     string = "C_RETURN"
	CCALL       string = "C_CALL"

	PUSH string = "push"
	POP  string = "pop"

	ADD string = "add"
	SUB string = "sub"
	NEG string = "neg"
	EQ  string = "eq"
	GT  string = "gt"
	LT  string = "lt"
	AND string = "and"
	OR  string = "or"
	NOT string = "not"

	ARGUMENT string = "argument"
	LOCAL    string = "local"
	STATIC   string = "static"
	CONSTANT string = "constant"
	THIS     string = "this"
	THAT     string = "that"
	POINTER  string = "pointer"
	TEMP     string = "temp"
)
