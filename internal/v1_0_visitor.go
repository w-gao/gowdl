package internal

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/w-gao/gowdl/parsers"
)

type WdlV1Visitor struct {
	// Move away from the generated interface which is ambiguous with typing.
	// parsers.BaseWdlV1ParserVisitor

}

func NewWdlV1Visitor() *WdlV1Visitor {
	return &WdlV1Visitor{}
}

// Visit takes in any parse tree and visits the child nodes from there.
func (v *WdlV1Visitor) Visit(tree antlr.ParseTree) interface{} {
	switch t := tree.(type) {
	case *parsers.DocumentContext:
		return v.VisitDocument(t)
	default:
		return nil
	}
}

func (v *WdlV1Visitor) VisitChildren(node antlr.RuleNode) interface{}     { return nil }
func (v *WdlV1Visitor) VisitTerminal(node antlr.TerminalNode) interface{} { return nil }
func (v *WdlV1Visitor) VisitErrorNode(node antlr.ErrorNode) interface{}   { return nil }

// func (v *WdlV1Visitor) VisitMap_type(ctx *parsers.Map_typeContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitArray_type(ctx *parsers.Array_typeContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitPair_type(ctx *parsers.Pair_typeContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitType_base(ctx *parsers.Type_baseContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitWdl_type(ctx *parsers.Wdl_typeContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitUnbound_decls(ctx *parsers.Unbound_declsContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitBound_decls(ctx *parsers.Bound_declsContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitAny_decls(ctx *parsers.Any_declsContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitNumber(ctx *parsers.NumberContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitExpression_placeholder_option(ctx *parsers.Expression_placeholder_optionContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

func (v *WdlV1Visitor) VisitString_part(ctx *parsers.String_partContext) string {
	var parts []string
	for _, part := range ctx.AllStringPart() {
		parts = append(parts, part.GetText())
	}

	return strings.Join(parts, "")
}

// func (v *WdlV1Visitor) VisitString_expr_part(ctx *parsers.String_expr_partContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitString_expr_with_string_part(ctx *parsers.String_expr_with_string_partContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

func (v *WdlV1Visitor) VisitString(ctx *parsers.StringContext) string { // interface{} {
	// TODO: parse actual string, which can contain expressions!
	for _, ctx := range ctx.GetChildren() {
		switch tt := ctx.(type) {
		case *parsers.String_partContext:
			return v.VisitString_part(tt)
		default:
			fmt.Printf("WARN: NotImplemented: %v\n", reflect.TypeOf(tt))
		}
	}

	// return v.VisitChildren(ctx)
	return ""
}

// func (v *WdlV1Visitor) VisitPrimitive_literal(ctx *parsers.Primitive_literalContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitExpr(ctx *parsers.ExprContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitInfix0(ctx *parsers.Infix0Context) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitInfix1(ctx *parsers.Infix1Context) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitLor(ctx *parsers.LorContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitInfix2(ctx *parsers.Infix2Context) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitLand(ctx *parsers.LandContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitEqeq(ctx *parsers.EqeqContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitLt(ctx *parsers.LtContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitInfix3(ctx *parsers.Infix3Context) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitGte(ctx *parsers.GteContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitNeq(ctx *parsers.NeqContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitLte(ctx *parsers.LteContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitGt(ctx *parsers.GtContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitAdd(ctx *parsers.AddContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitSub(ctx *parsers.SubContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitInfix4(ctx *parsers.Infix4Context) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitMod(ctx *parsers.ModContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitMul(ctx *parsers.MulContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitDivide(ctx *parsers.DivideContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitInfix5(ctx *parsers.Infix5Context) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitExpr_infix5(ctx *parsers.Expr_infix5Context) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitPair_literal(ctx *parsers.Pair_literalContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitUnarysigned(ctx *parsers.UnarysignedContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitApply(ctx *parsers.ApplyContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitExpression_group(ctx *parsers.Expression_groupContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitPrimitives(ctx *parsers.PrimitivesContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitLeft_name(ctx *parsers.Left_nameContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitAt(ctx *parsers.AtContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitNegate(ctx *parsers.NegateContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitMap_literal(ctx *parsers.Map_literalContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitIfthenelse(ctx *parsers.IfthenelseContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitGet_name(ctx *parsers.Get_nameContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitObject_literal(ctx *parsers.Object_literalContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitArray_literal(ctx *parsers.Array_literalContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// VisitVersion returns the version as a string. This is a terminal operation.
func (v *WdlV1Visitor) VisitVersion(ctx *parsers.VersionContext) string {
	return ctx.ReleaseVersion().GetText()
}

