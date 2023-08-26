package main


import (

"fmt"
"strings"
)


func Compare(a, b string) string {
    var short, long, res string
    if len(a) < len(b){
        short,long = a,b
    }else{
        short,long = b,a
    }
    for i := 0; i<len(short); i++{
        for j:=len(short); j>=i; j--{
            if strings.Contains(long, short[i:j]) && len(short[i:j]) > len(res){
              res = short[i:j]
            }
        }
    }
  return res
}


func main() {
    fmt.Println(Compare("AKA", "AKASHI")) //AKA
    fmt.Println(Compare("KANGOORO", "KANG")) //KANG
    fmt.Println(Compare("KI", "KIJANG")) //KI
    fmt.Println(Compare("KUPU-KUPU", "KUPU")) //KUPU
    fmt.Println(Compare("ILALANG", "ILA")) //ILA
}