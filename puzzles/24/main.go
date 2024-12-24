package main

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
	"github.com/PhilAldridge/aoc-2024-go/pkg/ints"
)

func main() {
	start := time.Now()
	fmt.Println("Part 1: ", part1("input.txt"), " in:", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", part2("input.txt"), " in:", time.Since(start))
}

func part1(name string) int {
	return calcPart1(name, [][2]string{})
}

func part2(name string) string {
	// outs:= parseOuts(name)
	// ins,_:= parseInput(name)
	// fmt.Println(len(outs))
	// wantedOutput:= getWantedOutput(ins)
	// for a:=0; a<len(outs); a++ {
	// 	for b:=a+1; b<len(outs); b++ {
	// 		for c:=b+1; c<len(outs); c++ {
	// 			fmt.Println(c)
	// 			for d:=c+1; d<len(outs); d++ {
	// 				for e:= d+1; e<len(outs); e++ {
	// 					for f:= e+1; f<len(outs); f++ {
	// 						for g:=f+1; g<len(outs); g++ {
	// 							for h:=g+1; h<len(outs); h++ {
	// 								if calcPart1(name,[][2]string{
	// 									{outs[a],outs[b]},
	// 									{outs[c],outs[d]},
	// 									{outs[e],outs[f]},
	// 									{outs[g],outs[h]},
	// 								}) == wantedOutput {
	// 									return outs[a] + "," +
	// 										outs[b] + "," +
	// 										outs[c] + "," +
	// 										outs[d] + "," +
	// 										outs[e] + "," +
	// 										outs[f] + "," +
	// 										outs[g] + "," +
	// 										outs[h] + "," 
	// 								}
	// 							}
	// 						}
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}
	// }
	return ""
}

type gate struct {
	out string
	operation string
}

func parseInput(name string) (map[string]bool, map[[2]string][]gate) {
	parts:= files.ReadParagraphs(name)
	inputs:= make(map[string]bool)
	gates:= make(map[[2]string][]gate)
	for _,v:= range parts[0] {
		inputVal:= strings.Split(v,": ")
		inputs[inputVal[0]] = inputVal[1]=="1"
	}
	for _,g:= range parts[1] {
		gateVal:= strings.Split(g," ")
		gates[[2]string{gateVal[0],gateVal[2]}] = append(gates[[2]string{gateVal[0],gateVal[2]}],
			gate{
			out: gateVal[4],
			operation: gateVal[1],
		})
	}
	return inputs,gates
}

func parseOuts(name string) []string {
	parts:= files.ReadParagraphs(name)
	outs:= make(map[string]bool)
	for _,g:= range parts[1] {
		gateVal:= strings.Split(g," ")
		outs[gateVal[4]]=true
	}
	// create an empty slice of key-value pairs
	s := make([]string, 0, len(outs))
	// append all map keys-value pairs to the slice
	for k := range outs {
		s = append(s, k)
	}
	slices.Sort(s)
	return s
}

func calcGate(v1 bool, v2 bool, operation string) bool {
	switch operation {
	case "OR":
		return v1 || v2
	case "XOR":
		return v1!=v2
	case "AND":
		return v1 && v2
	}
	panic("bad operation")
}

func parseOut(out bool, outVal int) int {
	val:= 0
	if out { val = 1}
	return val << outVal
}

func getWantedOutput(inputs map[string]bool) int {
	res:=0
	for k,v:=range inputs {
		if k[0]=='x' || k[0]=='y' {
			vInt:= 0
			if v {vInt=1}
			res += vInt<<ints.FromString(k[1:])
		}
	}
	return res
}

func calcPart1(name string, swaps [][2]string) int {
	values, gates:= parseInput(name)
	output:= 0
	for {
		finished:= true
		toDelete:= [][2]string{}
		for inputs,gate:= range gates {
			if v1,ok:= values[inputs[0]]; ok {
				if v2,ok2:=values[inputs[1]];ok2 {
					finished = false
					for g:=0; g<len(gate);g++ {
						val:= calcGate(v1,v2,gate[g].operation)
						outputTo:= gate[g].out
						for _,s := range swaps {
							if s[0] == outputTo {
								outputTo = s[1]
								break
							}
							if s[1] == outputTo {
								outputTo = s[0]
								break
							}
						}
						values[outputTo] = val
						if outputTo[0]=='z' {
							outNum:= ints.FromString(outputTo[1:])
							output  += parseOut(val,outNum)
						}
						
					}
					toDelete = append(toDelete, inputs)
				}
			}
		}
		if finished {
			return output
		}
		for _,v:= range toDelete {
			delete(gates,v)
		}
	}
}