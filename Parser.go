package Parser

import (
    "fmt"
    "os"
    "GoImp"
)

type Parser struct{
    Lex lexer
    Listing *io.Writer
    Token int
}

func New(src *io.Reader, list *io.Writer) *Parser
{
    Parse := new(Parser)
    Parse.Lex := lexer.New(src,list)
    
}

func (parse *Parser) Parse()
{
    parse.Token = parse.Lex.GetToken()
    Program()
}

func (parse *Parser) Accept(t int)
{
    if parse.Token == t {
        parse.Token = parse.Lex.GetToken()
    }
    else {
        parse.ErrorMessage(t, parse.Token)
    }
}

func (parse *Parser) ErrorMessage(expected int, found int)
{
    buffer := bytes.NewBuffer(make([]byte,0))
    buffer.Write("\n\n I was expecting a " + parse.Lex.TokenToString(expected) + " but I found a " + parse.Lex.TokenToString(found) + " " + parse.Lex.Lexeme)
    io.WriteString(*parse.Listing, buffer.String())
    os.Exit(1)
}

func (parse *Parse) StringErrorMessage(message string)
{
    buffer := bytes.NewBuffer(make([]byte,0))
    buffer.Write(("\n" + message))
    io.WriteString(*parse.Listing, buffer.String())
    os.Exit(1)
}

func (parse *Parse) BeginsFactor(t int) ret bool
{
    switch(t){
    case lexer.T_minus:
    case lexer.T_not:
    case lexer.T_number:
    case lexer.T_false:
    case lexer.T_true:
    case lexer.T_id:
    case lexer.T_lparen:
        return true
    default:
        return false
    }
}

//PROGRAM ::= VARDECS SUBROUTINES
func (parse *Parser) Program()
{
    parse.Vardecs()
    parse.Subroutines()
}

//VARDECS ::= var VARDECLIST { VARDECLIST } | empty
func (parse *Parser) Vardecs()
{
    //TRACER t("vardecs " + lex.getLexeme())
    if parse.Token != lexer.T_var {return}
    parse.Accept(lexer.T_var)
    parse.Vardeclist()
    for parse.Token == lexer.T_id {
        parse.Vardeclist()
    }
}

//VARDECLIST ::= id {, id } : TYPE
func (parse *Parser) Vardeclist()
{
    //TRACER t("vardeclist " + lex.getLexeme())
    parse.Accept(t_id)
    for parse.Token == lexer.T_comma {
        parse.Accept(lexer.T_comma)
        parse.Accept(lexer.T_id)
    }
    parse.Accept(lexer.T_colon)
    parse.Type()
    parse.Accept(lexer.T_semi)
}

//TYPE ::= integer | boolean
func (parse *Parser) Type()
{
    //TRACER t("type " + lex.getLexeme())
    if parse.Token == lexer.T_integer {
        parse.Accept(lexer.T_integer)
    }

    else if parse.Token == lexer.T_boolean {
        parse.Accept(lexer.T_boolean)
    }

    else {ErrorMessage(lexer.T_boolean,parse.Token)}
}

//SUBROUTINES ::= PROCFUN { PROCFUN }
func (parse *Parser) Subroutines()
{
    //TRACER t("subroutines " + lex.getLexeme())
    parse.Procfun()
    for parser.Token == lexer.T_procedure || parse.Token == lexer.T_function {
        parse.Procfun()
    }
}

//PROCFUN ::= PROC | FUN
func (parse *Parser) Procfun()
{
    //TRACER t("procfun " + lex.getLexeme())
    if parse.Token == lexer.T_procedure {
        parse.Proc()
    }

    if parse.Token == lexer.T_function {
        parse.Fun()
    }
}

//PROC ::= procedure id ( PARAMETERS ) is PROCBODY
func (parse *Parser) Proc()
{
    //TRACER t("proc " + lex.getLexeme())
    parse.Accept(lexer.T_procedure)
    parse.Accept(lexer.T_id)
    parse.Accept(lexer.T_lparen)
    parse.Parameters()
    parse.Accept(lexer.T_rparen)
    parse.Accept(lexer.T_is)
    parse.Procbody()
}

//FUN ::= function id ( PARAMETERS ) return TYPE is PROCBODY
func (parse *Parser) Fun()
{
    //TRACER t("fun " + lex.getLexeme())
    parse.Accept(lexer.T_function)
    parse.Accept(lexer.T_id)
    parse.Accept(lexer.T_lparen)
    parse.Parameters()
    parse.Accept(lexer.T_rparen)
    parse.Accept(lexer.T_return)
    parse.Type()
    parse.Accept(lexer.T_is)
    parse.Procbody()
}

