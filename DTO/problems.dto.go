package dto

type TestCase struct {
	Input  string `bson:"input" json:"input" validate:"required"`
	Output string `bson:"output" json:"output" validate:"required"`
}

type CodeSnippet struct {
	Language     string `bson:"language" json:"language" validate:"required"`
	StartSnippet string `bson:"startSnippet" json:"startSnippet" validate:"required"`
	MidSnippet   string `bson:"midSnippet,omitempty" json:"midSnippet,omitempty" validate:"required"`
	EndSnippet   string `bson:"endSnippet" json:"endSnippet" validate:"required"`
}

type GenerateProblem struct {
	ID           string        `bson:"_id,omitempty" json:"id,omitempty"`
	Title        string        `bson:"title" json:"title"`
	Description  string        `bson:"description" json:"description"`
	Difficulty   string        `bson:"difficulty" json:"difficulty"`
	TestCases    []TestCase    `bson:"testCases" json:"testCases"`
	CodeSnippets []CodeSnippet `bson:"codeSnippets" json:"codeSnippets"`
	Editorial    string        `bson:"editorial" json:"editorial,omitempty"`
	Topic        string        `bson:"topic" json:"topic,omitempty"`
}
