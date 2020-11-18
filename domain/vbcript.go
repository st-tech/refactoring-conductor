package domain

// VBScriptControlFlowPattern defines the VBScript control flow pattern.
const (
	VBScriptIfPattern          = `(?i)(If)`
	VBScriptElsePattern        = `(?i)(Else|ElseIf)`
	VBScriptEndIfPattern       = `(?i)(End if)`
	VBScriptFunctionPattern    = `(?i)(Function)`
	VBScriptEndFunctionPattern = `(?i)(End Function)`
	VBScriptProcedurePattern   = `[\r\n](\s*?)(private|public)?\s*(function|sub|property +(get|set|let))\s\[?([^(\r\n\]]+).*?[\r\n]([\s\S]*?)end +(function|sub|Property)`
)

// VBScript is to set VBScript Code Architecture.
type VBScript struct {
	NestState           int
	CognitiveComplexity int
	Maintainability     int
	Functions           []FunctionState
}

// FunctionState is to set function states.
type FunctionState struct {
	Name                string
	NestState           int
	CognitiveComplexity int
}
