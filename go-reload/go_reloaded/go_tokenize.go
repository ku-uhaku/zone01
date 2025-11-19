package go_reloaded

import (
	"fmt"
	"unicode"
)

type TokenType int

const (
	WORD TokenType = iota
	COMMAND
	PUNCT
	QUOTE
)

type Token struct {
	Type     TokenType
	Value    string
	Children []Token
}


func Tokenize(text string) []Token {
	tokens, _ := tokenizeRunes([]rune(text), 0)
	PrintAllTokens(tokens)
	return tokens
}

// -----------------------------------------------------------------------------
// INTERNAL TOKENIZER
// -----------------------------------------------------------------------------

func tokenizeRunes(runes []rune, start int) ([]Token, int) {
	var tokens []Token
	i := start

	for i < len(runes) {
		c := runes[i]

		// COMMAND ------------------------------------------------------------
		if c == '(' {
			cmdToken, nextPos := readCommand(runes, i)
			tokens = append(tokens, cmdToken)
			i = nextPos
			continue
		}

		// WORD ---------------------------------------------------------------
		if isWordRune(c) {
			start := i
			for i < len(runes) && isWordRune(runes[i]) {
				i++
			}
			tokens = append(tokens, Token{
				Type:  WORD,
				Value: string(runes[start:i]),
			})
			continue
		}

		// QUOTE --------------------------------------------------------------
		if c == '\'' {
			tokens = append(tokens, Token{Type: QUOTE, Value: "'"})
			i++
			continue
		}

		// PUNCT --------------------------------------------------------------
		if isPunctuation(c) {
			tokens = append(tokens, Token{Type: PUNCT, Value: string(c)})
			i++
			continue
		}

		i++
	}

	return tokens, i
}

// -----------------------------------------------------------------------------
// READ COMMAND + CHILDREN
// -----------------------------------------------------------------------------

func readCommand(runes []rune, start int) (Token, int) {
	i := start + 1
	depth := 1

	for i < len(runes) && depth > 0 {
		if runes[i] == '(' {
			depth++
		} else if runes[i] == ')' {
			depth--
		}
		i++
	}

	value := string(runes[start:i])
	inner := value[1 : len(value)-1]

	children, _ := tokenizeRunes([]rune(inner), 0)

	return Token{
		Type:     COMMAND,
		Value:    value,
		Children: children,
	}, i
}

// -----------------------------------------------------------------------------
// HELPERS
// -----------------------------------------------------------------------------

// allowed WORD runes except: ' ( ) . ! ? , : ;
func isWordRune(r rune) bool {
	switch r {
	case '\'', '(', ')', '.', '!', '?', ',', ':', ';':
		return false
	}

	// allow letters and digits
	if unicode.IsLetter(r) || unicode.IsDigit(r) {
		return true
	}

	// emojis (Symbol-Other)
	if unicode.Is(unicode.So, r) {
		return true
	}

	// other symbols (math, currency, etc.)
	if unicode.IsSymbol(r) {
		return true
	}

	return false
}


func isPunctuation(r rune) bool {
	switch r {
	case '.', '!', '?', ',', ':', ';':
		return true
	}
	return false
}

// -----------------------------------------------------------------------------
// DEBUG PRINTING
// -----------------------------------------------------------------------------

func PrintTokens(tokens []Token, indent int) {
	pad := func(n int) string { return string(make([]rune, n)) }

	for i, t := range tokens {
		p := pad(indent)

		typeName := map[TokenType]string{
			WORD:    "WORD",
			COMMAND: "COMMAND",
			PUNCT:   "PUNCT",
			QUOTE:   "QUOTE",
		}[t.Type]

		fmt.Printf("%sToken %d: Type=%s, Value='%s'\n", p, i, typeName, t.Value)

		if len(t.Children) > 0 {
			PrintTokens(t.Children, indent+4)
		}
	}
}

func PrintAllTokens(tokens []Token) {
	printTokensRecursive(tokens, 0)
}

func printTokensRecursive(tokens []Token, level int) {
	indent := ""
	for i := 0; i < level; i++ {
		indent += "  "
	}

	for i, t := range tokens {
		typeName := map[TokenType]string{
			WORD:    "WORD",
			COMMAND: "COMMAND",
			PUNCT:   "PUNCT",
			QUOTE:   "QUOTE",
		}[t.Type]

		fmt.Printf("%sToken %d: Type=%s, Value='%s'\n",
			indent, i, typeName, t.Value)

		if len(t.Children) > 0 {
			printTokensRecursive(t.Children, level+1)
		}
	}
}
