package main

import (
	"math/rand"
	"strings"
)

type Bobbot struct {
	answers []string
}

var answers []string = []string{
	"Controllo",
	"A che ti serve?",
	"Te lo dico dopo",
	"Non lo devi chiedere a me",
	"Hai fatto una RDA?",
	"Chiedi a [insert random name]",
	"[Silenzio per 2 secondi, poi se ne va]",
	"Bobbot.exe is rebooting, please wait...",
}

func NewBobbot() *Bobbot {
	return &Bobbot{cloneSlice(answers)}
}

func (b *Bobbot) WelcomeMessage() string {
	return strings.Join([]string{
		"Bobbot v0.2.1",
		"===",
		"Il simulatore di risposte di Bobbers.",
		"",
	}, "\n")
}

func (b *Bobbot) Ask(question string) string {
	answersMaxIndex := len(b.answers) - 1
	answerIndex := rand.Intn(answersMaxIndex)
	return b.answers[answerIndex]
}

func cloneSlice[T any](s []T) []T {
	theCopy := make([]T, len(s))
	copy(theCopy, s)
	return theCopy
}
