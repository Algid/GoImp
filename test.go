package main

import(
    "os"
    "fmt"
    "bytes"
    "io"
    "bufio"
    "strings"
    "GoImp/lexer"
)

func main(){
    r := new(io.Reader)
    w := new(io.Writer)
    lex := lexer.New(r, w)
    fmt.Println(lex.t_while)
}

func readLines(path String) (lines []string, err error){
    var (
        file *os.File
        part []byte
        prefix bool
    )
    if file, err = os.Open(path); err != nil {
        return
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    buffer := bytes.NewBuffer(make([]byte, 0))

}
    
