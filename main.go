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
	mandarinDictionary, errZh := os.Open("dictionaries/Mandarin.txt")
	taiwaneseDictionary, errTw := os.Open("dictionaries/Taiwanese.txt")
	hakkaDictionary, errHk := os.Open("dictionaries/Hakka.txt")
	set := tools.NewSet()
	if errZh == nil && errTw == nil && errHk == nil {
		scannerZh := bufio.NewScanner(mandarinDictionary)
		scannerTw := bufio.NewScanner(taiwaneseDictionary)
		scannerHk := bufio.NewScanner(hakkaDictionary)

		for scannerZh.Scan() {
			set.Add(scannerZh.Text())
		}
		for scannerTw.Scan() {
			set.Add(scannerTw.Text())
		}
		for scannerHk.Scan() {
			set.Add(scannerHk.Text())
		}
		return set
	} else {
		fmt.Println("error")
	}
	return nil
}

func wordSegmentation(subSentence string, dictionary *tools.Set, root *Tree) {
	// build a tree
	// if word found in dictionary, then make the current word as root
	// and left subtree is string[:foundWordStartIndex]
	// right subtree is string[foundWordEndIndex:]
	// keep using n-gram for every node of subtrees

	// and last, output the tree by using Inorder fashion
	// you will get the segmented result
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
}

func scanInput() string {
	fmt.Print("Sentence: ")
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
	dictionary := readDictionary()
	sentence := scanInput()
	fmt.Println("Result: ")
	regex, _ := regexp.Compile("[A-Za-z0-9 -]+")
	// break sentence into several pieces
	// sentence -> English|(Zh, Tw, Hk) -> (Zh, Tw, Hk) fragments
	splitEnglishAndOtherLanguage := convertToStringArray(regex, sentence)
	for _, chineseSentence := range splitEnglishAndOtherLanguage {
		if regex.MatchString(chineseSentence) {
			fmt.Print(chineseSentence, " | ")
		} else {
			regexChinesePunctuation, _ := regexp.Compile("[，。、]")
			splitByChinesePunctuation := convertToStringArray(regexChinesePunctuation, chineseSentence)
			for _, chineseFragment := range splitByChinesePunctuation {
				if regexChinesePunctuation.MatchString(chineseFragment) {
					fmt.Print(chineseFragment, " | ")
				} else {
					// perform word segmentation
					segmentationTree := &Tree{}
					wordSegmentation(chineseFragment, dictionary, segmentationTree)
					inOrderTraverseTree(segmentationTree)
				}
			}

		}
	}
}
