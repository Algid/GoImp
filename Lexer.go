package Lexer

import(
    "bufio"
    "byte"
    "io"
    "fmt"
)

type LexIO struct{
    Source Reader
    Listing Writer
   
    Lexeme string
}
func GetChar() char {
}
