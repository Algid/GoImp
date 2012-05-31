package lexer

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
	"unicode"
)

type Lexer struct {
	source  *bufio.Reader
	listing *io.Writer

	TokenToStringVector []string
	LexemeTokenMap      map[string]int

	Lexeme      string
	ch          rune
	currentLine *strings.Reader
}

const (
	_ = iota
	T_and
	T_begin
	T_boolean
	T_break
	T_call
	T_end
	T_else
	T_elsif
	T_false
	T_function
	T_halt
	T_if
	T_input
	T_integer
	T_is
	T_loop
	T_not
	T_null
	T_newline
	T_or
	T_output
	T_procedure
	T_return
	T_then
	T_true
	T_var
	T_while
	T_comma
	T_colon
	T_lparen
	T_rparen
	T_semi
	T_lt
	T_le
	T_gt
	T_ge
	T_eq
	T_ne
	T_plus
	T_minus
	T_mult
	T_div
	T_mod
	T_assign
	T_error
	T_id
	T_number
	T_string
	T_eof
)

func New(src *io.Reader, list *io.Writer) *Lexer {
	lex := new(Lexer)
	lex.source = bufio.NewReader(*src)
	lex.listing = list
	lex.ch, _ = lex.GetChar()
	lex.Lexeme = ""
	lex.TokenToStringVector = make([]string, 50)
	lex.LexemeTokenMap = map[string]int{}
	tokenNames := []string{"and", "begin", "boolean", "break", "call", "end", "else", "elseif", "false", "function", "halt", "if", "input", "integer", "is", "loop", "not", "null", "newline", "or", "output", "procedure", "return", "then", "true", "var", "while", "@comma", "@colon", "@lparen", "@rparen", "@semi", "@lt", "@le", "@gt", "@ge", "@eq", "@ne", "@plus", "@minus", "@mult", "@div", "@mod", "@assign", "@t_error", "@t_id", "@t_number", "@t_string", "@t_eof"}
	for t := T_and; t <= T_eof; t++ {
		lex.TokenToStringVector[t] = tokenNames[t-1]
		lex.LexemeTokenMap[tokenNames[t-1]] = t
	}

	return lex
}

func (lex *Lexer) GetChar() (char rune, err error) {
	var (
		part        []byte
		prefix      bool
		currentRune rune
	)
	buffer := bytes.NewBuffer(make([]byte, 0))
	if lex.currentLine == nil || lex.currentLine.Len() == 0 {
		if part, prefix, err = lex.source.ReadLine(); err != nil {
			return
		}
		buffer.Write(part)
		fmt.Fprint(buffer, "\n")
		if !prefix {
			lex.currentLine = strings.NewReader(buffer.String())
			fmt.Println(buffer.String())
			if _, err = io.WriteString(*lex.listing, buffer.String()); err != nil {
				return
			}
			buffer.Reset()
		}
	}
	if currentRune, _, err = lex.currentLine.ReadRune(); err != nil {
		fmt.Println(err)
		return
	}
	char = currentRune
	return
}

func (lex *Lexer) GetToken() (token int, err error) {

	lex.Lexeme = ""
	for unicode.IsSpace(lex.ch) {
		lex.ch, err = lex.GetChar()
	}
	if err != nil {
		token = T_eof
		return
	}
	if unicode.IsDigit(lex.ch) {
		for unicode.IsDigit(lex.ch) {
			lex.Lexeme += string(lex.ch)
			lex.ch, err = lex.GetChar()
		}
		token = T_number
		return
	}
	if unicode.IsLetter(lex.ch) || (string(lex.ch)) == "_" {
		for unicode.IsLetter(lex.ch) || (string(lex.ch)) == "_" {
			lex.Lexeme += string(lex.ch)
			lex.ch, err = lex.GetChar()
		}
		if lex.LexemeTokenMap[lex.Lexeme] == 0 {
			token = T_id
			return
		} else {
			token = lex.LexemeTokenMap[lex.Lexeme]
			return
		}
	}

	switch string(lex.ch) {

	case "#":

		lex.Lexeme += string(lex.ch)

		lex.ch, err = lex.GetChar()
		token = T_ne
		return
		break
	case ",":

		lex.Lexeme += string(lex.ch)

		lex.ch, err = lex.GetChar()
		token = T_comma
		return
		break
	case ":":

		lex.Lexeme += string(lex.ch)

		lex.ch, err = lex.GetChar()
		token = T_colon
		return
		break
	case "(":

		lex.Lexeme += string(lex.ch)

		lex.ch, err = lex.GetChar()
		token = T_lparen
		return
		break
	case ")":

		lex.Lexeme += string(lex.ch)

		lex.ch, err = lex.GetChar()
		token = T_rparen
		return
		break
	case ";":

		lex.Lexeme += string(lex.ch)

		lex.ch, err = lex.GetChar()
		token = T_semi
		return
		break
	case "<":

		lex.Lexeme += string(lex.ch)

		lex.ch, err = lex.GetChar()
		if string(lex.ch) == "=" {
			lex.Lexeme += string(lex.ch)
			lex.ch, err = lex.GetChar()
			token = T_le
		} else {
			token = T_lt
		}
		return
		break
	case ">":

		lex.Lexeme += string(lex.ch)

		lex.ch, err = lex.GetChar()
		if string(lex.ch) == "=" {
			lex.Lexeme += string(lex.ch)
			lex.ch, err = lex.GetChar()
			token = T_ge
		} else {
			token = T_gt
		}
		return
		break
	case "=":

		lex.Lexeme += string(lex.ch)

		lex.ch, err = lex.GetChar()
		if string(lex.ch) == "=" {
			lex.Lexeme += string(lex.ch)
			lex.ch, err = lex.GetChar()
			token = T_eq
		} else {
			token = T_assign
		}
		return
		break
	case "/":

		lex.Lexeme += string(lex.ch)

		lex.ch, err = lex.GetChar()
		if string(lex.ch) == "/" {
			for string(lex.ch) != "\n" {
				lex.ch, err = lex.GetChar()
			}
			token, err = lex.GetToken()
			return

		} else {
			token = T_div
		}
		return
		break
	case "!":

		lex.Lexeme += string(lex.ch)

		lex.ch, err = lex.GetChar()
		token = T_error
		return
		break
	case "+":

		lex.Lexeme += string(lex.ch)

		lex.ch, err = lex.GetChar()
		token = T_plus
		return
	case "-":

		lex.Lexeme += string(lex.ch)

		lex.ch, err = lex.GetChar()
		token = T_minus
		return
		break
	case "*":

		lex.Lexeme += string(lex.ch)

		lex.ch, err = lex.GetChar()
		token = T_mult
		return
		break
	case "%":

		lex.Lexeme += string(lex.ch)

		lex.ch, err = lex.GetChar()
		token = T_mod
		return
	case "\"":
		lex.ch, err = lex.GetChar()
		for string(lex.ch) != "\"" {
			lex.Lexeme += string(lex.ch)
			lex.ch, err = lex.GetChar()
		}
		lex.ch, err = lex.GetChar()
		token = T_string
		return
		break
	case "":
		token = T_eof
		return
		break
	default:
		lex.Lexeme += string(lex.ch)
		lex.ch, err = lex.GetChar()
		token = T_error
		return
		break
	}
	return
}
