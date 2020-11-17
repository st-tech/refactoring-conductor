package domain

// VBScriptControlFlowPattern defines the VBScript control flow pattern.
const VBScriptIfPattern = `(?i)(If)`
const VBScriptElsePattern = `(?i)(Else|ElseIf)`
const VBScriptEndIfPattern = `(?i)(End if)`
const VBScriptProcedurePattern = `[\r\n](\s*?)(private|public)?\s*(function|sub|property +(get|set|let))\s\[?([^(\r\n\]]+).*?[\r\n]([\s\S]*?)end +(function|sub|Property)`
const VBScriptOperatorPattern = ``

// VBScript is to set VBScript Code Architecture.
type VBScript struct {
	NestState           int
	CognitiveComplexity int
	Maintainability     int
}
