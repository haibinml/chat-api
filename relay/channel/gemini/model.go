package gemini

type ChatRequest struct {
	Contents         []ChatContent        `json:"contents"`
	SafetySettings   []ChatSafetySettings `json:"safety_settings,omitempty"`
	GenerationConfig ChatGenerationConfig `json:"generation_config,omitempty"`
	Tools            []ChatTools          `json:"tools,omitempty"`
}

type InlineData struct {
	MimeType string `json:"mimeType"`
	Data     string `json:"data"`
}

type FunctionCall struct {
	FunctionName string `json:"name"`
	Arguments    any    `json:"args"`
}
type EmbeddingRequest struct {
	Model                string      `json:"model"`
	Content              ChatContent `json:"content"`
	TaskType             string      `json:"taskType,omitempty"`
	Title                string      `json:"title,omitempty"`
	OutputDimensionality int         `json:"outputDimensionality,omitempty"`
}
type BatchEmbeddingRequest struct {
	Requests []EmbeddingRequest `json:"requests"`
}
type EmbeddingResponse struct {
	Embeddings []EmbeddingData `json:"embeddings"`
	Error      *Error          `json:"error,omitempty"`
}
type EmbeddingData struct {
	Values []float64 `json:"values"`
}
type Error struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Status  string `json:"status,omitempty"`
}
type Part struct {
	Text         string        `json:"text,omitempty"`
	InlineData   *InlineData   `json:"inlineData,omitempty"`
	FunctionCall *FunctionCall `json:"functionCall,omitempty"`
}

type ChatContent struct {
	Role  string `json:"role,omitempty"`
	Parts []Part `json:"parts"`
}

type ChatSafetySettings struct {
	Category  string `json:"category"`
	Threshold string `json:"threshold"`
}

type ChatTools struct {
	FunctionDeclarations any `json:"function_declarations,omitempty"`
}

type ChatGenerationConfig struct {
	Temperature     *float64 `json:"temperature,omitempty"`
	TopP            float64  `json:"topP,omitempty"`
	TopK            float64  `json:"topK,omitempty"`
	MaxOutputTokens uint     `json:"maxOutputTokens,omitempty"`
	CandidateCount  int      `json:"candidateCount,omitempty"`
	StopSequences   []string `json:"stopSequences,omitempty"`
}

type ChatResponse struct {
	Candidates     []ChatCandidate    `json:"candidates"`
	ModelVersion   string             `json:"modelVersion"`
	UsageMetadata  UsageMetadata      `json:"usageMetadata"`
	PromptFeedback ChatPromptFeedback `json:"promptFeedback"`
}

type UsageMetadata struct {
	TotalTokens          int `json:"totalTokens"`
	PromptTokenCount     int `json:"promptTokenCount,omitempty"`
	TotalTokenCount      int `json:"totalTokenCount,omitempty"`
	CandidatesTokenCount int `json:"candidatesTokenCount,omitempty"`
	ThoughtsTokenCount   int `json:"thoughtsTokenCount,omitempty"`
}
type ChatPromptFeedback struct {
	SafetyRatings []ChatSafetyRating `json:"safetyRatings"`
}
type TokenDetail struct {
	Modality   string `json:"modality"`
	TokenCount int    `json:"tokenCount"`
}

type ChatCandidate struct {
	Content       ChatContent        `json:"content"`
	FinishReason  string             `json:"finishReason"`
	Index         int64              `json:"index"`
	SafetyRatings []ChatSafetyRating `json:"safetyRatings"`
}

type ChatSafetyRating struct {
	Category    string `json:"category"`
	Probability string `json:"probability"`
}