func (v *WdlV1Visitor) VisitImport_alias(ctx *parsers.Import_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *WdlV1Visitor) VisitImport_as(ctx *parsers.Import_asContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *WdlV1Visitor) VisitImport_doc(ctx *parsers.Import_docContext) Import {
	var url string
	var as string // optional
	// var aliases []lib.ImportAlias

	for _, ctx := range ctx.GetChildren() {
		switch tt := ctx.(type) {
		case *parsers.StringContext:
			url = v.VisitString(tt)
		case *parsers.Import_asContext:
			as = tt.Identifier().GetText()
		default:
			fmt.Printf("WARN: NotImplemented: %v\n", reflect.TypeOf(tt))
		}
	}

	return Import{
		Url: url,
		As:  as,
	}
}

// func (v *WdlV1Visitor) VisitStruct(ctx *parsers.StructContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitMeta_value(ctx *parsers.Meta_valueContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitMeta_string_part(ctx *parsers.Meta_string_partContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitMeta_string(ctx *parsers.Meta_stringContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitMeta_array(ctx *parsers.Meta_arrayContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitMeta_object(ctx *parsers.Meta_objectContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitMeta_object_kv(ctx *parsers.Meta_object_kvContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitMeta_kv(ctx *parsers.Meta_kvContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitParameter_meta(ctx *parsers.Parameter_metaContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitMeta(ctx *parsers.MetaContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitTask_runtime_kv(ctx *parsers.Task_runtime_kvContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitTask_runtime(ctx *parsers.Task_runtimeContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitTask_input(ctx *parsers.Task_inputContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitTask_output(ctx *parsers.Task_outputContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitTask_command_string_part(ctx *parsers.Task_command_string_partContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitTask_command_expr_part(ctx *parsers.Task_command_expr_partContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitTask_command_expr_with_string(ctx *parsers.Task_command_expr_with_stringContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitTask_command(ctx *parsers.Task_commandContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitTask_element(ctx *parsers.Task_elementContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitTask(ctx *parsers.TaskContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitInner_workflow_element(ctx *parsers.Inner_workflow_elementContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitCall_alias(ctx *parsers.Call_aliasContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitCall_input(ctx *parsers.Call_inputContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitCall_inputs(ctx *parsers.Call_inputsContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitCall_body(ctx *parsers.Call_bodyContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitCall_name(ctx *parsers.Call_nameContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitCall(ctx *parsers.CallContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitScatter(ctx *parsers.ScatterContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitConditional(ctx *parsers.ConditionalContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitWorkflow_input(ctx *parsers.Workflow_inputContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitWorkflow_output(ctx *parsers.Workflow_outputContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitInput(ctx *parsers.InputContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitOutput(ctx *parsers.OutputContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitInner_element(ctx *parsers.Inner_elementContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitParameter_meta_element(ctx *parsers.Parameter_meta_elementContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *WdlV1Visitor) VisitMeta_element(ctx *parsers.Meta_elementContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

func (v *WdlV1Visitor) VisitWorkflow(ctx *parsers.WorkflowContext) *Workflow {
	return nil
}

func (v *WdlV1Visitor) VisitDocument_element(ctx *parsers.Document_elementContext) interface{} {
	fmt.Println("WARN: VisitDocument_element is deprecated.")
	return nil
}

func (v *WdlV1Visitor) VisitDocument(ctx *parsers.DocumentContext) *Document {
	var version string
	var workflow *Workflow // optional
	var imports []Import

	for _, ctx := range ctx.GetChildren() {
		switch tt := ctx.(type) {
		case *parsers.VersionContext:
			// The grammar enforces that there should be one and only one version statement
			version = v.VisitVersion(tt)
		case *parsers.WorkflowContext:
			workflow = v.VisitWorkflow(tt)
		case *parsers.Document_elementContext:
			// Let's flatten this.
			switch tt := tt.GetChild(0).(type) {
			case *parsers.Import_docContext:
				imports = append(imports, v.VisitImport_doc(tt))
			// TODO: tasks and structs
			default:
				fmt.Printf("WARN: NotImplemented: %v\n", reflect.TypeOf(tt))
			}

		default:
			fmt.Printf("WARN: NotImplemented: %v\n", reflect.TypeOf(tt))
		}
	}

	return &Document{
		Version:  version,
		Workflow: workflow,
		Imports:  imports,
	}
}
