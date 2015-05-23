package main

import (
    "fmt"
    "os/exec"
    "strconv"
)


func main() {
    n := 0
    fmt.Scan(&n)
    for ctr :=0; ctr < 10; ctr++ {
        uri := "http://it-ebooks.info/book/"
        n += 1
        pth := ""
        pth = strconv.Itoa(n)
        uri += pth
        exec.Command(`C:\Windows\System32\rundll32.exe`, "url.dll,FileProtocolHandler", uri).Start()
    }
}
