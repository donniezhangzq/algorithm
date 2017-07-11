/* 
author: donniezhangzq@gmail.com
last chagne date: 2017-07-10
// to judge that if all characters is uniq in a string
*/
package main

import (
    "fmt"
        )

func main(){
    var input string
    var m map[rune]int
    m = make(map[rune]int)
    var flag bool
    fmt.Println("Please input string:")
    fmt.Scanln(&input)
    var s = []rune(input)
    for k,val := range(s) {
        m[val]++
        if (m[val] > 1){
            fmt.Printf("string:%s is not uniq,it has duplicate char:%s",
                input, string(s[k]))
            flag = true
            break
        }
    }
    if (!flag) {
        fmt.Printf("string:%s's charaters is uniq", input)
    }
}
