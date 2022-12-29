package main

import (
	"fmt"

	"github.com/w-gao/gowdl/internal/cmd"
	"github.com/w-gao/gowdl/parsers"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type WdlV1Listener struct {
	*parsers.BaseWdlV1ParserListener
}

func NewWdlV1Listener() *WdlV1Listener {
	return new(WdlV1Listener)
}

func (this *WdlV1Listener) EnterVersion(ctx *parsers.VersionContext) {
	fmt.Printf(ctx.ReleaseVersion().GetText())
}

func (this *WdlV1Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	// fmt.Println(ctx.GetText())
}

func main() {
	// input, _ := antlr.NewFileStream(os.Args[1])
	// lexer := parsers.NewWdlV1Lexer(input)
	// stream := antlr.NewCommonTokenStream(lexer, 0)
	// p := parsers.NewWdlV1Parser(stream)
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	// p.BuildParseTrees = false
	// tree := p.Document()
	// antlr.ParseTreeWalkerDefault.Walk(NewWdlV1Listener(), tree)

	// version, err := internal.GetVersion(os.Args[1])
	// if err != nil {
	// 	fmt.Printf("error: %s\n", err)
	// 	return
	// }

	// fmt.Printf("WDL version: '%s'\n", version)

	cmd.Execute()
}
