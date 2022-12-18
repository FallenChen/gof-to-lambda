package main

import "strings"

type TextFormatter interface {
	filter(text string) bool
	format(text string) string
}

type PlainTextFormatter struct {
}

func (f *PlainTextFormatter) filter(text string) bool {
	return true
}

func (f *PlainTextFormatter) format(text string) string {
	return text
}

type ErrorTextFormatter struct {
}

func (f *ErrorTextFormatter) filter(text string) bool {
	return strings.HasPrefix(text, "ERROR")
}

func (f *ErrorTextFormatter) format(text string) string {
	return strings.ToUpper(text)
}

type ShortTextFormatter struct {
}

func (f *ShortTextFormatter) filter(text string) bool {
	return len(text) < 20
}

func (f *ShortTextFormatter) format(text string) string {
	return strings.ToLower(text)
}

type TextEditor struct {
	TextFormatter
}

func (f *TextEditor) publish(text string) {
	if f.filter(text) {
		println(f.format(text))
	}
}

//func main() {
//	errorTextFormatter := ErrorTextFormatter{}
//	textEditor := &TextEditor{&errorTextFormatter}
//
//	textEditor.publish("ERROR - something bad happened")
//	textEditor.publish("DEBUG - I'm here")
//}
