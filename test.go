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
    )
    str = strings.NewReader("Foo Bar")
    src = io.Reader(str)
    if file, err = os.Create("./list.imp"); err != nil{
        fmt.Println(err)
    }
    defer file.Close()
    list = io.Writer(file)
    lex := lexer.New(&src, &list)
    fmt.Println(lex.GetChar())
    fmt.Println(lex.GetChar())
}
