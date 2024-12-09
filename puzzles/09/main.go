package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
	"github.com/kylehoehns/aoc-2023-go/pkg/files"
	"github.com/kylehoehns/aoc-2023-go/pkg/ints"
)

func main() {
	start := time.Now()
	fmt.Println("Part 1: ", part1("input.txt"), " in:", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", part2("input.txt"), " in:", time.Since(start))
}

func part1(name string) int {
	fileSystem:= createFileSystem(name)
	l:=0
	r:=len(fileSystem)-1
	for l<r {
		//Move left pointer until there is a space
		if fileSystem[l] != -1 {
			l++
			continue
		}
		//Move right pointer until there is a file to move
		if fileSystem[r] ==-1 {
			r--
			continue
		}
		//Swap space at left pointer and file at right pointer
		fileSystem[l],fileSystem[r] = fileSystem[r], fileSystem[l]
	}
	return checkSum(fileSystem)
}

func part2(name string) int {
	fileSystem:= createFileSystem(name)
	r:=len(fileSystem)-1
	lastFileNum := 1000000
	for r >0 && fileSystem[r] != 0  {
		//skip until there is a file not yet moved
		if fileSystem[r] ==-1 || fileSystem[r] >= lastFileNum {
			r--
			continue
		}
		fileNum:= fileSystem[r]
		lastFileNum = fileNum
		fileLength := findFileLength(fileSystem,r)
		emtpySpaceIndex,err := findSpaceToFill(fileSystem,fileLength)
		
		//nowhere to move the file, so skip it
		if err!=nil || emtpySpaceIndex > r {
			r -= fileLength
			continue
		}

		//swap file with empty space found
		for i:=0; i<fileLength; i++ {
			fileSystem[emtpySpaceIndex+i] = fileNum
			fileSystem[r-i] = -1
		}
		
		r -= fileLength-1
	}
	return checkSum(fileSystem)
}

func createFileSystem(name string) []int {
	fileString := files.Read(name)
	fileString = strings.ReplaceAll(fileString,"\n","")
	fileSystem:= []int{}
	for i:=0; i<len(fileString); i+=2 {
		file:= ints.FromString(string(fileString[i]))
		for j:=0; j<file; j++ {
			fileSystem = append(fileSystem, i/2)
		}
		if i+1 == len(fileString) {
			break
		}
		space:= ints.FromString(string(fileString[i+1]))
		for j:=0; j<space;j++ {
			fileSystem = append(fileSystem, -1)
		}
	}
	return fileSystem
}

func findFileLength(fileSystem []int, rIndex int) int {
	counter:=0
	for fileSystem[rIndex-counter] == fileSystem[rIndex] && counter <= rIndex {
		counter ++
	}
	return counter
}

func findSpaceToFill(fileSystem []int, fileLength int) (int,error) {
	for i:=0; i<len(fileSystem); i++ {
		if fileSystem[i] != -1 {
			continue
		}
		for j:=1; j<len(fileSystem)-i; j++ {
			if j>= fileLength {
				return i,nil
			}
			if fileSystem[i+j]!= -1 {
				i = i+j
				break
			}
		}
	}
	return -1, errors.New("no space left")
}

func checkSum(fileSystem []int) int {
	total:= 0
	for i,v:= range fileSystem {
		if v==-1 {
			continue
		}
		total += i*v
	}
	return total
}