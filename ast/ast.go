package ast

// Statement Node
type Stmt interface {
	stmt()
}

// Expression Node
type Expr interface {
	expr()
}
