package main

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
)

func main() {
	start := time.Now()
	fmt.Println("Part 1: ", part1("input.txt"), " in:", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", part2("input.txt"), " in:", time.Since(start))
}

func part1(name string) int {
	connectionMap:= getConnectionMap(name)
	threeList:= make(map[string]int)
	for k,v:= range connectionMap {
		if k[0]!='t' || len(v.connections)<=2 {
			continue
		}
		for i:= 0 ; i<len(v.connections); i++ {
			for j:= i+1; j<len(v.connections); j++ {
				if !slices.Contains(connectionMap[v.connections[i]].connections,v.connections[j]) {
					continue
				}
				three:= []string{k,v.connections[i],v.connections[j]}
				slices.Sort(three)
				threeList[strings.Join(three,",")] ++
			}
		}
	}
	return len(threeList)
}

func part2(name string) string {
	connectionMap:= getConnectionMap(name)
	res:=""
	maxFound:= 2
	for k,v:= range connectionMap {
		if len(v.connections)<maxFound {
			continue
		}
		computerList:= []string{}
		for _,conn:= range v.connections {
			connected:= true
			for _,comp:= range computerList {
				if !slices.Contains(connectionMap[conn].connections,comp) {
					connected = false
					break
				}
			}
			if connected {
				computerList = append(computerList, conn)
			}
		}
		computerList = append(computerList,  k)
		if len(computerList) > maxFound {
			maxFound = len(computerList)
			slices.Sort(computerList)
			res = strings.Join(computerList,",")
		}
	}
	return res
}

type computer struct {
	connections []string
}

func getConnectionMap(name string) map[string]computer {
	lines:= files.ReadLines(name)
	mapping:= make(map[string]computer)

	for _,l:= range lines{
		computers:= strings.Split(l, "-")
		if entry,ok:= mapping[computers[0]];ok {
			exists:= false
			for _,c:=range entry.connections {
				if computers[1] == c {
					exists = true
					break
				}
			}
			if !exists {
				entry.connections = append(entry.connections, computers[1])
				mapping[computers[0]] = entry
			}
		} else {
			mapping[computers[0]] = computer{connections: []string{computers[1]}}
		}

		if entry,ok:= mapping[computers[1]];ok {
			exists:= false
			for _,c:=range entry.connections {
				if computers[0] == c {
					exists = true
					break
				}
			}
			if !exists {
				entry.connections = append(entry.connections, computers[0])
				mapping[computers[1]] = entry
			}
		} else {
			mapping[computers[1]] = computer{connections: []string{computers[0]}}
		}
	}
	
	return mapping
}