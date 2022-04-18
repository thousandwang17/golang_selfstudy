package main

import (
	"strings"

	"github.com/pkg/profile"
)

func main() {
	defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
	simplyAdd()
	builderAdd()
}

func simplyAdd() {
	text := "text"
	textArea := ""
	for i := 0; i < 100; i++ {
		textArea += text
	}
}

func builderAdd() {
	text := "text"
	var textArea strings.Builder
	for i := 0; i < 100; i++ {
		textArea.WriteString(text)
	}
}
