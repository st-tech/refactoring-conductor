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

type VBScriptJson struct {
	VBScript []VBScript
}

// VBScript is to set VBScript Code Architecture.
type VBScript struct {
	NestState           int  `json:"nest_state"`
	CognitiveComplexity int  `json:"cognitive_complexity"`
	IsBeginFunction     bool `json:"is_begin_function"`
	Functions           []Function
}

// Function is to set function state to VBSCript.
type Function struct {
	NestState           int    `json:"nest_state"`
	CognitiveComplexity int    `json:"cognitive_complexity"`
	FunctionName        string `json:"function_name"`
}
