package main

import (
	"fmt"
	"strings"
)

func main() {

	words := []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
	k := 16

	currentGroup := []string{words[0]}
	cnt := len(words[0]) // # of letters in the current group build
	groupings := [][]string{}

	for x := 1; x < len(words); x++ {
		if len(words[x])+1+cnt > k {
			// we are starting a new group
			groupings = append(groupings, currentGroup)
			fmt.Println(currentGroup)
			currentGroup = []string{}
			cnt = 0
		}
		cnt += len(words[x]) + 1
		currentGroup = append(currentGroup, words[x])
	}

	if len(currentGroup) > 0 { // last group
		fmt.Println(currentGroup)
		groupings = append(groupings, currentGroup)
	}

	//////////// now we have groupings
	fmt.Println("-------------------")
	fmt.Println(groupings)

	for x := 0; x < len(groupings); x++ {
		getGroupString(groupings[x], k)
	}

}

func getGroupString(sgroup []string, k int) {
	fmt.Printf("Welcome to getGroupString()  sgroup=%v  k=%v\n", sgroup, k)

	gaps := len(sgroup) - 1
	letterCount := 0
	for _, word := range sgroup {
		letterCount += len(word)
	}

	fmt.Printf("letters=%d\n", letterCount)

	spaceCount := k - letterCount
	fmt.Printf("spaces=%d\n", spaceCount)

	gap_spc := make([]int, gaps)

	if spaceCount > gaps {
		for x := 0; x < len(gap_spc); x++ {
			gap_spc[x] = spaceCount / gaps
		}
	}

	i := 0
	for x := spaceCount % gaps; x > 0; x-- {
		gap_spc[i]++
		i++
	}

	fmt.Printf("gap_spc[] = %v\n", gap_spc)

	ret := ""
	ret += sgroup[0]

	for x := 0; x < len(gap_spc); x++ {
		ret += strings.Repeat(" ", gap_spc[x])
		ret += sgroup[x+1]
	}

	if len(sgroup) > 1 {
		ret += sgroup[len(sgroup)-1]
	}

	fmt.Printf("ret = %v\n", ret)

}
