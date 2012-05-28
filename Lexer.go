package lexer

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

type Lexer struct {
	source  *bufio.Reader
	listing *io.Writer
    
    tokenToStringVector []string
    lexemeTokenMap map[string] int
    
	lexeme      string
	currentLine *strings.Reader
}

const (
	T_and = iota
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
    lex.lexeme = ""
    lex.tokenToStringVector = make([]string, 50)
    tokenNames := []string{"and","begin","boolean","break","call","end","else","elseif","false","function","halt","if","input","integer","is","loop","not","null","newline","or","output","procedure","return","then","true","var","while","@comma","@colon","@lparen","@rparen","@semi","@lt","@le","@gt","@ge","@eq","@ne","@plus","@minus","@mult","@div","@mod","@assign","@t_error","@t_id","@t_number","@t_string","@t_eof"}
    for t:= T_and; t <= T_eof; t++{
        lex.tokenToStringVector[t] = tokenNames[t]
        lex.lexemeTokenMap[tokenNames[t]] = t
    }
    
    return lex
}

func (lex *Lexer) GetLexeme() string {
	return lex.lexeme
}

func (lex *Lexer) GetChar() (char string, err error) {
	var (
		part        []byte
		prefix      bool
		currentRune rune
	)
	//lex.source = bufio.NewReader(source)
	buffer := bytes.NewBuffer(make([]byte, 0))
	if lex.currentLine == nil || lex.currentLine.Len() == 0 {
		if part, prefix, err = lex.source.ReadLine(); err != nil {
			return
		}
		buffer.Write(part)
		if !prefix {
			lex.currentLine = strings.NewReader(buffer.String())
			if _, err = io.WriteString(*lex.listing, buffer.String()); err != nil {
				return
			}
			buffer.Reset()
		}
	}
	if currentRune, _, err = lex.currentLine.ReadRune(); err != nil {
		return
	}
	char = string(currentRune)
	return
}
