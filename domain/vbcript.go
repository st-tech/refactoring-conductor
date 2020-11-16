package domain

// VBScriptControlFlowPattern defines the VBScript control flow pattern.
const VBScriptControlFlowPattern = `(?i)(If|For|Else|End\sif)`
const VBScriptProcedurePattern = `[\r\n](\s*?)(private|public)?\s*(function|sub|property +(get|set|let))\s\[?([^(\r\n\]]+).*?[\r\n]([\s\S]*?)end +(function|sub|Property)`
const VBScriptOperatorPattern = ``

// VBScript is to set VBScript Code Architecture.
type VBScript struct {
	NestState           int
	CognitiveComplexity int
}
