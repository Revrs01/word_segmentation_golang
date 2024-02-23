package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"testGo/tools"
	"unicode/utf8"
)

func readFile() {
	file, err := os.Open("dict_no_space.txt")
	var counter = 0
	if err == nil {
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			counter += 1
			fmt.Println(scanner.Text())
		}
		fmt.Println(counter)
	} else {
		fmt.Println("error")
	}
}

func wordSegmentation(subSentence string) {
	// use n-gram, build a tree
	// if word found in dictionary, then make the current word as root
	// and left subtree is string[:foundWordStartIndex]
	// right subtree is string[foundWordEndIndex:]
	// keep using n-gram for every node of subtrees

	// and last, output the tree by using Inorder fashion
	// you will get the segmented result
}

func main() {
	// 早上好，現在我有Ice Cream，Mi2S還有個老頭喜歡BB
	scanner := bufio.NewScanner(os.Stdin)
	regex, _ := regexp.Compile("[A-Za-z0-9 ]+")
	scanner.Scan()
	sentence := scanner.Text()
	fmt.Println()
	runeSlice := []rune(sentence)
	matchesIndices := regex.FindAllStringIndex(sentence, -1)

	var runeEnglishIdx [][2]int
	for _, indexPair := range matchesIndices {
		ByteStartIndex, ByteEndIndex := indexPair[0], indexPair[1]

		runeStartIndex := utf8.RuneCountInString(sentence[:ByteStartIndex])
		runeEndIndex := utf8.RuneCountInString(sentence[:ByteEndIndex])
		runeEnglishIdx = append(runeEnglishIdx, [2]int{runeStartIndex, runeEndIndex})

	}
	var result []string
	var prefix = 0
	for _, element := range runeEnglishIdx {
		result = append(result, string(runeSlice[prefix:element[0]]))
		result = append(result, string(runeSlice[element[0]:element[1]]))
		prefix = element[1]
	}

	for i, e := range result {
		if regex.MatchString(e) {
			fmt.Println(i, "This is english or number or space")
		} else {
			// perform word segmentation
		}
		fmt.Println(i, e)
	}

	set := tools.NewSet()
	set.Add("早上好")
	set.Add("老頭")
	fmt.Println(set.Contains("老頭"))

}
