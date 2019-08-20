package main

import (
        "fmt"
        "io/ioutil"
)

type Pointer struct{
        Position    int
}

func getArraySize(file []byte) (int, int){ //returns the needed size of the array
        var highest int                    //also returns the initial starting position of the pointer
        var lowest int
        var current int
        var starting_pos int
        for _, value := range file {
                if value == 62 {
                        current += 1
                        if current > highest{
                                highest = current
                        }
                } else if value == 60{
                        current -= 1
                        if current < lowest{
                                starting_pos += 1
                                lowest = current
                        }
                }
        }
        //fmt.Println(highest)
        //fmt.Println(lowest)
        //fmt.Println(highest - lowest)
        //fmt.Println(starting_pos)
        return highest - lowest + 1, starting_pos
}

func check(e error){
        if e != nil{
                panic(e)
        }
}

func main() {
        file, err := ioutil.ReadFile("./test")
        check(err)
        array_size, starting_position := getArraySize(file)
        p := Pointer{starting_position}
        nodes := make([]int, array_size)
        for _, value := range file {
                if value == 62 {
                        p.Position += 1
                } else if value == 60{
                        p.Position -= 1
                } else if value == 43{
                        nodes[p.Position] += 1
                } else if value == 45{
                        nodes[p.Position] -= 1
                }
        }
        fmt.Println(nodes)
}
