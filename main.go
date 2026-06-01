package main

import (
	"fmt"
	"github.com/indium114/slog/internal"
)

// MARK: generic print function so that i don't have to retype so much code :P
func genericPrint(prefix, text string, a ...any) {
	fmt.Printf(prefix+" "+text, a...)
}

// MARK: actual functions, let's go :D
func Ok(text string, a ...any) {
	genericPrint(internal.OkText, text, a...)
}

func Warn(text string, a ...any) {
	genericPrint(internal.WarnText, text, a...)
}

func Err(text string, a ...any) error {
	return fmt.Errorf(internal.ErrText+" "+text, a...)
}

func Sync(text string, a ...any) {
	genericPrint(internal.SyncText, text, a...)
}

func Update(text string, a ...any) {
	genericPrint(internal.UpdateText, text, a...)
}

func Log(text string, a ...any) {
	genericPrint(internal.LogText, text, a...)
}

func Clean(text string, a ...any) {
	genericPrint(internal.CleanText, text, a...)
}

func Add(text string, a ...any) {
	genericPrint(internal.AddText, text, a...)
}

func Hint(text string, a ...any) {
	genericPrint(internal.HintText, text, a...)
}

func Query(text string, a ...any) {
	genericPrint(internal.QueryText, text, a...)
}
