package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TestCase struct{
	Input string `bson:"input" json:"input"`
	Output string `bson:"output" json:"output"`
}

type CodeSnippet struct{
	Language string `bson:"language" json:"language"`
	StartSnippet string `bson:"startsnippet" json:"startsnippet"`
	MidSnippet string `bson:"midsnippet" json:"midsnippet,omitempty"`
	EndSnippet string `bson:"endsnippet" json:"endsnippet"`
}

type Problem struct {
	ID          primitive.ObjectID
	Title       string `bson:"title" json:"title"`
	Tescription string `bson:"Tescription" json:"Tescription"`
	Difficulty  string `bson:"Difficulty" json:"Difficulty"`
	TestCases []TestCase `bson:"testcases" json:"testcases"`
	CodeSnippets []CodeSnippet `bson:"codesnippets" json:"codesnippets"`
	Editorial string `bson:"editorial" json:"editorial,omitempty"`
	Topic string `bson:"topic" json:"topic,omitempty"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}