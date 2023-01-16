package internal

import (
	"fmt"
	urlparse "net/url"
	"path"
	"reflect"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/w-gao/gowdl/internal/domain"
	"github.com/w-gao/gowdl/internal/impl"
)

type IVisitorReporter interface {
	Warn(antlr.ParserRuleContext, string)
	Error(antlr.ParserRuleContext, error)
	NotImplemented(fmt.Stringer)
}

type CustomErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []error
}

func (c *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	c.Errors = append(c.Errors, fmt.Errorf("SyntaxError: line %d:%d %s", line, column, msg))
}

type WdlVisitor struct {
	// Let's move away from the generated interface which is ambiguous with
	// typing. This way we'd lose some functionality such as tree.Accept(v),
	// since the visitor no longer implements the antlr.ParseTreeVisitor
	// interface, but we mitigate that with the Visit() function.
	// parsers.BaseWdlV1ParserVisitor

	Url      string
	Version  string
	Reporter IVisitorReporter
}

func NewWdlVisitor(url string, version string) *WdlVisitor {
	reporter := &impl.FmtReporter{}
	return &WdlVisitor{
		Url:      url,
		Version:  version,
		Reporter: reporter,
	}
}

// Visit takes in any parse tree and visits the child nodes from there.
func (v *WdlVisitor) Visit(tree antlr.ParseTree) interface{} {
	// fmt.Printf("%v\n", reflect.TypeOf(tree).Elem().Name())

	switch t := tree.(type) {
	case domain.IDocumentContext:
		return v.VisitDocument(t)
	default:
		return nil
	}
}

func (v *WdlVisitor) VisitChildren(node antlr.RuleNode) interface{}     { return nil }
func (v *WdlVisitor) VisitTerminal(node antlr.TerminalNode) interface{} { return nil }
func (v *WdlVisitor) VisitErrorNode(node antlr.ErrorNode) interface{}   { return nil }

// func (v *WdlVisitor) VisitMap_type(ctx *parsers.Map_typeContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitArray_type(ctx *parsers.Array_typeContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitPair_type(ctx *parsers.Pair_typeContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitType_base(ctx *parsers.Type_baseContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

func (v *WdlVisitor) VisitWdl_type(ctx domain.IWdl_typeContext) domain.Type {
	// TODO: parse WDL type
	return domain.Type{
		Optional: true,
	}
}

func (v *WdlVisitor) VisitUnbound_decls(ctx domain.IUnbound_declsContext) domain.Declaration {
	identifier := domain.Identifier(ctx.Identifier().GetText())
	var type_ domain.Type

	for _, child := range ctx.GetChildren() {
		switch tt := child.(type) {
		case domain.IWdl_typeContext:
			type_ = v.VisitWdl_type(tt)
		}
	}

	return domain.Declaration{
		Type:       type_,
		Identifier: identifier,
		Expr:       nil,
	}
}

func (v *WdlVisitor) VisitBound_decls(ctx domain.IBound_declsContext) domain.Declaration {
	identifier := domain.Identifier(ctx.Identifier().GetText())
	var type_ domain.Type
	var expr domain.IExpression

	for _, child := range ctx.GetChildren() {
		switch tt := child.(type) {
		case domain.IWdl_typeContext:
			type_ = v.VisitWdl_type(tt)
		case domain.IExprContext:
			expr = v.VisitExpr(tt)
		}
	}

	return domain.Declaration{
		Type:       type_,
		Identifier: identifier,
		Expr:       &expr,
	}
}

func (v *WdlVisitor) VisitAny_decls(ctx domain.IAny_declsContext) domain.Declaration {
	// First child _should_ be the declaration.
	switch tt := ctx.GetChild(0).(type) {
	case domain.IBound_declsContext:
		return v.VisitBound_decls(tt)
	case domain.IUnbound_declsContext:
		return v.VisitUnbound_decls(tt)
	}

	// This should not happen. Should we panic?
	return domain.Declaration{}
}

// func (v *WdlVisitor) VisitNumber(ctx *parsers.NumberContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitExpression_placeholder_option(ctx *parsers.Expression_placeholder_optionContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

func (v *WdlVisitor) VisitString_part(ctx domain.IString_partContext) string {
	var parts []string
	for _, part := range ctx.AllStringPart() {
		parts = append(parts, part.GetText())
	}

	return strings.Join(parts, "")
}

