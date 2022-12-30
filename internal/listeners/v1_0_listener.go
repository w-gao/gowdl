package listeners

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/w-gao/gowdl/parsers"
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
