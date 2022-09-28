package domain

// VBScriptControlFlowPattern defines the VBScript control flow pattern.
const (
	VBScriptIfPattern          = `(?i)^(\s*?)(\t*?)(If )`
	VBScriptElsePattern        = `(?i)^(\s*?)(\t*?)(Else|ElseIf)`
	VBScriptEndIfPattern       = `(?i)^(\s*?)(\t*?)(End if)`
	VBScriptForPattern         = `(?i)^(\s*?)(\t*?)(For )`
	VBScriptNextPattern        = `(?i)^(\s*?)(\t*?)(Next)($|(\s+?)|(\t+?))`
	VBScriptDoPattern          = `(?i)^(\s*?)(\t*?)(Do)($|(\s+?)|(\t+?)| While | Until )`
	VBScriptLoopPattern        = `(?i)^(\s*?)(\t*?)(Loop)($|(\s+?)|(\t+?)| While | Until )`
	VBScriptWhilePattern       = `(?i)^(\s*?)(\t*?)(While )`
	VBScriptEndWhilePattern    = `(?i)^(\s*?)(\t*?)(End While|Wend)`
	VBScriptSelectPattern      = `(?i)^(\s*?)(\t*?)(Select Case )`
	VBScriptEndSelectPattern   = `(?i)^(\s*?)(\t*?)(End Select)($|(\s+?)|(\t+?))`
	VBScriptFunctionPattern    = `(?i)^(\s*?)(\t*?)(Private |Public (Default )?)?(Function|Sub)`
	VBScriptEndFunctionPattern = `(?i)^(\s*?)(\t*?)(End (Function|Sub))`
)

// VBScriptFileExtension defines the VBSCript file extensions.
const VBScriptFileExtension = `(?i)(.wsf|.vbs|.asp|.inc)$`

// VBScriptJSON defines the VBScript JSON.
type VBScriptJSON struct {
	VBScript []VBScript
}

// Cognitive is to set state of cognitive complexity
type Cognitive struct {
	NestState           int `json:"-"`
	CognitiveComplexity int `json:"cognitive_complexity"`
}

func (v *Cognitive) BeginNest() {
	v.CognitiveComplexity++
	v.CognitiveComplexity += v.NestState
	v.NestState++
}

func (v *Cognitive) EndNest() {
	v.NestState--
}

func (v *Cognitive) Increment() {
	v.CognitiveComplexity++
}

// VBScript is to set VBScript Code Architecture.
type VBScript struct {
	Cognitive
	FileName        string `json:"file_name"`
	IsBeginFunction bool   `json:"-"`
	Functions       []Function
}

// Function is to set function state to VBSCript.
type Function struct {
	Cognitive
	FunctionName string `json:"function_name"`
}
