package parser

import(
	"coding_language/ast"
	"coding_language/lexer"
	"coding_language/token"
)

type Parser struct{
	l *lexer.Lexer //lexer
	
	curToken token.Token  //current token
	peekToken token.Token //next token
}


func New(l *lexer.Lexer) *Parser{ //new parser, takes lexer as args
	p := &Parser{l: l}

	p.l.NextToken()
	p.l.NextToken()

	return p
}


func (p *Parser) ParseProgram() *ast.Program{
	return nil
}

