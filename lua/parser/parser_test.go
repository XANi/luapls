package parser

import (
	"fmt"
	"testing"

	"github.com/raiguard/luapls/lua/ast"
	"github.com/raiguard/luapls/lua/lexer"

	"github.com/stretchr/testify/require"
)

func checkParserErrors(t *testing.T, p *Parser) {
	if len(p.errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(p.errors))
	for _, msg := range p.errors {
		t.Errorf("parser error: %q", msg)
	}
	t.Fail()
}

func testNumberLiteral(t *testing.T, il ast.Expression, value float64) {
	integ, ok := il.(*ast.NumberLiteral)
	require.True(t, ok)
	require.Equal(t, value, integ.Value)
	require.Equal(t, fmt.Sprintf("%.0f", value), integ.Token.Literal)
}

type statementTest struct {
	input, expected string
}

func testStatements(t *testing.T, tests []statementTest) {
	for _, test := range tests {
		l := lexer.New(test.input)
		p := New(l)

		block := p.ParseBlock()
		require.NotNil(t, block)
		require.Equal(t, 1, len(block.Statements))

		require.Equal(t, test.expected, block.String())
	}
}
