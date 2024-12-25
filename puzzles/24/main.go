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
	values, gates:= parseInput(name)
	output:= 0
	for {
		toDelete:= [][2]string{}
		for inputs,gate:= range gates {
			if v1,ok:= values[inputs[0]]; ok {
				if v2,ok2:=values[inputs[1]];ok2 {
					for g:=0; g<len(gate);g++ {
						val:= calcGate(v1,v2,gate[g].operation)
						outputTo:= gate[g].out
					
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
		if len(toDelete)==0 {
			return output
		}
		for _,v:= range toDelete {
			delete(gates,v)
		}
	}
}

func part2(name string) string {
	values, gates:= parseInput(name)
	wantedOutput:= getWantedOutput(values)
	// swaps := [][2]string{}
	// fmt.Println(findAllBadGates(gates))
	for z:=3; z<=45; z++ {
		fmt.Println(z)
		fmt.Println(trySwaps(wantedOutput,values,gates,[][2]string{},z))
		
		badGates:= getBadGates(gates,z)
		if len(badGates) ==0 {continue}
		fmt.Println(badGates)
		fmt.Println(trySwaps(wantedOutput,values,gates,[][2]string{{badGates[0],badGates[1]}},z))
	}
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

func getWantedOutput(inputs map[string]bool) map[int]bool {
	res:=0
	resMap:= make(map[int]bool)
	for k,v:=range inputs {
		if k[0]=='x' || k[0]=='y' {
			vInt:= 0
			if v {vInt=1}
			res += vInt<<ints.FromString(k[1:])
		}
	}
	i:=0 
	for res>0 {
		resMap[i] = res%2 ==1
		i++
		res /= 2
	}

	return resMap
}

func trySwaps(wantedOutput map[int]bool, oldValues map[string]bool, oldGates map[[2]string][]gate, swaps [][2]string, valToCheck int) bool {
	outputsSoFar:=[]string{}
	values:= make(map[string]bool)
	gates:=make(map[[2]string][]gate)
	for k,v:= range oldValues {	values[k] = v}
	for k,v:= range oldGates {	gates[k] = v}
	for {
		toDelete:= [][2]string{}
		for inputs,gate:= range gates {
			if v1,ok:= values[inputs[0]]; ok {
				if v2,ok2:=values[inputs[1]];ok2 {
					for g:=0; g<len(gate);g++ {
						val:= calcGate(v1,v2,gate[g].operation)
						outputTo:= swapsCheck(gate[g].out,swaps)
						if slices.Contains(outputsSoFar, outputTo) {
							return false
						}
						outputsSoFar = append(outputsSoFar, outputTo)
						values[outputTo] = val
						if outputTo[0] == 'z' {
							outNum:= ints.FromString(outputTo[1:])
							if outNum==valToCheck {
								return wantedOutput[outNum] == val
							}
						}
						
						toDelete = append(toDelete, inputs)
						}
				}
			}
		}
		if len(toDelete)==0 {
			return false
		}
		for _,v:= range toDelete {
			delete(gates,v)
		}
	}
}

func swapsCheck(out string, swaps [][2]string) string {
	for _,s:= range swaps {
		if out == s[0] {
			return s[1]
		}
		if out == s[1] {
			return s[0]
		}
	}
	return out
}


func getBadGates(gateMap map[[2]string][]gate, zToCheck int) []string {
	lastGateKey, lastGate := findLastGate(gateMap,zToCheck)
	res:= []string{}
	if lastGate.operation != "XOR" {
		res = append(res, lastGate.out)
	}

	for k,v:= range gateMap {
		for _,gate:=range v {
			if gate.out == lastGateKey[0] || gate.out == lastGateKey[1] {
				if gate.operation == "XOR" {
					if !(k[0][0] == 'x' && k[1][0]=='y') &&
					!(k[0][0] == 'y' && k[1][0]=='x') {
						res = append(res, gate.out)
					}
					continue
				}
				if gate.operation == "OR" {
					found:= false
					for _,v2:= range gateMap {
						for _,gate2:=range v2 {
							if (gate2.out == k[0] || gate2.out == k[1]) && gate2.operation != "AND"{
								res = append(res, gate2.out)
							}
						}
						if found { break}
					}
					continue
				}
				res = append(res, gate.out)
			}
		}
	}

	return res
}

func findLastGate(gateMap map[[2]string][]gate, zToCheck int) ([]string,gate) {
	for k,v:=range gateMap {
		for _,gate:=range v {
			if gate.out[0] == 'z' && ints.FromString(gate.out[1:]) == zToCheck {
				return []string{k[0],k[1]},gate
			}
		}
	}
	panic("lastGate not found")
}

// func findAllBadGates(gateMap map[[2]string][]gate) map[[2]string][]gate {
// 	for k,v:= range gateMap {
// 		for _,gate:= range v{
// 			var nextGates []gate
// 			for k2,v2:= range gateMap {
// 				if k2[0] == gate.out || k2[1] == gate.out {
// 					nextGates = append(nextGates, v2...)
// 				}
// 			}
// 			switch gate.operation {
// 			case "AND":
// 				if gat
// 			case "XOR":

// 			case "OR":

// 			}
// 		}
// 	}
// }