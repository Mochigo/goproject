package main

import (
    "fmt"
       )

func factorial(n int64) int64 {
     if n > 1 {
         n *= factorial(n - 1)
     } else if n == 1 || n == 0 {
         n *= 1
     }
     return n
}

func main() {
     var n int64
     
     fmt.Printf("请输入一个整数：");
     fmt.Scanf("%d", &n)
     fmt.Printf("该整数的阶乘为：%d。", factorial(n))
}




   
