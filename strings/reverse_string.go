package main

import (
        "fmt"
        )

func main(){
    var s []rune = []rune("ABCs1DEF")
    var backup []rune = make([]rune, len(s))
    copy(backup, s)
    for i:=0; i<len(s)/2;i++ {
        s[i] ^= s[len(s)-i-1]
        s[len(s)-i-1] ^= s[i]
        s[i] ^= s[len(s)-i-1]
    }

    fmt.Printf("origin:%s, reverse string is %s\n", string(backup), string(s))
}
