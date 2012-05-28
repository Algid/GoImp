package main

import(
    "fmt"
    "os"
    "io"
    "strings"
    "GoImp"
)

func main(){
    var (
        str *strings.Reader
        src io.Reader
        file *os.File
        err error
        list io.Writer
        token int
    )
    str = strings.NewReader("x _ _x and begin boolean break call end else false\nbegin function half if input integer is loop not null\n newline or output procedure return then true var\n123,:()< <= > >= > = == # + - * / %!<\n \"This is a string token\" mom is nice\n so is dad //but little sis is naughty x\n y\nx \n bye ")
    src = io.Reader(str)
    if file, err = os.Create("./list.imp"); err != nil{
        fmt.Println(err)
    }
    defer file.Close()
    list = io.Writer(file)
    lex := lexer.New(&src, &list)
    token,_ = lex.GetToken()
    for token != lexer.T_eof {
        fmt.Println(lex.TokenToStringVector[token] + ":" + lex.Lexeme)
        token,_ = lex.GetToken()
    }
}
