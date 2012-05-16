package main

import(
    "fmt"
    "GoImp"
)

func main(){
    lex := lexer.New()
    for i := 0; i < 10; {
        fmt.Println(lex.GetChar())
        i++
    }
}
