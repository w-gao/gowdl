package internal

import (
	"fmt"
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

type ContainsIdentifierContext interface {
	antlr.ParserRuleContext
	Identifier() antlr.TerminalNode
}

type WdlVisitor struct {
	// Let's move away from the generated interface which is ambiguous with
	// typing. This way we'd lose some functionality such as tree.Accept(v),
	// since the visitor no longer implements the antlr.ParseTreeVisitor
	// interface, but we mitigate that with the Visit() function.
	// parsers.BaseWdlV1ParserVisitor

	Version  string
	Reporter IVisitorReporter
}

func NewWdlVisitor(version string) *WdlVisitor {
	reporter := &impl.FmtReporter{}
	return &WdlVisitor{
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

// VisitIdentifier is a special visitor that returns the identifier that is under the given context.
func (v *WdlVisitor) VisitIdentifier(ctx ContainsIdentifierContext) domain.Identifier {
	return domain.Identifier(ctx.Identifier().GetText())
}

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

// func (v *WdlVisitor) VisitWdl_type(ctx *parsers.Wdl_typeContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitUnbound_decls(ctx *parsers.Unbound_declsContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitBound_decls(ctx *parsers.Bound_declsContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitAny_decls(ctx *parsers.Any_declsContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

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

// func (v *WdlVisitor) VisitPrimitive_literal(ctx *parsers.Primitive_literalContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitExpr(ctx *parsers.ExprContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitInfix0(ctx *parsers.Infix0Context) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitInfix1(ctx *parsers.Infix1Context) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitLor(ctx *parsers.LorContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitInfix2(ctx *parsers.Infix2Context) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitLand(ctx *parsers.LandContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitEqeq(ctx *parsers.EqeqContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitLt(ctx *parsers.LtContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitInfix3(ctx *parsers.Infix3Context) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitGte(ctx *parsers.GteContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitNeq(ctx *parsers.NeqContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitLte(ctx *parsers.LteContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitGt(ctx *parsers.GtContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitAdd(ctx *parsers.AddContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitSub(ctx *parsers.SubContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitInfix4(ctx *parsers.Infix4Context) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitMod(ctx *parsers.ModContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitMul(ctx *parsers.MulContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitDivide(ctx *parsers.DivideContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitInfix5(ctx *parsers.Infix5Context) interface{} {
// 	return v.VisitChildren(ctx)
// }

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

func (v *WdlVisitor) VisitImport_alias(ctx domain.IImport_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *WdlVisitor) VisitImport_as(ctx domain.IImport_asContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *WdlVisitor) VisitImport_doc(ctx domain.IImport_docContext) domain.Import {
	var url string
	var as domain.Identifier // optional
	// var aliases []domain.ImportAlias

	for _, child := range ctx.GetChildren() {
		switch tt := child.(type) {
		case domain.IStringContext:
			url = v.VisitString(tt)
		case domain.IImport_asContext:
			as = v.VisitIdentifier(tt)
		default:
			v.Reporter.NotImplemented(reflect.TypeOf(tt))
		}
	}

	return domain.Import{
		Url: url,
		As:  as,
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

// func (v *WdlVisitor) VisitInner_workflow_element(ctx *parsers.Inner_workflow_elementContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

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

// func (v *WdlVisitor) VisitWorkflow_input(ctx *parsers.Workflow_inputContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlVisitor) VisitWorkflow_output(ctx *parsers.Workflow_outputContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

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
	// return nil
	return &domain.Workflow{}
}

func (v *WdlVisitor) VisitDocument_element(ctx domain.IDocument_elementContext) interface{} {
	v.Reporter.Warn(ctx, "VisitDocument_element is deprecated")
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
		Version:  version,
		Workflow: workflow,
		Imports:  imports,
	}
}
