package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"slices"
	"strings"
)

type letterStatus int

const (
	letterStatusCorrectPosition letterStatus = iota
	letterStatusWrongPosition
	letterStatusGuessed
)

func main() {
	targetWord := getWord()
	targetSlice := []rune(targetWord)
	alphabet := getAlphabet()
	guessedLetters := make(map[rune]letterStatus)
	out := os.Stdout
	maxGuesses := 5
	guessCount := 1
	var (
		guess string
	)

	for {
		if guessCount > maxGuesses {
			mustPrint(out, "You ran out of guesses. The word was %s\n", targetWord)
			break
		}
		mustPrint(out, "Guess (%d): ", guessCount)

		_, err := fmt.Scanln(&guess)
		if err != nil {
			panic(err)
		}

		guessCount++
		if guess == targetWord {
			mustPrint(out, "\u001b[42m%s\u001b[0m\n", targetWord)
			mustPrint(out, "You guessed it in %d guesses!\n", guessCount)
			break
		}

		for ind, c := range guess {
			switch {
			case slices.Contains(targetSlice, c) && slices.Index(targetSlice, c) == ind:
				mustPrint(out, "\u001b[42m%c\u001b[0m", c)
				guessedLetters[c] = letterStatusCorrectPosition
			case slices.Contains(targetSlice, c):
				mustPrint(out, "\u001b[43m%c\u001b[0m", c)
				guessedLetters[c] = letterStatusWrongPosition
			default:
				mustPrint(out, "%c", c)
				guessedLetters[c] = letterStatusGuessed
			}
		}
		mustPrint(out, "\n")

		mustPrint(out, "Guessed letters: ")
		for _, c := range alphabet {
			status, ok := guessedLetters[c]
			switch {
			case !ok:
				mustPrint(out, "%c", c)
			case status == letterStatusCorrectPosition:
				mustPrint(out, "\u001b[42m%c\u001b[0m", c)
			case status == letterStatusWrongPosition:
				mustPrint(out, "\u001b[43m%c\u001b[0m", c)
			case status == letterStatusGuessed:
				mustPrint(out, "\u001b[90m%c\u001b[0m", c)
			}
		}
		mustPrint(out, "\n")
	}
}

func mustPrint(out io.Writer, s string, args ...any) {
	_, err := fmt.Fprintf(out, s, args...)
	if err != nil {
		panic(err)
	}
}

func getWord() string {
	contents, err := os.ReadFile("data/words.txt")
	if err != nil {
		panic(err)
	}
	words := strings.Split(string(contents), "\n")

	index := rand.Int() % len(words)
	return words[index]
}

func getAlphabet() []rune {
	return []rune("abcdefghijklmnopqrstuvwxyz")
}
