package parser

import (
    "os"
    "fmt"
    "strings"
    "strconv"
    "io"
    "bytes"
    "GoImp/lexer"
)

type Parser struct{
    Lex *lexer.Lexer
    Listing *io.Writer
    Token int
}

func New(src *io.Reader, list *io.Writer) *Parser{
    parse := new(Parser)
    parse.Lex = lexer.New(src,list)
    return parse
}

func (parse *Parser) Parse(){
    parse.Token,_ = parse.Lex.GetToken()
    parse.Program()
}

func (parse *Parser) Accept(t int){
    if parse.Token == t {
        parse.Token,_ = parse.Lex.GetToken()
    } else {
        parse.ErrorMessage(t, parse.Token)
    }
}

func (parse *Parser) ErrorMessage(expected int, found int){
    buff := make([]byte,0)
    buffer := bytes.NewBuffer(make([]byte,0))
    str := strings.NewReader("\n\n I was expecting a " + parse.Lex.TokenToStringVector[expected] + " but I found a " + parse.Lex.TokenToStringVector[found] + " " + parse.Lex.Lexeme)
    _,_ = str.Read(buff)
    io.WriteString(*parse.Listing, buffer.String())
    os.Exit(1)
}

func (parse *Parser) StringErrorMessage(message string){
    buff := make([]byte,0)
    buffer := bytes.NewBuffer(make([]byte,0))
    str := strings.NewReader("\n" + message)
    _,_ = str.Read(buff)
    io.WriteString(*parse.Listing, buffer.String())
    os.Exit(1)
}

func (parse *Parser) BeginsFactor(t int) (ret bool){
    switch(t){
    case lexer.T_minus, lexer.T_not, lexer.T_number, lexer.T_false, lexer.T_true, lexer.T_id, lexer.T_lparen:
        ret = true
    default:
        ret = false
    }
    return
}

//PROGRAM ::= VARDECS SUBROUTINES
func (parse *Parser) Program(){
    parse.Vardecs()
    parse.Subroutines()
}

//VARDECS ::= var VARDECLIST { VARDECLIST } | empty
func (parse *Parser) Vardecs(){
    parse.Tt("vardecs " + parse.Lex.Lexeme)
    if parse.Token != lexer.T_var {return}
    parse.Accept(lexer.T_var)
    parse.Vardeclist()
    for parse.Token == lexer.T_id {
        parse.Vardeclist()
    }
}

//VARDECLIST ::= id {, id } : TYPE
func (parse *Parser) Vardeclist(){
    parse.Tt("vardeclist " + parse.Lex.Lexeme)
    parse.Accept(lexer.T_id)
    for parse.Token == lexer.T_comma {
        parse.Accept(lexer.T_comma)
        parse.Accept(lexer.T_id)
    }
    parse.Accept(lexer.T_colon)
    parse.Type()
    parse.Accept(lexer.T_semi)
}

//TYPE ::= integer | boolean
func (parse *Parser) Type(){
    parse.Tt("type " + parse.Lex.Lexeme)
    if parse.Token == lexer.T_integer {
        parse.Accept(lexer.T_integer)
    } else if parse.Token == lexer.T_boolean {
        parse.Accept(lexer.T_boolean)
    } else {
        parse.ErrorMessage(lexer.T_boolean,parse.Token)
    }
}

//SUBROUTINES ::= PROCFUN { PROCFUN }
func (parse *Parser) Subroutines(){
    parse.Tt("subroutines " + parse.Lex.Lexeme)
    parse.Procfun()
    for parse.Token == lexer.T_procedure || parse.Token == lexer.T_function {
        parse.Procfun()
    }
}

//PROCFUN ::= PROC | FUN
func (parse *Parser) Procfun(){
    parse.Tt("procfun " + parse.Lex.Lexeme)
    if parse.Token == lexer.T_procedure {
        parse.Proc()
    }

    if parse.Token == lexer.T_function {
        parse.Fun()
    }
}

