package internal

import (
	"encoding/json"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/stretchr/testify/suite"

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

func (suite *VisitorTestSuite) TestVisitExpr() {
	// TODO: encapsulate this
	// input := antlr.NewInputStream("5 || 3") // Lor
	// input := antlr.NewInputStream("5 || 5 && false") // Lor, right is Land
	input := antlr.NewInputStream("5 || 5 >= false") // Lor, right is Land with comparison (branch at infix3)

	// input := antlr.NewInputStream("5 * (3 + 2)")
	lexer := v1_0.NewWdlV1Lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := v1_0.NewWdlV1Parser(stream)
	ctx := p.Expr()

	expr := suite.visitor.VisitExpr(ctx)
	// expected := domain.Expression{}
	// assert.Equal(suite.T(), expr, expected)

	out, _ := json.MarshalIndent(expr, "", "    ")
	// if err != nil {
	// 	fmt.Printf("%v\n", out)
	// }

	suite.T().Log(string(out))
}

func TestVisitorTestSuite(t *testing.T) {
	suite.Run(t, new(VisitorTestSuite))
}
