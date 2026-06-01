package main

import (
	"fmt"
	"github.com/indium114/slog/internal"
)

func genericPrint(prefix, text string, a ...any) {
	fmt.Printf(prefix+text, a...)
}
