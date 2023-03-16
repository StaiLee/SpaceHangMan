package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
)

type Hangman struct {
	TryLeft        int
	NbrTry         int
	WordHidden     string
	Word           string
	Guessedletter  []string
	Guessedletter1 []string
	Answer         string
	GameState      int
}

var hangman Hangman

func HangmanStart() {
	hangman = Hangman{
		TryLeft:        12,
		NbrTry:         0,
		Guessedletter:  []string{},
		Guessedletter1: []string{},
		GameState:      0,
	}
}

func StartGame(filename string) {
	HangmanStart()
	ReadWord(filename)
	PickWord(filename)

}
func PickWord(filename string) {
	tw := ReadWord(filename)
	hangman.Word = tw[rand.Intn(len(tw)-1)]
	hangman.Word = hangman.Word[:len(hangman.Word)-1]
	hangman.WordHidden = wordToUnderScore()
}

func ReadWord(filename string) []string {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	word1 := ""
	var todowordos []string
	for _, char := range string(file) {
		if char == '\n' {
			todowordos = append(todowordos, word1)
			word1 = ""
		} else {
			word1 += string(char)
		}
	}
	return todowordos
}

func wordToUnderScore() string {
	sampleRegexp := regexp.MustCompile("[a-z,A-Z]")

	input := hangman.Word

	result := sampleRegexp.ReplaceAllString(input, "_")
	return (string(result))
}

func FindAndReplace(letterToReplace string) {
	isALetter, err := regexp.MatchString("^[a-zA-Z]", letterToReplace)
	if !isALetter || err != nil {
		return
	}

	hangman.NbrTry++
	if len(letterToReplace) != 0 {
		for _, guess := range hangman.Guessedletter {
			if letterToReplace == guess {
				return
			}
		}
		hangman.Guessedletter = append(hangman.Guessedletter, letterToReplace)
	}
	if len(letterToReplace) > 1 {
		if letterToReplace == hangman.Word {
			print(2)
			hangman.WordHidden = hangman.Word
		} else {
			hangman.TryLeft -= 2
		}
		if hangman.TryLeft < 0 {
			hangman.TryLeft = 0
		}
		return
	}

	isFound := strings.Index(hangman.Word, letterToReplace)
	if isFound == -1 {
		if hangman.TryLeft >= 1 {
			hangman.TryLeft--
			//deathCountStage(hangman.DeathCount)
			fmt.Println("raté")
			fmt.Println("Il vous reste", hangman.TryLeft, "essais")
			// mettre à jour le score
		}

	} else {
		str3 := []rune(hangman.WordHidden)
		for i, lettre := range hangman.Word {
			if string(lettre) == letterToReplace {
				str3[i] = lettre
				hangman.WordHidden = string(str3)
			}
		}
	}
}

func testEndGame() {
	if hangman.WordHidden == hangman.Word {
		hangman.GameState = 1
	} else if hangman.TryLeft <= 0 {
		hangman.GameState = 2
	}
}

func testWord() bool {
	hangman.NbrTry++
	// créer une var scanner qui va lire ce que l'utilisateur va écrire
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // l'utilisateur input dans la console
	// lis ce que l'utilisateur a écrit
	lettreoumot := scanner.Text()
	lettreoumot = strings.ToLower(lettreoumot)
	// peret à l'utilisateur de savoir qu'il ne doit mettre que des lettres contenues dans l'alphabet latin
	isALetter, err := regexp.MatchString("^[a-zA-Z]", lettreoumot)
	if Contains1(hangman.Guessedletter, lettreoumot) {
	} else {
		if err != nil {
			return testWord()
		}
		if !isALetter {
			return testWord()
		}
		if len(lettreoumot) == 1 {

			FindAndReplace(lettreoumot)
		} else if lettreoumot == hangman.Word {
			return true
		} else if (len(lettreoumot) == len(hangman.Word)) && hangman.WordHidden == hangman.Word {
			return true
		} else {
			hangman.TryLeft -= 2
			//deathCountStage(hangman.DeathCount)
		}

	}
	return false
}

func deathCountStage() int {

	file, err := os.Open("../hangman.txt")
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	fileScanner := bufio.NewScanner(file)
	index := 0
	var death int
	var start int
	var end int
	if death == 9 {
		start = 0
		end = 7
		return 9
	}

	if death == 8 {
		start = 8
		end = 15
		return 8
	}
	if death == 7 {
		start = 16
		end = 23
		return 7
	}
	if death == 6 {
		start = 24
		end = 31
		return 6
	}
	if death == 5 {
		start = 32
		end = 39
		return 5
	}
	if death == 4 {
		start = 40
		end = 47
		return 4
	}
	if death == 3 {
		start = 48
		end = 55
		return 3
	}
	if death == 2 {
		start = 56
		end = 63
		return 2
	}
	if death == 1 {
		start = 64
		end = 71
		return 1
	}
	if death == 0 {
		start = 72
		end = 79
		return 0
	}
	for fileScanner.Scan() {
		if index >= start && index <= end {
			println(fileScanner.Text())
		}
		index++
	}
	return index
}

func GameState() {
	if testWord() || !Contains(hangman.WordHidden, '_') {
		hangman.GameState = 1
	}
	if deathCountStage() == 0 {
		hangman.GameState = 2
	}
}

func Retry() {
	hangman.NbrTry = 0
	hangman.TryLeft = 12
	hangman.Guessedletter = hangman.Guessedletter1

}

func Contains(s string, char rune) bool { // Si une string est contenue dans un tableau
	for _, a := range s {
		if a == char {
			return true
		}
	}
	return false
}

func Contains1(s []string, char string) bool { // Si une string est contenue dans un tableau
	for _, a := range s {
		if a == char {
			return true
		}
	}
	return false
}

func IsLetter2(r rune) bool {
	lettermaj := ('A' <= r) && (r <= 'Z')
	lettermin := ('a' <= r) && (r <= 'z')
	if lettermaj || lettermin {
		return true
	}
	return false
}

func IsMin(r rune) bool {
	lettermin := ('a' <= r) && (r <= 'z')
	return lettermin
}

func IsMaj(r rune) bool {
	lettermaj := ('A' <= r) && (r <= 'Z')
	return lettermaj
}

func Capitalize(s string) string {
	str2 := []rune(s)
	for index := range str2 {
		if index == 0 {
			if IsMin(str2[index]) {
				str2[index] -= 32
			}
		} else if !IsLetter2(str2[index-1]) && IsMin(str2[index]) {
			str2[index] -= 32
		} else if IsLetter2(str2[index-1]) && IsMaj(str2[index]) {
			str2[index] += 32
		}
	}
	return string(str2)
}
