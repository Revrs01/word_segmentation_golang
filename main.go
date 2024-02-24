package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"unicode/utf8"
	"wordSegmentation/tools"
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

func wordSegmentation(subSentence string, dictionary *tools.Set, root *Tree) {
	if len(subSentence) < 2 {
		return
	}
	subSentenceSlice := []rune(subSentence)
	var foundWord = false
	for gramLength := len(subSentenceSlice); gramLength > 0; gramLength-- {
		if foundWord == true || gramLength < 1 {
			//root.Value = subSentence
			break
		}

		for startIndex := 0; startIndex < len(subSentenceSlice)-gramLength+1; startIndex++ {
			currentSubSentence := string(subSentenceSlice[startIndex : startIndex+gramLength])
			//fmt.Println("currentWord: ", currentSubSentence)

			if dictionary.Contains(currentSubSentence) {
				//fmt.Println("splitIndex: ", startIndex, startIndex+gramLength, currentSubSentence)
				root.Value = currentSubSentence
				root.Left = &Tree{}
				root.Right = &Tree{}
				foundWord = true
				wordSegmentation(string(subSentenceSlice[:startIndex]), dictionary, root.Left)
				wordSegmentation(string(subSentenceSlice[startIndex+gramLength:]), dictionary, root.Right)
				break
			}
		}
	}
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

func inOrderTraverseTree(root *Tree) {
	if root == nil {
		return
	}

	inOrderTraverseTree(root.Left)
	if root.Value != "" {
		fmt.Print(root.Value, " | ")
	}
	inOrderTraverseTree(root.Right)
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

	if prefix != len(runeSlice) {
		result = append(result, string(runeSlice[prefix:]))
	}

	return result
}

func main() {
	// 早上好，現在我有Ice Cream，Mi-2S還有個老頭喜歡BB。
	// 大家好，我是盧文祥老師，我來自Mi-2S實驗室，我是一個超亮亮的老師，最喜歡帶著同學做沒有用的Project，然後不斷地改A改，超爽der

	dictionary := readDictionary()
	sentence := scanInput()
	regex, _ := regexp.Compile("[A-Za-z0-9 -]+")

	splitEnglishAndOtherLanguage := convertToStringArray(regex, sentence)
	for _, e := range splitEnglishAndOtherLanguage {
		if regex.MatchString(e) {
			fmt.Print(e, " | ")
		} else {
			// perform word segmentation
			segmentationTree := &Tree{}
			wordSegmentation(e, dictionary, segmentationTree)
			inOrderTraverseTree(segmentationTree)
		}
		//fmt.Println(i, len(e))
	}
}
