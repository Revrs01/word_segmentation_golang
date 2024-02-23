package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"testGo/tools"
	"unicode/utf8"
)

type Tree struct {
	Value       string
	Left, Right *Tree
}

func readDictionary() *tools.Set {
	file, err := os.Open("dict_no_space.txt")
	set := tools.NewSet()
	if err == nil {
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			set.Add(scanner.Text())
		}
		return set
	} else {
		fmt.Println("error")
	}
	return nil
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

func scanInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func convertToStringArray(regex *regexp.Regexp, sentence string) []string {
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

	return result
}

func main() {
	// 早上好，現在我有Ice Cream，Mi2S還有個老頭喜歡BB

	dictionary := readDictionary()
	fmt.Print(dictionary)
	sentence := scanInput()
	regex, _ := regexp.Compile("[A-Za-z0-9 ]+")

	splitEnglishAndOtherLanguage := convertToStringArray(regex, sentence)
	for i, e := range splitEnglishAndOtherLanguage {
		if regex.MatchString(e) {
			fmt.Println(i, "This is english or number or space")
		} else {
			// perform word segmentation
			// segmentationResult := wordSegmentation(splitEnglishAndOtherLanguage)
		}
		fmt.Println(i, e)
	}

	set := tools.NewSet()
	set.Add("早上好")
	set.Add("老頭")
	fmt.Println(set.Contains("老頭"))

}
