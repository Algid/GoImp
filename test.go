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
        char string
        src io.Reader
        file *os.File
        err error
        list io.Writer
    )
    str = strings.NewReader("Foo Bar 123\nwolololol")
    src = io.Reader(str)
    if file, err = os.Create("./list.imp"); err != nil{
        fmt.Println(err)
    }
    defer file.Close()
    list = io.Writer(file)
    lex := lexer.New(&src, &list)
    char,_ = lex.GetChar()
    for char != "" {
        fmt.Println(char)
        char,_ = lex.GetChar()
    }
}