//PARAMETERS ::= PARAMLIST { PARAMLIST} | E
func (parse *Parser) Parameters()
{
    //TRACER t("parameters " + lex.getLexeme())
    if parse.Token != lexer.T_id {
        return
    }
    parse.Paramlist()
    for parse.Token == lexer.T_semi {
        parse.Accept(lexer.T_semi)
        parse.Paramlist()
    }
}

//PARAMLIST ::= id {, id} : TYPE
func (parse *Parser) Paramlist()
{
    //TRACER t("paramlist " + lex.getLexeme())
    parse.Accept(lexer.T_id)
    for parse.Token == lexer.T_comma {
        parse.Accept(lexer.T_comma)
        parse.Accept(lexer.T_id)
    }
    parse.Accept(lexer.T_colon)
    parse.Type()
}

//PROCBODY ::= VARDECS begin STMTLIST end id
func (parse *Parser) Procbody()
{
    //TRACER t("procbody " + lex.getLexeme())
    parse.Vardecs()
    parse.Accept(lexer.T_begin)
    parse.Stmtlist()
    parse.Accept(lexer.T_end)
    parse.Accept(lexer.T_id)
    parse.Accept(lexer.T_semi)
}

//STMTLIST ::= E | STMT STMTLIST
func (parse *Parser) Stmtlist()
{
    //TRACER t("stmtlist " + lex.getLexeme())
    switch (parse.Token) {
    case lexer.T_if:
    case lexer.T_while:
    case lexer.T_id:
    case lexer.T_call:
    case lexer.T_input:
    case lexer.T_output:
    case lexer.T_null:
    case lexer.T_break:
    case lexer.T_return:
    case lexer.T_halt:
    case lexer.T_newline:
        parse.Stmt()
        parse.Accept(lexer.T_semi)
        parse.Stmtlist()
        return
    default:
        return
    }
}

//STMT ::= IFSTMT | WHILESTMT | ASSIGNSTMT
//STMT ::= CALLSTMT | INPUTSTMT | OUTPUTSTMT
//STMT ::= null |  | return
//STMT ::= return EXPR | halt | newline
func (parse *Parser) Stmt()
{
    //TRACER t("stmt " + lex.getLexeme())
    switch (parse.Token) {
    case lexer.T_if:
        parse.Ifstmt()
        break
    case lexer.T_while:
        parse.Whilestmt()
        break
    case lexer.T_id:
        parse.Assignstmt()
        break
    case lexer.T_call:
        parse.Callstmt()
        break
    case lexer.T_input:
        parse.Inputstmt()
        break
    case lexer.T_output:
        parse.Outputstmt()
        break
    case lexer.T_null:
        parse.Accept(lexer.T_null)
        break
    case lexer.T_break:
        parse.Accept(lexer.T_break)
        break
    case lexer.T_return:
        parse.Accept(lexer.T_return)
        if parse.BeginsFactor(parse.Token) {
            parse.Expr()
        }
        break
    case lexer.T_halt:
        parse.Accept(lexer.T_halt)
        break
    case lexer.T_newline:
        parse.Accept(lexer.T_newline)
        break
    default:
        parse.StringErrorMessage("Invalid Statement")
        break
    }
    return
}

//IFSTMT ::= if EXPR then STMTLIST [else STMTLIST] end if
func (parse *Parser) Ifstmt()
{
    //TRACER t("ifstmt " + lex.getLexeme())
    parse.Accept(lexer.T_if)
    parse.Expr()
    parse.Accept(T_then)
    stmtlist()
    if parse.Token == lexer.T_else {
        parse.Accept(lexer.T_else)
        parse.Stmtlist()
    }

    parse.Accept(lexer.T_end)
    parse.Accept(lexer.T_if)
}

//WHILESTMT ::= while EXPR loop STMTLIST end loop
func (parse *Parser) Whilestmt()
{
    //TRACER whilestmt("type " + lex.getLexeme())
    parse.Accept(lexer.T_while)
    parse.Expr()
    parse.Accept(lexer.T_loop)
    parse.Stmtlist()
    parse.Accept(lexer.T_end)
    parse.Accept(lexer.T_loop)
}

