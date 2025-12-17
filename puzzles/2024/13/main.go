package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/coords"
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
	machines := getMachines(name)
	total := 0
	for _, machine := range machines {
		total += machineVal(machine)
	}
	return total
}

func part2(name string) int {
	machines := getMachines(name)
	total := 0
	toAdd := coords.NewCoord(10000000000000, 10000000000000)
	for _, machine := range machines {
		machine.target = machine.target.Add(toAdd)
		total += machineVal(machine)
	}
	return total
}

type machine struct {
	a      coords.Coord
	b      coords.Coord
	target coords.Coord
}

func getMachines(name string) []machine {
	machineStr := files.ReadParagraphs(name)
	machines := []machine{}
	for _, str := range machineStr {
		aSplit := strings.Split(strings.ReplaceAll(strings.ReplaceAll(str[0], "Button A: X", ""), "+", ""), ", Y")
		a := coords.NewCoord(ints.FromString(aSplit[0]), ints.FromString(aSplit[1]))
		bSplit := strings.Split(strings.ReplaceAll(strings.ReplaceAll(str[1], "Button B: X", ""), "+", ""), ", Y")
		b := coords.NewCoord(ints.FromString(bSplit[0]), ints.FromString(bSplit[1]))
		tSplit := strings.Split(strings.ReplaceAll(strings.ReplaceAll(str[2], "Prize: X=", ""), "+", ""), ", Y=")
		t := coords.NewCoord(ints.FromString(tSplit[0]), ints.FromString(tSplit[1]))
		machines = append(machines, machine{
			a:      a,
			b:      b,
			target: t,
		})

	}
	return machines
}

func machineVal(machine machine) int {
	aVal := 3
	bVal := 1
	if machine.a.IsSameDirectionAs(machine.b) {
		//Parallel
		//None of the inputs ended up being  parallel, so this code is redundant. Left in for completeness
		if machine.a.J > 3*machine.b.J {
			//Pressing a is cheaper than pressing b. Guess and check starting with max possible a presses.
			for j:= machine.target.J/machine.a.J; j>=0; j-- {
				distanceLeft:= machine.target.J - j*machine.a.J
				for i:=0; i<=distanceLeft/machine.b.I; i++ {
					bDist:= machine.b.I * i
					if distanceLeft < bDist {
						break
					}
					if distanceLeft == bDist && machine.target.I == machine.a.I*j + machine.b.I*i {
						return aVal*j + bVal*i
					}
				}
			} 
		} else {
			//Pressing b is cheaper than pressing a. Guess and check starting with max possible b presses.
			for i := machine.target.J / machine.b.J; i >= 0; i-- {
				distanceLeft := machine.target.J - i*machine.b.J
				for j := 0; j <= distanceLeft/machine.a.J; j++ {
					aDist := machine.a.J * j
					if distanceLeft < aDist {
						break
					}
					if distanceLeft == aDist && machine.target.I == machine.a.I*j + machine.b.I*i {
						return aVal*j + bVal*i
					}
				}
			}
		}
	} else {
		//Not parallel - only one solution
		//To solve, Multiply by inverse matrix
		det := machine.a.I*machine.b.J - machine.a.J*machine.b.I
		aAmount := machine.b.J*machine.target.I - machine.b.I*machine.target.J
		bAmount := machine.a.I*machine.target.J - machine.a.J*machine.target.I
		if aAmount%det == 0 && bAmount%det == 0 && aAmount/det>=0 && bAmount/det>=0 {
			return aVal*aAmount/det + bVal*bAmount/det
		}
	}
	return 0
}
