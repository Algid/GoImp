package lexer

import(
    "io"
)

type Lexer struct{
    source io.Reader
    listing io.Writer

    lexeme string
}

const(
    t_and = iota; t_begin; t_boolean; t_break; t_call; t_end; t_else; t_elsif; t_false;
    t_function; t_halt; t_if; t_input; t_integer; t_is; t_loop; t_not; t_null;
    t_newline; t_or; t_output; t_procedure; t_return; t_then; t_true; t_var; t_while;
    t_comma; t_colon; t_lparen; t_rparen; t_semi; t_lt; t_le; t_gt; t_ge; t_eq; t_ne;
    t_plus; t_minus; t_mult; t_div; t_mod; t_assign; t_error; t_id; t_number; t_string; t_eof;
)

func New(src io.Reader, list io.Writer) *Lexer{
    return new(Lexer)
}

func (lex Lexer) GetLexeme() string{
    return lex.lexeme
}

func GetChar() string {
    return "l"
}
