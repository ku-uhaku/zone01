package go_reloaded

import (
	"strconv"
	"strings"
	"unicode"
)

// -----------------------------------------------------------------------------
// PUBLIC ENTRY POINT
// -----------------------------------------------------------------------------

func ProcessTokens(tokens []Token) string {
	var words []string

	for i := 0; i < len(tokens); i++ {
		t := tokens[i]

		switch t.Type {

		case WORD:
			words = append(words, t.Value)

		case COMMAND:
			applyCommand(&words, t)

		case PUNCT:
			words = append(words, t.Value)
		}
	}

	return joinSmart(words)
}

// -----------------------------------------------------------------------------
// COMMAND EXECUTION
// -----------------------------------------------------------------------------

func applyCommand(words *[]string, cmd Token) {
	clean := removeCommas(cmd.Children)

	if len(clean) == 0 {
		return
	}

	// ------------------------------------------
	// CASE 1: simple (up) or (low)
	// modifies previous word
	// ------------------------------------------
	if len(clean) == 1 {
		name := clean[0].Value
		applyToLast(words, name)
		return
	}

	// ------------------------------------------
	// CASE 2: command with number: (low, 4)
	// ------------------------------------------
	name := processToken(clean[0]) // e.g. "low"
	numStr := processToken(clean[1])
	n, err := strconv.Atoi(numStr)
	if err != nil {
		return
	}

	applyToPreviousN(words, name, n)
}

// -----------------------------------------------------------------------------
// REMOVE COMMAS INSIDE COMMAND
// -----------------------------------------------------------------------------

func removeCommas(tokens []Token) []Token {
	var out []Token
	for _, t := range tokens {
		if t.Type == PUNCT && t.Value == "," {
			continue
		}
		out = append(out, t)
	}
	return out
}

// -----------------------------------------------------------------------------
// APPLY TO PREVIOUS WORD(S)
// -----------------------------------------------------------------------------

func applyToLast(words *[]string, cmd string) {
	if len(*words) == 0 {
		return
	}

	last := len(*words) - 1

	if isWord((*words)[last]) {
		(*words)[last] = applyFunction((*words)[last], cmd)
	}
}

func applyToPreviousN(words *[]string, cmd string, n int) {
	count := 0

	for i := len(*words) - 1; i >= 0; i-- {

		if isWord((*words)[i]) {

			(*words)[i] = applyFunction((*words)[i], cmd)
			count++

			if count == n {
				return
			}
		}
	}
}

// -----------------------------------------------------------------------------
// APPLY FUNCTION: up, low, cap, bin, hex
// -----------------------------------------------------------------------------

func applyFunction(str, cmd string) string {
	switch cmd {
	case "up":
		return ToUp(str)
	case "low":
		return ToLow(str)
	case "cap":
		return ToCap(str)
	case "bin":
		return ToBin(str)
	case "hex":
		return (str)
	}
	return str
}

// -----------------------------------------------------------------------------
// TOKEN EVALUATION
// -----------------------------------------------------------------------------

func processToken(t Token) string {
	switch t.Type {

	case WORD:
		return t.Value

	case COMMAND:
		// nested command recursion:
		clean := removeCommas(t.Children)

		if len(clean) == 1 {
			return clean[0].Value
		}

		// handle nested (bin) or (hex)
		if len(clean) == 1 {
			return clean[0].Value
		}

		// could expand if nested function needed
		return ""

	case PUNCT:
		return t.Value
	}

	return ""
}

// -----------------------------------------------------------------------------
// WORD / PUNCT HELPERS
// -----------------------------------------------------------------------------

func isWord(s string) bool {
	if len(s) == 0 {
		return false
	}
	r := rune(s[0])
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func isPunctuationRune(r rune) bool {
	switch r {
	case ',', '.', '!', '?', ':', ';':
		return true
	}
	return false
}

// -----------------------------------------------------------------------------
// JOIN WITH CORRECT SPACING
// -----------------------------------------------------------------------------

func joinSmart(words []string) string {
	var out strings.Builder

	for i, w := range words {
		if i > 0 {
			// no space before punctuation
			if !isPunctuationRune(rune(w[0])) {
				out.WriteRune(' ')
			}
		}
		out.WriteString(w)
	}

	return out.String()
}
