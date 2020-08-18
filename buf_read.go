package main

import (
"os"
"bufio"
"fmt"
"time"
)
 

func main(){
    file,err:=os.Open("a.txt")
    if err!=nil {
        fmt.Println(err)
    }
    defer file.Close()

    //创建一个新的io.Reader，它实现了Read方法
    reader:=bufio.NewReader(file)
    //设置读取的长度
    buf:=make([]byte,19)
    //读取文件
for {
    _,err=reader.Read(buf)
    if err!=nil {
        fmt.Println(err)
    }
time.Sleep(100 * time.Millisecond)
//   fmt.Print(string(buf))
    fmt.Print(buf)
}
}
