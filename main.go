package main

import (
	"fmt"
	"strings"
)

func main() {

	words := []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog", "nifty!"}
	k := 16

	// 1) Build arrays of word groupings for each line
	groupings := [][]string{}
	currentGroup := []string{words[0]}
	letterCount := len(words[0]) // # of letters in the current group build

	for x := 1; x < len(words); x++ {
		if len(words[x])+1+letterCount > k {
			// we are starting a new group
			groupings = append(groupings, currentGroup)
			fmt.Println(currentGroup)
			currentGroup = []string{}
			letterCount = 0
		}
		letterCount += len(words[x]) + 1
		currentGroup = append(currentGroup, words[x])
	}

	if len(currentGroup) > 0 { // last group
		fmt.Println(currentGroup)
		groupings = append(groupings, currentGroup)
	}

	// have getGroupString() process each line and put them in results
	fmt.Println("-------------------")
	fmt.Println(groupings)

	results := []string{}
	for x := 0; x < len(groupings); x++ {
		results = append(results, getGroupString(groupings[x], k))
	}

	fmt.Println(strings.Repeat("*", k+4))
	for x := 0; x < len(results); x++ {
		fmt.Printf("* %v *\n", results[x])
	}
	fmt.Printf(strings.Repeat("*", k+4))

}

func getGroupString(sgroup []string, k int) string {
	fmt.Printf("Welcome to getGroupString()  sgroup=%v  k=%v\n", sgroup, k)

	if len(sgroup) == 1 {
		ret := sgroup[0] + strings.Repeat(" ", k-len(sgroup[0]))
		return ret
	}

	gaps := len(sgroup) - 1
	letterCount := 0
	for _, word := range sgroup {
		letterCount += len(word)
	}

	fmt.Printf("letters=%d\n", letterCount)

	spaceCount := k - letterCount
	fmt.Printf("spaces=%d\n", spaceCount)

	gapSpc := make([]int, gaps)

	if spaceCount > gaps {
		for x := 0; x < len(gapSpc); x++ {
			gapSpc[x] = spaceCount / gaps
		}
	}

	i := 0
	for x := spaceCount % gaps; x > 0; x-- {
		gapSpc[i]++
		i++
	}

	fmt.Printf("gapSpc[] = %v\n", gapSpc)

	ret := ""
	ret += sgroup[0]

	for x := 0; x < len(gapSpc); x++ {
		ret += strings.Repeat(" ", gapSpc[x])
		ret += sgroup[x+1]
	}

	fmt.Printf("ret = %v\n", ret)
	return ret

}
