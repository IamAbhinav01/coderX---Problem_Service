package models

import (
	"time"
)

type TestCase struct {
	Input  string `bson:"input" json:"input"`
	Output string `bson:"output" json:"output"`
}

type CodeSnippet struct {
	Language     string `bson:"language" json:"language"`
	StartSnippet string `bson:"startSnippet" json:"startSnippet"`
	MidSnippet   string `bson:"midSnippet,omitempty" json:"midSnippet,omitempty"`
	EndSnippet   string `bson:"endSnippet" json:"endSnippet"`
}

type Problem struct {
	ID           string        `bson:"_id,omitempty" json:"id,omitempty"`
	Title        string        `bson:"title" json:"title"`
	Description  string        `bson:"description" json:"description"`
	Difficulty   string        `bson:"difficulty" json:"difficulty"`
	TestCases    []TestCase    `bson:"testCases" json:"testCases"`
	CodeSnippets []CodeSnippet `bson:"codeSnippets" json:"codeSnippets"`
	Editorial    string        `bson:"editorial" json:"editorial,omitempty"`
	Topic        string        `bson:"topic" json:"topic,omitempty"`
	CreatedAt    time.Time     `bson:"createdAt" json:"createdAt"`
}
