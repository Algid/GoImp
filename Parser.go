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
func procbody()
{
    //TRACER t("procbody " + lex.getLexeme())
    vardecs()
    accept(t_begin)
    stmtlist()
    accept(t_end)
    accept(t_id)
    accept(t_semi)
}

//STMTLIST ::= E | STMT STMTLIST
func stmtlist()
{
    //TRACER t("stmtlist " + lex.getLexeme())
    if token == t_end || token == t_else {return}
    stmt()
    accept(t_semi)
    stmtlist()
}

//STMT ::= IFSTMT | WHILESTMT | ASSIGNSTMT
//STMT ::= CALLSTMT | INPUTSTMT | OUTPUTSTMT
//STMT ::= null |  | return
//STMT ::= return EXPR | halt | newline
func stmt()
{
    //TRACER t("stmt " + lex.getLexeme())
    switch (token) 
    {
    case t_if: ifstmt()
        
    case t_while: whilestmt()
        
    case t_id: assignstmt()
        
    case t_call: callstmt()
        
    case t_input: inputstmt()
        
    case t_output: outputstmt()
        
    case t_null: accept(t_null)
        
    case t_break: accept(t_break)
        
    case t_return: accept(t_return)
        if token != t_semi {
            expr()
        }
        
    case t_halt: accept(t_halt)
        
    case t_newline: accept(t_newline)
    }
}

//IFSTMT ::= if EXPR then STMTLIST [else STMTLIST] end if
func ifstmt()
{
    //TRACER t("ifstmt " + lex.getLexeme())
    accept(t_if)
    expr()
    accept(t_then)
    stmtlist()
    if token == t_else {
        accept(t_else)
        stmtlist()
    }

    accept(t_end)
    accept(t_if)
}

//WHILESTMT ::= while EXPR loop STMTLIST end loop
func whilestmt()
{
    //TRACER whilestmt("type " + lex.getLexeme())
    accept(t_while)
    expr()
    accept(t_loop)
    stmtlist()
    accept(t_end)
    accept(t_loop)
}

//ASSIGNSTMT ::= id = EXPR
func assignstmt()
{
    //TRACER t("assignstmt " + lex.getLexeme())
    accept(t_id)
    accept(t_assign)
    expr()
}

//CALLSTMT ::= call id( ARGLIST )
func callstmt()
{
    //TRACER t("callstmt " + lex.getLexeme())
    accept(t_call)
    accept(t_id)
    accept(t_lparen)
    arglist()
    accept(t_rparen)
}

//INPUTSTMT ::= input(id)
func inputstmt()
{
    //TRACER t("inputstmt " + lex.getLexeme())
    accept(t_input)
    accept(t_lparen)
    accept(t_id)
    accept(t_rparen)
}

//OUTPUTSTMT ::= output(EXPR) | output(string)
func outputstmt()
{
    //TRACER t("outputstmt " + lex.getLexeme())
    accept(t_output)
    accept(t_lparen)
    if token == t_string {
        accept(t_string)
        accept(t_rparen)
    }

    else {
        expr()
        accept(t_rparen)
    }
}

//EXPR ::= SIMPLEEXPR [ RELOP SIMPLEEXPR ]
func expr()
{
    //TRACER t("expr " + lex.getLexeme())
    simpleexpr()
    for ;isRelop(token); {
        token = getToken()
        simpleexpr()
    }
}

//SIMPLEEXPR ::= TERM { ADDOP TERM }
func simpleexpr()
{
    //TRACER t("simpleexpr " + lex.getLexeme())
    term()
    for ;isAddop(token); {
        token = getToken()
        term()
    }
}

//TERM ::= FACTOR {MULTOP FACTOR }
func term()
{
    //TRACER t("term " + lex.getLexeme())
    factor()
    for ;isMultop(token); {
        token = getToken()
        factor()
    }
}

//FACTOR ::= -FACTOR | not FACTOR
//FACTOR ::= number | false | true
//FACTOR ::= id | id(ARGLIST) | (EXPR)
func factor()
{
    //TRACER t("factor " + lex.getLexeme())
    switch (token)
    {
    case t_minus: accept(t_minus)
        factor()
        
    case t_not: accept(t_not)
        factor()
        
    case t_number: accept(t_number)
        
    case t_false: accept(t_false)
        
    case t_true: accept(t_true)
        
    case t_id: accept(t_id)
        if (token == t_lparen) {
            accept(t_lparen)
            arglist()
            accept(t_rparen)
        }
        
    case t_lparen: accept(t_lparen)
        expr()
        accept(t_rparen)
    }
}

//ARGLIST ::= EXPR {, EXPR } | E
func arglist()
{
    //TRACER t("arglist " + lex.getLexeme())
    if token == t_rparen {return}
    expr()
    for ; token == t_comma; {
        accept(t_comma)
        expr()
    }
}

bool isRelop(TokenType t)
{
    return t == t_eq || t == t_ne || t == t_le || t== t_lt ||
    t == t_ge || t == t_gt
}

bool isAddop(TokenType t)
{
    return t == t_plus || t == t_minus || t == t_or
}

bool isMultop(TokenType t)
{
    return t == t_mult || t == t_div || t == t_mod || t== t_and
}
