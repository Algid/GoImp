package Parser

import "fmt"
import "os"
import "GoImp"

func parse()
{
    token = getToken()
    program()
}

func accept(TokenType t)
{
    if token == t {
        token = getToken()
    }
    else {
        errorMessage(t, token)
    }
}

func errorMessage(expected, found TokenType)
{
    fmt.Println("\n\n I was expecting a " + lex.tokenToString(expected) + " but I found a " + lex.tokenToString(found) + " " + lex.getLexeme())
    os.exit(1)
}

func errorMessage(message string)
{
    fmt.Println("\n" + message)
    os.exit(1)
}

//PROGRAM ::= VARDECS SUBROUTINES
func program()
{
    vardecs()
    subroutines()
}

//VARDECS ::= var VARDECLIST { VARDECLIST } | empty
func vardecs()
{
    //TRACER t("vardecs " + lex.getLexeme())
    if token != t_var {return}
    accept(t_var)
    vardeclist()
    for ; token == t_id; {
        vardeclist()
    }
}

//VARDECLIST ::= id {, id } : TYPE
func vardeclist()
{
    //TRACER t("vardeclist " + lex.getLexeme())
    accept(t_id)

    if token == t_colon {
        accept(t_colon)
        type()
        accept(t_semi)
        if token == t_id {
            vardeclist()
        }
    }

    if token == t_comma {
        accept(t_comma)
    }
}

//TYPE ::= integer | boolean
func type()
{
    //TRACER t("type " + lex.getLexeme())
    if token == t_integer {
        accept(t_integer)
    }

    else if token == t_boolean {
        accept(t_boolean)
    }

    else {errorMessage(t_boolean,token)}
}

//SUBROUTINES ::= PROCFUN { PROCFUN }
func subroutines()
{
    //TRACER t("subroutines " + lex.getLexeme())
    for ; token == t_procedure || token == t_function; {
        procfun()
    }
}

//PROCFUN ::= PROC | FUN
func procfun()
{
    //TRACER t("procfun " + lex.getLexeme())
    if token == t_procedure {
        proc()
    }

    if token == t_function {
        fun()
    }
}

//PROC ::= procedure id ( PARAMETERS ) is PROCBODY
func proc()
{
    //TRACER t("proc " + lex.getLexeme())
    accept(t_procedure)
    accept(t_id)
    accept(t_lparen)
    parameters()
    accept(t_rparen)
    accept(t_is)
    procbody()
}

//FUN ::= function id ( PARAMETERS ) return TYPE is PROCBODY
func fun()
{
    //TRACER t("fun " + lex.getLexeme())
    accept(t_function)
    accept(t_id)
    accept(t_lparen)
    parameters()
    accept(t_rparen)
    accept(t_return)
    type()
    accept(t_is)
    procbody()
}

//PARAMETERS ::= PARAMLIST { PARAMLIST} | E
func parameters()
{
    //TRACER t("parameters " + lex.getLexeme())
    if token == t_rparen {return}
    paramlist()
}

//PARAMLIST ::= id {, id} : TYPE
func paramlist()
{
    //TRACER t("paramlist " + lex.getLexeme())
    if token != t_id {
        errorMessage(t_id,token)
    }

    for ; token == t_id; {
        accept(t_id)
        if token == t_comma {
            accept(t_comma)
        }
    }
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
