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
        p := Pointer{}
        nodes := make([]int, 30000)
        jmp_index := 0
        for index, value := range file {
                if value == 62 {
                        p.Position += 1
                } else if value == 60{
                        p.Position -= 1
                } else if value == 43{
                        nodes[p.Position] += 1
                } else if value == 45{
                        nodes[p.Position] -= 1
                } else if value == 91{
                        jmp_index =  index
                        fmt.Printf("Jmp_index: %d\n", jmp_index)
                } else if value == 93{
                        fmt.Printf("node value %d\n", nodes[p.Position])
                        if nodes[p.Position] != 0{
                                p.Position = jmp_index
                                fmt.Printf("Looped back")
                        }
                }
        }
}
