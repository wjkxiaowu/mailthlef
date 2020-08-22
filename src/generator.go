package main

import (
    "fmt"
    "io/ioutil"
    "encoding/base64"
)

func main() {
    b, _ := ioutil.ReadFile("./release/mail")
    encodeString := base64.StdEncoding.EncodeToString(b)
    fmt.Printf(`var enStr string = "%s"`, encodeString)
}