// func (v *WdlVisitor) VisitString_expr_part(ctx *parsers.String_expr_partContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitString_expr_with_string_part(ctx *parsers.String_expr_with_string_partContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

func (v *WdlVisitor) VisitString(ctx domain.IStringContext) string { // interface{} {
	// TODO: parse actual string, which can contain expressions!
	for _, child := range ctx.GetChildren() {
		switch tt := child.(type) {
		case domain.IString_partContext:
			return v.VisitString_part(tt)
		default:
			v.Reporter.NotImplemented(reflect.TypeOf(tt))
		}
	}

	// return v.VisitChildren(ctx)
	return ""
}

func (v *WdlVisitor) VisitPrimitive_literal(ctx domain.IPrimitive_literalContext) interface{} {

	// TODO: figure out the appropriate return type.
	return 123
}

func (v *WdlVisitor) VisitExpr(ctx domain.IExprContext) domain.IExpression {
	infixCtx := ctx.GetChild(0).(domain.IExpr_infixContext)

	// 	expr
	// 	: expr_infix
	// 	;
	//
	//   expr_infix
	// 	: expr_infix0 #infix0
	// 	;

	infix0Ctx := infixCtx.GetChild(0).(domain.IExpr_infix0Context)
	return v.VisitInfix0(infix0Ctx)
}

func (v *WdlVisitor) VisitInfix0(ctx domain.IExpr_infix0Context) domain.IExpression {

	// expr_infix0
	// : expr_infix0 OR expr_infix1 #lor
	// | expr_infix1 #infix1
	// ;

	if count := ctx.GetChildCount(); count == 1 {
		// No branches
		infix1Ctx := ctx.GetChild(0).(domain.IExpr_infix1Context)
		return v.VisitInfix1(infix1Ctx)
	} else if count != 3 {
		v.Reporter.Error(ctx, fmt.Errorf("infix0 contains unexpected number of children"))
		return domain.NewUnknownExpr()
	}

	infix0Ctx := ctx.GetChild(0).(domain.IExpr_infix0Context) // 0
	infix1Ctx := ctx.GetChild(2).(domain.IExpr_infix1Context) // 2

	return domain.NewLorExpr(
		v.VisitInfix0(infix0Ctx), // left
		v.VisitInfix1(infix1Ctx), // right
	)
}

func (v *WdlVisitor) VisitInfix1(ctx domain.IExpr_infix1Context) domain.IExpression {

	// expr_infix1
	// : expr_infix1 AND expr_infix2 #land
	// | expr_infix2 #infix2
	// ;

	if count := ctx.GetChildCount(); count == 1 {
		// No branches
		infix2Ctx := ctx.GetChild(0).(domain.IExpr_infix2Context)
		return v.VisitInfix2(infix2Ctx)
	} else if count != 3 {
		v.Reporter.Error(ctx, fmt.Errorf("infix1 contains unexpected number of children"))
		return domain.NewUnknownExpr()
	}

	infix1Ctx := ctx.GetChild(0).(domain.IExpr_infix1Context) // 0
	infix2Ctx := ctx.GetChild(2).(domain.IExpr_infix2Context) // 2

	return domain.NewLandExpr(
		v.VisitInfix1(infix1Ctx), // left
		v.VisitInfix2(infix2Ctx), // right
	)
}

func (v *WdlVisitor) VisitInfix2(ctx domain.IExpr_infix2Context) domain.IExpression {

	// A bunch of comparisons:

	// expr_infix2
	// : expr_infix2 EQUALITY expr_infix3 #eqeq
	// | expr_infix2 NOTEQUAL expr_infix3 #neq
	// | expr_infix2 LTE expr_infix3 #lte
	// | expr_infix2 GTE expr_infix3 #gte
	// | expr_infix2 LT expr_infix3 #lt
	// | expr_infix2 GT expr_infix3 #gt
	// | expr_infix3 #infix3
	// ;

	if count := ctx.GetChildCount(); count == 1 {
		// No branches
		infix3Ctx := ctx.GetChild(0).(domain.IExpr_infix3Context)
		return v.VisitInfix3(infix3Ctx)
	} else if count != 3 {
		v.Reporter.Error(ctx, fmt.Errorf("infix2 contains unexpected number of children"))
		return domain.NewUnknownExpr()
	}

	infix2Ctx := ctx.GetChild(0).(domain.IExpr_infix2Context) // 0
	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()      // 1
	infix3Ctx := ctx.GetChild(2).(domain.IExpr_infix3Context) // 2

	// `op` is one of []string{"==", "!=", "<=", ">=", "<", ">"}.  We could further parse this, but string would be fine for now.

	return domain.NewComparisonExpr(
		v.VisitInfix2(infix2Ctx), // left
		op,                       // operation
		v.VisitInfix3(infix3Ctx), // right
	)
}

