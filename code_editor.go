// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gxui

type CodeSuggestion interface {
	Name() string
	Code() string
}

type CodeSuggestionProvider interface {
	SuggestionsAt(runeIndex int) []CodeSuggestion
}

type CodeEditor interface {
	TextBox
	SyntaxLayers() CodeSyntaxLayers
	SetSyntaxLayers(CodeSyntaxLayers)
	TabWidth() int
	SetTabWidth(int)
	TabSpaces() bool
	SetTabSpaces(bool)
	SuggestionProvider() CodeSuggestionProvider
	SetSuggestionProvider(CodeSuggestionProvider)
	ShowSuggestionList()
	HideSuggestionList()
	HideLines(int, int)
	RevealLines(int, int)
}
