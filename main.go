package main

import (
        "fmt"
        "io/ioutil"
)

type Pointer struct{
        Position    int
}

func check(e error){
        if e != nil{
                panic(e)
        }
}


func main() {
        file, err := ioutil.ReadFile("./test")
        check(err)
        p := Pointer{15000}
        nodes := make([]int, 30000)
        jmp_index := 0
        index := 0
        for index < len(file) {
                if file[index] == 62 {
                        p.Position += 1
                } else if file[index] == 60{
                        p.Position -= 1
                } else if file[index] == 43{
                        nodes[p.Position] += 1
                } else if file[index] == 45{
                        nodes[p.Position] -= 1
                } else if file[index] == 91{
                        jmp_index = index
                } else if file[index] == 46{
                        fmt.Printf("%c", nodes[p.Position])
                } else if file[index] == 93{
                        if nodes[p.Position] != 0{
                                index = jmp_index
                                continue
                        }
                }
                index += 1
        }
}
