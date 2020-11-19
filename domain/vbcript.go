package domain

// VBScriptControlFlowPattern defines the VBScript control flow pattern.
const (
	VBScriptIfPattern          = `(?i)(If )`
	VBScriptElsePattern        = `(?i)(Else|ElseIf)`
	VBScriptEndIfPattern       = `(?i)(End if)`
	VBScriptFunctionPattern    = `(?i)(Function)`
	VBScriptEndFunctionPattern = `(?i)(End Function)`
	VBScriptProcedurePattern   = `[\r\n](\s*?)(private|public)?\s*(function|sub|property +(get|set|let))\s\[?([^(\r\n\]]+).*?[\r\n]([\s\S]*?)end +(function|sub|Property)`
)

// VBScriptFileExtension defines the VBSCript file extensions.
const VBScriptFileExtension = `(?i)(.wsf|.vbs|.asp|.inc)$`

// VBScript is to set VBScript Code Architecture.
type VBScript struct {
	NestState           int
	CognitiveComplexity int
	Maintainability     int
	IsBeginFunction     bool
	Functions           []Function
}

// Function is to set function state to VBSCript.
type Function struct {
	NestState           int
	CognitiveComplexity int
	FunctionName        string
}