func (v *WdlVisitor) VisitInfix3(ctx domain.IExpr_infix3Context) domain.IExpression {

	// addition and subtraction

	// expr_infix3
	// : expr_infix3 PLUS expr_infix4 #add
	// | expr_infix3 MINUS expr_infix4 #sub
	// | expr_infix4 #infix4
	// ;

	if count := ctx.GetChildCount(); count == 1 {
		// No branches
		infix4Ctx := ctx.GetChild(0).(domain.IExpr_infix4Context)
		return v.VisitInfix4(infix4Ctx)
	} else if count != 3 {
		v.Reporter.Error(ctx, fmt.Errorf("infix3 contains unexpected number of children"))
		return domain.NewUnknownExpr()
	}

	infix3Ctx := ctx.GetChild(0).(domain.IExpr_infix3Context) // 0
	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()      // 1
	infix4Ctx := ctx.GetChild(2).(domain.IExpr_infix4Context) // 2

	// `op` is one of []string{"+", "-"}.  We could further parse this, but string would be fine for now.

	return domain.NewBinaryOpExpr(
		v.VisitInfix3(infix3Ctx), // left
		op,                       // operation
		v.VisitInfix4(infix4Ctx), // right
	)
}

func (v *WdlVisitor) VisitInfix4(ctx domain.IExpr_infix4Context) domain.IExpression {

	// multiplication, division, and mod

	// expr_infix4
	// : expr_infix4 STAR expr_infix5 #mul
	// | expr_infix4 DIVIDE expr_infix5 #divide
	// | expr_infix4 MOD expr_infix5 #mod
	// | expr_infix5 #infix5
	// ;

	if count := ctx.GetChildCount(); count == 1 {
		// No branches
		infix5Ctx := ctx.GetChild(0).(domain.IExpr_infix5Context)
		return v.VisitInfix5(infix5Ctx)
	} else if count != 3 {
		v.Reporter.Error(ctx, fmt.Errorf("infix4 contains unexpected number of children"))
		return domain.NewUnknownExpr()
	}

	infix4Ctx := ctx.GetChild(0).(domain.IExpr_infix4Context) // 0
	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()      // 1
	infix5Ctx := ctx.GetChild(2).(domain.IExpr_infix5Context) // 2

	// `op` is one of []string{"*", "/", "%"}.  We could further parse this, but string would be fine for now.

	return domain.NewBinaryOpExpr(
		v.VisitInfix4(infix4Ctx), // left
		op,                       // operation
		v.VisitInfix5(infix5Ctx), // right
	)
}

func (v *WdlVisitor) VisitInfix5(ctx domain.IExpr_infix5Context) domain.IExpression {
	coreCtx := ctx.GetChild(0).(domain.IExpr_coreContext)

	// expr_infix5
	// : expr_core
	// ;
	//
	// expr_core
	// : Identifier LPAREN (expr (COMMA expr)* COMMA?)? RPAREN #apply
	// | LBRACK (expr (COMMA expr)* COMMA?)* RBRACK #array_literal
	// | LPAREN expr COMMA expr RPAREN #pair_literal
	// | LBRACE (expr COLON expr (COMMA expr COLON expr)* COMMA?)* RBRACE #map_literal
	// | OBJECT_LITERAL LBRACE (Identifier COLON expr (COMMA Identifier COLON expr)* COMMA?)* RBRACE #object_literal
	// | IF expr THEN expr ELSE expr #ifthenelse
	// | LPAREN expr RPAREN #expression_group
	// | expr_core LBRACK expr RBRACK #at
	// | expr_core DOT Identifier #get_name
	// | NOT expr #negate
	// | (PLUS | MINUS) expr #unarysigned
	// | primitive_literal #primitives
	// | Identifier #left_name
	// ;

	// We probably **have** to use reflection here to get the type of context we're dealing with (the label). I don't
	// think ANTLR4 stores this information in the node, it just assumes that we're working with the original interface
	// so we could type check that way, but we're not.
	fmt.Printf("%v\n", reflect.TypeOf(coreCtx).Elem().Name())

	switch reflect.TypeOf(coreCtx).Elem().Name() {
	case "PrimitivesContext":
		rv := domain.NewTerminalExpr()
		rv.Value = v.VisitPrimitive_literal(coreCtx.GetChild(0).(domain.IPrimitive_literalContext))
		return rv
	}

	return domain.NewUnknownExpr()
}

