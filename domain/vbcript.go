package domain

// VBScriptControlFlowPattern defines the VBScript control flow pattern.
const (
	VBScriptIfPattern          = `(?i)^(\s*?)(\t*?)(If )`
	VBScriptElsePattern        = `(?i)^(\s*?)(\t*?)(Else|ElseIf)`
	VBScriptEndIfPattern       = `(?i)^(\s*?)(\t*?)(End if)`
	VBScriptFunctionPattern    = `(?i)^(\s*?)(\t*?)(Function)`
	VBScriptEndFunctionPattern = `(?i)^(\s*?)(\t*?)(End Function)`
	VBScriptProcedurePattern   = `[\r\n](\s*?)(private|public)?\s*(function|sub|property +(get|set|let))\s\[?([^(\r\n\]]+).*?[\r\n]([\s\S]*?)end +(function|sub|Property)`
)

// VBScriptFileExtension defines the VBSCript file extensions.
const VBScriptFileExtension = `(?i)(.wsf|.vbs|.asp|.inc)$`

// VBScriptJSON defines the VBScript JSON.
type VBScriptJSON struct {
	VBScript []VBScript
}

// VBScript is to set VBScript Code Architecture.
type VBScript struct {
	FileName            string `json:"file_name"`
	NestState           int    `json:"-"`
	CognitiveComplexity int    `json:"cognitive_complexity"`
	IsBeginFunction     bool   `json:"-"`
	Functions           []Function
}

// Function is to set function state to VBSCript.
type Function struct {
	NestState           int    `json:"-"`
	CognitiveComplexity int    `json:"cognitive_complexity"`
	FunctionName        string `json:"function_name"`
}
