package main
import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    var i string
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan(){
        i = scanner.Text()
        fmt.Println("Hello, World.")
        fmt.Println(i)
    }
}