// func (v *WdlVisitor) VisitExpr_infix5(ctx *parsers.Expr_infix5Context) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitPair_literal(ctx *parsers.Pair_literalContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitUnarysigned(ctx *parsers.UnarysignedContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitApply(ctx *parsers.ApplyContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitExpression_group(ctx *parsers.Expression_groupContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitPrimitives(ctx *parsers.PrimitivesContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitLeft_name(ctx *parsers.Left_nameContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitAt(ctx *parsers.AtContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitNegate(ctx *parsers.NegateContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitMap_literal(ctx *parsers.Map_literalContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitIfthenelse(ctx *parsers.IfthenelseContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitGet_name(ctx *parsers.Get_nameContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitObject_literal(ctx *parsers.Object_literalContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitArray_literal(ctx *parsers.Array_literalContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// VisitVersion returns the version as a string. This is a terminal operation.
func (v *WdlVisitor) VisitVersion(ctx domain.IVersionContext) string {
	version := ctx.ReleaseVersion().GetText()

	if version != v.Version {
		v.Reporter.Warn(ctx, fmt.Sprintf("Version mismatch! Using parser for version %s but workflow has version %s", v.Version, version))
	}

	return version
}

// VisitImport_alias returns the import alias as a mapping from the original
// identifer -> the aliased identifier. This is a terminal operation.
func (v *WdlVisitor) VisitImport_alias(ctx domain.IImport_aliasContext) (domain.Identifier, domain.Identifier) {
	return domain.Identifier(ctx.Identifier(0).GetText()),
		domain.Identifier(ctx.Identifier(1).GetText())
}

// VisitImport_as returns the identifier that refers to the import. This is a terminal operation.
func (v *WdlVisitor) VisitImport_as(ctx domain.IImport_asContext) domain.Identifier {
	return domain.Identifier(ctx.Identifier().GetText())
}

func (v *WdlVisitor) VisitImport_doc(ctx domain.IImport_docContext) domain.Import {
	var url string
	var as domain.Identifier                             // optional
	aliases := map[domain.Identifier]domain.Identifier{} // optional

	for _, child := range ctx.GetChildren() {
		switch tt := child.(type) {
		case domain.IStringContext:
			url = v.VisitString(tt)
		case domain.IImport_asContext:
			as = v.VisitImport_as(tt)
		case domain.IImport_aliasContext:
			original, alias := v.VisitImport_alias(tt)
			aliases[original] = alias
		default:
			v.Reporter.NotImplemented(reflect.TypeOf(tt))
		}
	}

	var absoluteUrl string
	if strings.Contains(url, "://") {
		absoluteUrl = url
	}

	u, err := urlparse.Parse(v.Url)
	if err == nil {
		u.Path = path.Join(path.Dir(u.Path), url)
		absoluteUrl = u.String()
	}

	return domain.Import{
		Url:         url,
		AbsoluteUrl: absoluteUrl,
		As:          as,
		Aliases:     aliases,
	}
}

// func (v *WdlVisitor) VisitStruct(ctx *parsers.StructContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitMeta_value(ctx *parsers.Meta_valueContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitMeta_string_part(ctx *parsers.Meta_string_partContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitMeta_string(ctx *parsers.Meta_stringContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitMeta_array(ctx *parsers.Meta_arrayContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitMeta_object(ctx *parsers.Meta_objectContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitMeta_object_kv(ctx *parsers.Meta_object_kvContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitMeta_kv(ctx *parsers.Meta_kvContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitParameter_meta(ctx *parsers.Parameter_metaContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitMeta(ctx *parsers.MetaContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitTask_runtime_kv(ctx *parsers.Task_runtime_kvContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitTask_runtime(ctx *parsers.Task_runtimeContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitTask_input(ctx *parsers.Task_inputContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitTask_output(ctx *parsers.Task_outputContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitTask_command_string_part(ctx *parsers.Task_command_string_partContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitTask_command_expr_part(ctx *parsers.Task_command_expr_partContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitTask_command_expr_with_string(ctx *parsers.Task_command_expr_with_stringContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitTask_command(ctx *parsers.Task_commandContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitTask_element(ctx *parsers.Task_elementContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitTask(ctx *parsers.TaskContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

