package impl

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// An IVisitorReporter that uses the fmt package to report messages.
type FmtReporter struct{}

func (r *FmtReporter) Warn(ctx antlr.ParserRuleContext, warn string) {
	fmt.Printf("WARN: %v\n", warn)
}

func (r *FmtReporter) Error(ctx antlr.ParserRuleContext, err error) {
	fmt.Printf("Error: %v\n", err)
}

func (r *FmtReporter) NotImplemented(message fmt.Stringer) {
	fmt.Printf("NotImplemented: %v\n", message)
}