//PROC ::= procedure id ( PARAMETERS ) is PROCBODY
func (parse *Parser) Proc(){
    parse.Tt("proc " + parse.Lex.Lexeme)
    parse.Accept(lexer.T_procedure)
    parse.Accept(lexer.T_id)
    parse.Accept(lexer.T_lparen)
    parse.Parameters()
    parse.Accept(lexer.T_rparen)
    parse.Accept(lexer.T_is)
    parse.Procbody()
}

//FUN ::= function id ( PARAMETERS ) return TYPE is PROCBODY
func (parse *Parser) Fun(){
    parse.Tt("fun " + parse.Lex.Lexeme)
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
func (parse *Parser) Parameters(){
    parse.Tt("parameters " + parse.Lex.Lexeme)
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
func (parse *Parser) Paramlist(){
    parse.Tt("paramlist " + parse.Lex.Lexeme)
    parse.Accept(lexer.T_id)
    for parse.Token == lexer.T_comma {
        parse.Accept(lexer.T_comma)
        parse.Accept(lexer.T_id)
    }
    parse.Accept(lexer.T_colon)
    parse.Type()
}

//PROCBODY ::= VARDECS begin STMTLIST end id
func (parse *Parser) Procbody(){
    parse.Tt("procbody " + parse.Lex.Lexeme)
    parse.Vardecs()
    parse.Accept(lexer.T_begin)
    parse.Stmtlist()
    parse.Accept(lexer.T_end)
    parse.Accept(lexer.T_id)
    parse.Accept(lexer.T_semi)
}

//STMTLIST ::= E | STMT STMTLIST
func (parse *Parser) Stmtlist(){
    parse.Tt("stmtlist " + parse.Lex.Lexeme)
    switch (parse.Token) {
    case lexer.T_if, lexer.T_while, lexer.T_id, lexer.T_call, lexer.T_input, lexer.T_output, lexer.T_null, lexer.T_break, lexer.T_return, lexer.T_halt, lexer.T_newline:
        parse.Stmt()
        parse.Accept(lexer.T_semi)
        parse.Stmtlist()
        break
    default:
        break
    }
}

//STMT ::= IFSTMT | WHILESTMT | ASSIGNSTMT
//STMT ::= CALLSTMT | INPUTSTMT | OUTPUTSTMT
//STMT ::= null |  | return
//STMT ::= return EXPR | halt | newline
func (parse *Parser) Stmt(){
    parse.Tt("stmt " + parse.Lex.Lexeme)
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
func (parse *Parser) Ifstmt(){
    parse.Tt("ifstmt " + parse.Lex.Lexeme)
    parse.Accept(lexer.T_if)
    parse.Expr()
    parse.Accept(lexer.T_then)
    parse.Stmtlist()
    if parse.Token == lexer.T_else {
        parse.Accept(lexer.T_else)
        parse.Stmtlist()
    }

    parse.Accept(lexer.T_end)
    parse.Accept(lexer.T_if)
}

//WHILESTMT ::= while EXPR loop STMTLIST end loop
func (parse *Parser) Whilestmt(){
    parse.Tt("whilestmt " + parse.Lex.Lexeme)
    parse.Accept(lexer.T_while)
    parse.Expr()
    parse.Accept(lexer.T_loop)
    parse.Stmtlist()
    parse.Accept(lexer.T_end)
    parse.Accept(lexer.T_loop)
}

//ASSIGNSTMT ::= id = EXPR
func (parse *Parser) Assignstmt(){
    parse.Tt("assignstmt " + parse.Lex.Lexeme)
    parse.Accept(lexer.T_id)
    parse.Accept(lexer.T_assign)
    parse.Expr()
}

//CALLSTMT ::= call id( ARGLIST )
func (parse *Parser) Callstmt(){
    parse.Tt("callstmt " + parse.Lex.Lexeme)
    parse.Accept(lexer.T_call)
    parse.Accept(lexer.T_id)
    parse.Accept(lexer.T_lparen)
    parse.Arglist()
    parse.Accept(lexer.T_rparen)
}

//INPUTSTMT ::= input(id)
func (parse *Parser) Inputstmt(){
    parse.Tt("inputstmt " + parse.Lex.Lexeme)
    parse.Accept(lexer.T_input)
    parse.Accept(lexer.T_lparen)
    parse.Accept(lexer.T_id)
    parse.Accept(lexer.T_rparen)
}

//OUTPUTSTMT ::= output(EXPR) | output(string)
func (parse *Parser) Outputstmt(){
    parse.Tt("outputstmt " + parse.Lex.Lexeme)
    parse.Accept(lexer.T_output)
    parse.Accept(lexer.T_lparen)
    if parse.Token == lexer.T_string {
        parse.Accept(lexer.T_string)
    } else {
        parse.Expr()
    }
    parse.Accept(lexer.T_rparen)
}

//EXPR ::= SIMPLEEXPR [ RELOP SIMPLEEXPR ]
func (parse *Parser) Expr(){
    parse.Tt("expr " + parse.Lex.Lexeme)
    parse.Simpleexpr()
    if parse.IsRelop(parse.Token) {
        parse.Token,_ = parse.Lex.GetToken()
        parse.Simpleexpr()
    }
}

//SIMPLEEXPR ::= TERM { ADDOP TERM }
func (parse *Parser) Simpleexpr(){
    parse.Tt("simpleexpr " + parse.Lex.Lexeme)
    parse.Term()
    for parse.IsAddop(parse.Token) {
        parse.Token,_ = parse.Lex.GetToken()
        parse.Term()
    }
}

//TERM ::= FACTOR {MULTOP FACTOR }
func (parse *Parser) Term(){
    parse.Tt("term " + parse.Lex.Lexeme)
    parse.Factor()
    for parse.IsMultop(parse.Token) {
        parse.Token,_ = parse.Lex.GetToken()
        parse.Factor()
    }
}

//FACTOR ::= -FACTOR | not FACTOR
//FACTOR ::= number | false | true
//FACTOR ::= id | id(ARGLIST) | (EXPR)
func (parse *Parser) Factor(){
    parse.Tt("factor " + parse.Lex.Lexeme)
    if parse.BeginsFactor(parse.Token) {
        if parse.Token == lexer.T_id {
            parse.Accept(lexer.T_id)
            if parse.Token == lexer.T_lparen {
                parse.Accept(lexer.T_lparen)
                parse.Arglist()
                parse.Accept(lexer.T_rparen)
            }
        } else if parse.Token == lexer.T_minus || parse.Token == lexer.T_not {
            parse.Accept(parse.Token)
            parse.Factor()
        } else if parse.Token == lexer.T_lparen {
            parse.Accept(lexer.T_lparen)
            parse.Expr()
            parse.Accept(lexer.T_rparen)
        } else {
            parse.Accept(parse.Token)
        }
    } else {
        parse.StringErrorMessage("Invalid factor token: " + strconv.Itoa(parse.Token))
    }
}

//ARGLIST ::= EXPR {, EXPR } | E
func (parse *Parser) Arglist(){
    parse.Tt("arglist " + parse.Lex.Lexeme)
    if parse.BeginsFactor(parse.Token) {
        parse.Expr()
        for parse.Token == lexer.T_comma {
            parse.Accept(lexer.T_comma)
            parse.Expr()
        }
    }
}

func (parse *Parser) IsRelop(t int) bool{
    return t == lexer.T_eq || t == lexer.T_ne || t == lexer.T_le || t== lexer.T_lt || t == lexer.T_ge || t == lexer.T_gt
}

func (parse *Parser) IsAddop(t int) bool{
    return t == lexer.T_plus || t == lexer.T_minus || t == lexer.T_or
}

func (parse *Parser) IsMultop(t int) bool{
    return t == lexer.T_mult || t == lexer.T_div || t == lexer.T_mod || t== lexer.T_and
}

func (parse *Parser) Tt(str string){
    fmt.Println("\n" + str + "\n")
}