func (v *WdlVisitor) VisitInner_workflow_element(ctx domain.IInner_workflow_elementContext) interface{} {
	return nil
}

// func (v *WdlVisitor) VisitCall_alias(ctx *parsers.Call_aliasContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitCall_input(ctx *parsers.Call_inputContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitCall_inputs(ctx *parsers.Call_inputsContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitCall_body(ctx *parsers.Call_bodyContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitCall_name(ctx *parsers.Call_nameContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitCall(ctx *parsers.CallContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitScatter(ctx *parsers.ScatterContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitConditional(ctx *parsers.ConditionalContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

func (v *WdlVisitor) VisitWorkflow_input(ctx domain.IWorkflow_inputContext) []domain.Declaration {
	len := 0
	for _, ctx := range ctx.GetChildren() {
		if _, ok := ctx.(domain.IAny_declsContext); ok {
			len++
		}
	}

	inputs := make([]domain.Declaration, len)
	idx := 0

	for _, child := range ctx.GetChildren() {
		switch tt := child.(type) {
		case domain.IAny_declsContext:
			inputs[idx] = v.VisitAny_decls(tt)
			idx++
		}
	}

	return inputs
}

func (v *WdlVisitor) VisitWorkflow_output(ctx domain.IWorkflow_outputContext) interface{} {
	return nil
}

// func (v *WdlVisitor) VisitInput(ctx *parsers.InputContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitOutput(ctx *parsers.OutputContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitInner_element(ctx *parsers.Inner_elementContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitParameter_meta_element(ctx *parsers.Parameter_meta_elementContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitMeta_element(ctx *parsers.Meta_elementContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

func (v *WdlVisitor) VisitWorkflow(ctx domain.IWorkflowContext) *domain.Workflow {
	name := domain.Identifier(ctx.Identifier().GetText())

	var inputs []domain.Declaration

	for _, child := range ctx.GetChildren() {
		switch tt := child.(type) {
		case domain.IWorkflow_elementContext:
			// Let's flatten this.
			// TODO: we need to support expression and declaration first.
			switch tt := tt.GetChild(0).(type) {
			case domain.IWorkflow_inputContext:
				inputs = v.VisitWorkflow_input(tt)
			case domain.IInner_workflow_elementContext:
				v.Reporter.Warn(ctx, "IInner_workflow_elementContext")
			case domain.IWorkflow_outputContext:
				v.Reporter.Warn(ctx, "IWorkflow_outputContext")
			default:
				v.Reporter.NotImplemented(reflect.TypeOf(tt))
			}
		}
	}

	return &domain.Workflow{
		Name:   name,
		Inputs: inputs,
	}
}

func (v *WdlVisitor) VisitDocument_element(ctx domain.IDocument_elementContext) interface{} {
	// We've flattened this structure.
	v.Reporter.Warn(ctx, "VisitDocument_element should not be called.")
	return nil
}

func (v *WdlVisitor) VisitDocument(ctx domain.IDocumentContext) *domain.Document {
	var version string
	var workflow *domain.Workflow // optional
	var imports []domain.Import

	for _, ctx := range ctx.GetChildren() {
		switch tt := ctx.(type) {
		// There should be exactly one version statement and one or zero workflows.
		case domain.IVersionContext:
			version = v.VisitVersion(tt)
		case domain.IWorkflowContext:
			workflow = v.VisitWorkflow(tt)
		// There could be a number of document elements.
		case domain.IDocument_elementContext:
			// Let's flatten this.
			switch tt := tt.GetChild(0).(type) {
			case domain.IImport_docContext:
				imports = append(imports, v.VisitImport_doc(tt))
			// TODO: tasks
			// TODO: structs
			default:
				v.Reporter.NotImplemented(reflect.TypeOf(tt))
			}
		default:
			v.Reporter.NotImplemented(reflect.TypeOf(tt))
		}
	}

	return &domain.Document{
		Url:      v.Url,
		Version:  version,
		Workflow: workflow,
		Imports:  imports,
	}
}
