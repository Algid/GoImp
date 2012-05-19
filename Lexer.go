package lexer

import(
    "io"
    "bufio"
    "bytes"
    "strings"
)

type Lexer struct{
    source *bufio.Reader
    listing *io.Writer

    lexeme string
    currentLine string
}

const(
    T_and = iota; T_begin; T_boolean; T_break; T_call; T_end; T_else; T_elsif; T_false;
    T_function; T_halt; T_if; T_input; T_integer; T_is; T_loop; T_not; T_null;
    T_newline; T_or; T_output; T_procedure; T_return; T_then; T_true; T_var; T_while;
    T_comma; T_colon; T_lparen; T_rparen; T_semi; T_lt; T_le; T_gt; T_ge; T_eq; T_ne;
    T_plus; T_minus; T_mult; T_div; T_mod; T_assign; T_error; T_id; T_number; T_string; T_eof;
)

func New(src *io.Reader, list *io.Writer) *Lexer{
    lex := new(Lexer)
    lex.source = bufio.NewReader(*src)
    //lex.listing = bufio.NewWriter(*list)
    lex.listing = list
    return lex
}

func (lex *Lexer) GetLexeme() string{
    return lex.lexeme
}

func (lex *Lexer) GetChar() (char rune, err error) {
    var (
        part []byte
        prefix bool
    )
    //lex.source = bufio.NewReader(source)
    buffer := bytes.NewBuffer(make([]byte,0))
    if lex.currentLine == ""{
        if part, prefix, err = lex.source.ReadLine(); err != nil {
            return
        }
        buffer.Write(part)
        if !prefix {
            lex.currentLine = buffer.String()
            if _, err = io.WriteString(*lex.listing, lex.currentLine); err != nil {
                return
            }
            buffer.Reset()
        }
    }
    runeReader := strings.NewReader(lex.currentLine)
    if char, _, err = runeReader.ReadRune(); err != nil {
        return
    }
    return
}
