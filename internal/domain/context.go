package domain

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Interfaces to support dynamic dispatch for different WDL versions.

// Limitation: we can not dynamically support functions that return another context (or array of
// contexts) Instead, the user should loop through ctx.GetChildren() and cast to child contexts
// there. This is the more efficient way too because it will only loop through the children once.

type IMap_typeContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsMap_typeContext()
}

type IArray_typeContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsArray_typeContext()
}

type IPair_typeContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsPair_typeContext()
}

type IType_baseContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsType_baseContext()
}

type IWdl_typeContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsWdl_typeContext()
}

type IUnbound_declsContext interface {
	antlr.ParserRuleContext

	Identifier() antlr.TerminalNode

	GetParser() antlr.Parser
	IsUnbound_declsContext()
}

type IBound_declsContext interface {
	antlr.ParserRuleContext

	Identifier() antlr.TerminalNode

	GetParser() antlr.Parser
	IsBound_declsContext()
}

type IAny_declsContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsAny_declsContext()
}

type INumberContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsNumberContext()
}

type IExpression_placeholder_optionContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsExpression_placeholder_optionContext()
}

type IString_partContext interface {
	antlr.ParserRuleContext

	AllStringPart() []antlr.TerminalNode
	StringPart(i int) antlr.TerminalNode

	GetParser() antlr.Parser
	IsString_partContext()
}

type IString_expr_partContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsString_expr_partContext()
}

type IString_expr_with_string_partContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsString_expr_with_string_partContext()
}

type IStringContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsStringContext()
}

type IPrimitive_literalContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsPrimitive_literalContext()
}

type IExprContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsExprContext()
}

type IExpr_infixContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsExpr_infixContext()
}

type IExpr_infix0Context interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsExpr_infix0Context()
}

type IExpr_infix1Context interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsExpr_infix1Context()
}

type IExpr_infix2Context interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsExpr_infix2Context()
}

type IExpr_infix3Context interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsExpr_infix3Context()
}

type IExpr_infix4Context interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsExpr_infix4Context()
}

type IExpr_infix5Context interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsExpr_infix5Context()
}

type IExpr_coreContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsExpr_coreContext()
}

type IVersionContext interface {
	antlr.ParserRuleContext

	ReleaseVersion() antlr.TerminalNode

	GetParser() antlr.Parser
	IsVersionContext()
}

type IImport_aliasContext interface {
	antlr.ParserRuleContext

	AllIdentifier() []antlr.TerminalNode
	Identifier(i int) antlr.TerminalNode

	GetParser() antlr.Parser
	IsImport_aliasContext()
}

type IImport_asContext interface {
	antlr.ParserRuleContext

	Identifier() antlr.TerminalNode

	GetParser() antlr.Parser
	IsImport_asContext()
}

type IImport_docContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsImport_docContext()
}

type IStructContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsStructContext()
}

type IMeta_valueContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsMeta_valueContext()
}

type IMeta_string_partContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsMeta_string_partContext()
}

type IMeta_stringContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsMeta_stringContext()
}

type IMeta_arrayContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsMeta_arrayContext()
}

type IMeta_objectContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsMeta_objectContext()
}

type IMeta_object_kvContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsMeta_object_kvContext()
}

type IMeta_kvContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsMeta_kvContext()
}

type IParameter_metaContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsParameter_metaContext()
}

type IMetaContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsMetaContext()
}

type ITask_runtime_kvContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsTask_runtime_kvContext()
}

type ITask_runtimeContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsTask_runtimeContext()
}

type ITask_inputContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsTask_inputContext()
}

type ITask_outputContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsTask_outputContext()
}

type ITask_command_string_partContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsTask_command_string_partContext()
}

type ITask_command_expr_partContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsTask_command_expr_partContext()
}

type ITask_command_expr_with_stringContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsTask_command_expr_with_stringContext()
}

type ITask_commandContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsTask_commandContext()
}

type ITask_elementContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsTask_elementContext()
}

type ITaskContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsTaskContext()
}

type IInner_workflow_elementContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsInner_workflow_elementContext()
}

type ICall_aliasContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsCall_aliasContext()
}

type ICall_inputContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsCall_inputContext()
}

type ICall_inputsContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsCall_inputsContext()
}

type ICall_bodyContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsCall_bodyContext()
}

type ICall_nameContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsCall_nameContext()
}

type ICallContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsCallContext()
}

type IScatterContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsScatterContext()
}

type IConditionalContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsConditionalContext()
}

type IWorkflow_inputContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsWorkflow_inputContext()
}

type IWorkflow_outputContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsWorkflow_outputContext()
}

type IWorkflow_elementContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsWorkflow_elementContext()
}

type IWorkflowContext interface {
	antlr.ParserRuleContext

	Identifier() antlr.TerminalNode

	GetParser() antlr.Parser
	IsWorkflowContext()
}

type IDocument_elementContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsDocument_elementContext()
}

type IDocumentContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser
	IsDocumentContext()
}
