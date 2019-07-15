package main

import (
    "fmt"
    "encoding/json"
)

var exists = struct{}{}

type set struct {
    m map[string]struct{}
}

func NewSet() *set {
    s := &set{}
    s.m = make(map[string]struct{})
    return s
}

func (s *set) Add(value string) {
    s.m[value] = exists
}

func (s *set) Remove(value string) {
    delete(s.m, value)
}

func (s *set) Contains(value string) bool {
    _, c := s.m[value]
    return c
}

func robotSim(commands []int, obstacles [][]int) int {
    var result, dir, x, y int = 0, 0, 0, 0
    var dirX = []int{0,1,0,-1}
    var dirY = []int{1,0,-1,0}

    obstaclesSet := NewSet()
    for _, obs := range obstacles {
        tmp, _ := json.Marshal(obs)
        obstaclesSet.Add(string(tmp))
    }

    for  _, cmd := range commands{
        switch(cmd){
            case -2:
                dir = (dir - 1) % 4
            case -1:
                dir = (dir + 1) % 4
            default:
                for j:=0; j < cmd; j++ {
                    tmpArr := make([]int, 2)
                    tmpArr[0] = x + dirX[dir]
                    tmpArr[1] =  y + dirY[dir]
                    tmp, _ := json.Marshal(tmpArr)
                    if ! obstaclesSet.Contains(string(tmp)) {
                        x += dirX[dir]
                        y += dirY[dir]
                        max := x*x + y*y
                        if max > result {
                            result = max
                        }
                    }
                }
        }
    }

    return result
}

func main() {
    fmt.Printf("%v", robotSim([]int{4,-1,3}, [][]int{}));
    fmt.Println("\n")
    fmt.Printf("%v", robotSim([]int{4,-1,4,-2,4}, [][]int{{2,4}}));
}