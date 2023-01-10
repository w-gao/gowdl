package internal

import (
	"encoding/json"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/stretchr/testify/suite"

	"github.com/w-gao/gowdl/internal/domain"
	"github.com/w-gao/gowdl/parsers/v1_0"
)

type VisitorTestSuite struct {
	suite.Suite
	visitor *WdlVisitor
}

func (suite *VisitorTestSuite) SetupTest() {
	// Called before each test
	suite.visitor = NewWdlVisitor("mock_url", "1.0")
}

func (suite *VisitorTestSuite) parseExpr(expr string) domain.IExpression {
	input := antlr.NewInputStream(expr)
	lexer := v1_0.NewWdlV1Lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := v1_0.NewWdlV1Parser(stream)
	ctx := p.Expr()

	return suite.visitor.VisitExpr(ctx)
}

func (s *VisitorTestSuite) TestVisitExpr() {
	// expr := suite.parseExpr("5 || 5 >= false") // Lor, right is Land with comparison (branch at infix3)
	// expr := suite.parseExpr("5 * (3 + 2)")

	// expected := domain.Expression{}
	// assert.Equal(suite.T(), expr, expected)

	s.Run("Lor", func() {
		expr := s.parseExpr("5 || 3")
		out, _ := json.MarshalIndent(expr, "", "    ")
		s.T().Log(string(out))
	})

	s.Run("Lor, right is Land", func() {
		expr := s.parseExpr("5 || 5 && false")
		out, _ := json.MarshalIndent(expr, "", "    ")
		s.T().Log(string(out))
	})

	s.Run("Land", func() {
		expr := s.parseExpr("5 && 3")
		out, _ := json.MarshalIndent(expr, "", "    ")
		s.T().Log(string(out))
	})

	s.Run("Lor, left is Land", func() {
		expr := s.parseExpr("5 && 3 || false")
		out, _ := json.MarshalIndent(expr, "", "    ")
		s.T().Log(string(out))
	})

	s.Run("Land, right is BinaryOp (infix 4)", func() {
		expr := s.parseExpr("5 && 5 + 3")
		out, _ := json.MarshalIndent(expr, "", "    ")
		s.T().Log(string(out))
	})

	s.Run("Land, right is BinaryOp (infix 5)", func() {
		expr := s.parseExpr("5 && 5 / 3")
		out, _ := json.MarshalIndent(expr, "", "    ")
		s.T().Log(string(out))
	})

	s.Run("BinaryOp (infix 4 and 5)", func() {
		expr := s.parseExpr("5 + 5 / 3")
		out, _ := json.MarshalIndent(expr, "", "    ")
		s.T().Log(string(out))
	})

}

func TestVisitorTestSuite(t *testing.T) {
	suite.Run(t, new(VisitorTestSuite))
}
