package main

import (
        "fmt"
        )

func main(){
    var s string = "abcdefgha"
    var checker int = 0
    var flag bool
    for i:=0; i< len(s); i++ {
        var index uint = uint(s[i]) - uint('a')
        if (index <0 || index >26 || (checker & (1 << index )) > 0) {
           flag = true
        }
        checker |= (1 << index)
        fmt.Printf("char:%s,index:%d,1<<index:%d,checker:%d\n",
            string(s[i]), index, (1<<index), checker)
    }
    if flag {
        fmt.Printf("print string:%s is not uniq\n", s)
    } else {
        fmt.Printf("print string:%s is uniq\n", s)
    }
}
