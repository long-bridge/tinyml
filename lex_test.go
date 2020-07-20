package tinyml

import (
	"bytes"
	"io"
	"testing"

	"github.com/tdewolff/test"
)

type TTs []TokenType

func TestTokens(t *testing.T) {
	var tokenTests = []struct {
		css     string
		ttypes  []TokenType
		lexemes []string
	}{
		{"Hello world", TTs{TextToken}, []string{"Hello world"}},
		{"Hello world\n\nThis new next line", TTs{TextToken, NewLineToken, TextToken}, []string{"Hello world", "\n\n", "This new next line"}},
		{"\nHello world\n\nThis is new line", TTs{NewLineToken, TextToken, NewLineToken, TextToken}, []string{"\n", "Hello world", "\n\n", "This is new line"}},
		// {"==== Hello World ====", TTs{Heading2StartToken, TextToken}, []string{"==== ", "Hello World", " ===="}},
	}

	for _, tt := range tokenTests {
		t.Run(tt.css, func(t *testing.T) {
			l := NewLexer(bytes.NewBufferString(tt.css))
			i := 0
			tokens := []TokenType{}
			lexemes := []string{}
			for {
				token, data := l.Next()
				// fmt.Println("token:", token, string(data))
				if token == ErrorToken {
					test.T(t, l.Err(), io.EOF)
					break
				}

				tokens = append(tokens, token)
				// fmt.Printf("-- %q\n", string(data))
				lexemes = append(lexemes, string(data))
				i++
			}

			test.T(t, tokens, tt.ttypes, "token types must match")
			test.T(t, lexemes, tt.lexemes, "token data must match")
		})
	}
}