//ASSIGNSTMT ::= id = EXPR
func (parse *Parser) Assignstmt()
{
    //TRACER t("assignstmt " + lex.getLexeme())
    parser.Accept(lexer.T_id)
    parser.Accept(lexer.T_assign)
    parser.Expr()
}

//CALLSTMT ::= call id( ARGLIST )
func (parse *Parser) Callstmt()
{
    //TRACER t("callstmt " + lex.getLexeme())
    parse.Accept(lexer.T_call)
    parse.Accept(lexer.T_id)
    parse.Accept(lexer.T_lparen)
    parse.Arglist()
    parse.Accept(lexer.T_rparen)
}

//INPUTSTMT ::= input(id)
func (parse *Parser) Inputstmt()
{
    //TRACER t("inputstmt " + lex.getLexeme())
    parse.Accept(lexer.T_input)
    parse.Accept(lexer.T_lparen)
    parse.Accept(lexer.T_id)
    parse.Accept(lexer.T_rparen)
}

//OUTPUTSTMT ::= output(EXPR) | output(string)
func outputstmt()
{
    //TRACER t("outputstmt " + lex.getLexeme())
    parse.Accept(lexer.T_output)
    parse.Accept(lexer.T_lparen)
    if parse.Token == lexer.T_string {
        parse.Accept(lexer.T_string)
    }
    else {
        parse.Expr()
    }
    parse.Accept(lexer.T_rparen)
}

//EXPR ::= SIMPLEEXPR [ RELOP SIMPLEEXPR ]
func (parse *Parser) Expr()
{
    //TRACER t("expr " + lex.getLexeme())
    parse.Simpleexpr()
    if parse.IsRelop(parse.Token) {
        parse.Token = parse.Lex.GetToken()
        parse.Simpleexpr()
    }
}

//SIMPLEEXPR ::= TERM { ADDOP TERM }
func (parse *Parser) Simpleexpr()
{
    //TRACER t("simpleexpr " + lex.getLexeme())
    parse.Term()
    for parse.IsAddop(parse.Token) {
        parse.Token = parse.Lex.GetToken()
        parse.Term()
    }
}

//TERM ::= FACTOR {MULTOP FACTOR }
func (parse *Parser) Term()
{
    //TRACER t("term " + lex.getLexeme())
    parse.Factor()
    for parse.IsMultop(parse.Token) {
        parse.Token = parse.Lex.GetToken()
        parse.Factor()
    }
}

//FACTOR ::= -FACTOR | not FACTOR
//FACTOR ::= number | false | true
//FACTOR ::= id | id(ARGLIST) | (EXPR)
func (parse *Parser) Factor()
{
    //TRACER t("factor " + lex.getLexeme())
    if parse.BeginsFactor(parse.Token) {
        if parse.Token == lexer.T_id {
            parse.Accept(lexer.T_id)
            if(parse.Token == lexer.T_lparen {
                parse.Accept(lexer.T_lparen)
                parse.Arglist()
                parse.Accept(lexer.T_rparen)
            }
        }
        else if paren.Token == lexer.T_minus || paren.Token == lexer.T_not {
            parse.Accept(parse.Token)
            parse.Factor()
        }
        else if paren.Token == lexer.T_lparen {
            parse.Accept(lexer.T_lparen)
            parse.Expr()
            parse.Accept(lexer.T_rparen)
        }
        else {
            parse.Accept(parse.Token)
        }
    }
    else {
        parse.StringErrorMessage("Invalid factor token: " + parse.Token)
    }
}

//ARGLIST ::= EXPR {, EXPR } | E
func (parse *Parser) Arglist()
{
    //TRACER t("arglist " + lex.getLexeme())
    if parse.BeginFactor(parse.Token) {
        parse.Expr()
        for parse.Token == lexer.T_comma {
            parse.Accept(lexer.T_comma)
            parse.Expr()
        }
    }
}

func (parse *Parser) IsRelop(t int) bool
{
    return t == lexer.T_eq || t == lexer.T_ne || t == lexer.T_le || t== lexer.T_lt || t == lexer.T_ge || t == lexer.T_gt
}

func (parse *Parser) IsAddop(t int) bool
{
    return t == lexer.T_plus || t == lexer.T_minus || t == lexer.T_or
}

func (parse *Parser) IsMultop(t int) bool
{
    return t == lexer.T_mult || t == lexer.T_div || t == lexer.T_mod || t== lexer.T_and
}
