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

func getBracketPairs(f []byte) map[int]int{
        opening_indexes := []int{}
        bracket_map := make(map[int]int)
        for index, value := range f{
                if value == 91{
                     opening_indexes = append(opening_indexes, index)
                }
                if value == 93{
                        if len(opening_indexes) > 0{
                                bracket_map[index] = opening_indexes[len(opening_indexes)-1]
                                opening_indexes = opening_indexes[:len(opening_indexes)-1]
                        }
                }
        }
        return bracket_map
}

func main() {
        file, err := ioutil.ReadFile("./test")
        check(err)
        p := Pointer{15000}
        nodes := make([]int, 30000)
        index := 0
        bracketMap := getBracketPairs(file)
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
                        fmt.Printf("")
                } else if file[index] == 46{
                        fmt.Printf("%c", nodes[p.Position])
                } else if file[index] == 93{
                        if nodes[p.Position] != 0 {
                                index = bracketMap[index]
                        }
                }
                //fmt.Println(p.Position)
                index += 1
        }
}
