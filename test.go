package main

import(
    "fmt"
    "os"
    "io"
    "bufio"
    "GoImp/parser"
)

func main(){
    var (
        src io.Reader
        infile *os.File
        outfile *os.File
        err error
        list io.Writer
    )
    fmt.Println("What is the input file?")
    input := bufio.NewReader(os.Stdin)
    str, err := input.ReadString('\n')

    if infile, err = os.Open("./ImpTestFiles/" + str[0:len(str)-1] + ".imp"); err != nil{
        fmt.Println(err)
        os.Exit(1)
    }
    defer infile.Close()
    src = io.Reader(infile)

    if outfile, err = os.Create("./ImpTestFiles/" + str[0:len(str)-1] + ".out"); err != nil{
        fmt.Println(err)
        os.Exit(1)
    }
    defer outfile.Close()
    list = io.Writer(outfile)
    parse := parser.New(&src, &list)
    parse.Parse()
}
