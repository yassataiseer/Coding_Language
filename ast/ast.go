package ast

import(
	"coding_language/token"
)


type Node interface{
	TokenLiteral() string
}

type Statement interface{
	Node
	statementNode()
}


type Expression interface{
	Node 
	expressionNode()
}

type Program struct{
	Statements []Statement
}
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
	}
type LetStatement struct{
	Token token.Token
	Name *Identifier
	Value Expression

}

func (ls *LetStatement) statementNode() {

}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal 
}

func (p *Program) TokenLiteral() string{
	if len(p.Statements)>0{
		return p.Statements[0].TokenLiteral()
	}else{
		return ""
	}

}

func (i *Identifier) expressionNode() {

}
func (i *Identifier) TokenLiteral() string { 
	return i.Token.Literal 
}

