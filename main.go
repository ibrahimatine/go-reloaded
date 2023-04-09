package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func Aan(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words)-1; i++ {
		if (words[i] == "a") && (isVowel(words[i+1][0])) {
			words[i] = "an"
		} else if (words[i] == "A") && (isVowel(words[i+1][0])) {
			words[i] = "An"
		} else if (words[i] == "'A") && (isVowel(words[i+1][0])) {
			words[i] = "'An"
		} else if (words[i] == "'a") && (isVowel(words[i+1][0])) {
			words[i] = "'an"
		}
	}
	return strings.Join(words, " ")
}
func isVowel(c byte) bool {
	vowels := "aeiouhAEIOUH"
	return strings.ContainsRune(vowels, rune(c))
}

func main() {
	arguments := os.Args[1:]
	if len(arguments) == 2 {
		data, err := os.ReadFile("sample.txt")
		if err == nil {
			re := regexp.MustCompile(`\s+`)
			nchaine := re.ReplaceAllString(string(data), " ")
			chaine := strings.Split(nchaine, " ")
			//nchaine = string(data)
			//numberStringCap := ""
			for i, mot := range chaine {

				if chaine[i] == "(hex)" && i != 0 {
					dc, err := strconv.ParseInt(chaine[i-1], 16, 64)
					if err == nil {
						dci := strconv.FormatInt(dc, 10)
						chaine[i-1] = dci
						nchaine = strings.Join(chaine, " ")
					}
				}

				if chaine[i] == "(bin)" && i != 0 {
					dc, err := strconv.ParseInt(chaine[i-1], 2, 64)
					if err == nil {
						dci := strconv.FormatInt(dc, 10)
						chaine[i-1] = dci
						nchaine = strings.Join(chaine, " ")
					}
				}
				if chaine[i] == "(up)" && i != 0 {
					up := strings.ToUpper(chaine[i-1])
					chaine[i-1] = up
					nchaine = strings.Join(chaine, " ")
				}
				if chaine[i] == "(up)" && i == 0 {
					chaine[i] = strings.Trim(chaine[i], chaine[i])
				}

				if chaine[i] == "(low)" && i != 0 {
					low := strings.ToLower(chaine[i-1])
					chaine[i-1] = low
					nchaine = strings.Join(chaine, " ")
				}
				if chaine[i] == "(low)" && i == 0 {
					chaine[i] = strings.Trim(chaine[i], chaine[i])
				}

				if chaine[i] == "(cap)" && i != 0 {
					up := strings.ToLower(chaine[i-1])
					chaine[i-1] = up
					chaine[i-1] = strings.Title(up)
					nchaine = strings.Join(chaine, " ")

				}
				if chaine[i] == "(cap)" && i == 0 {
					chaine[i] = strings.Trim(chaine[i], chaine[i])
				}

				if mot == "(cap," && i != 0 {
					nbr := strings.TrimFunc(chaine[i+1], func(r rune) bool {
						return !unicode.IsDigit(r)
					})
					conv, err := strconv.Atoi(nbr)
					if err == nil {
						if conv <= i {
							for j := i - conv; j <= i-1; j++ {
								chaine[j] = strings.ToLower(chaine[j])
								chaine[j] = strings.Title(chaine[j])
							}
						} else {
							for j := 0; j < i; j++ {
								chaine[j] = strings.ToLower(chaine[j])
								chaine[j] = strings.Title(chaine[j])
							}
						}
						if i == 1 {
							chaine[i] = strings.Trim(chaine[i], chaine[i])
						}

					}
					chaine[i] = ""
					chaine[i+1] = ""
					nchaine = strings.Join(chaine, " ")
				}

				if mot == "(up," && i != 0 {
					nbr := strings.TrimFunc(chaine[i+1], func(r rune) bool {
						return !unicode.IsDigit(r)
					})
					conv, err := strconv.Atoi(nbr)
					if err == nil {
						if conv <= i {
							for j := i - conv; j <= i-1; j++ {
								chaine[j] = strings.ToUpper(chaine[j])
							}
						} else {
							for j := 0; j < i; j++ {
								chaine[j] = strings.ToUpper(chaine[j])
							}
						}
					}
					chaine[i] = ""
					chaine[i+1] = ""
					nchaine = strings.Join(chaine, " ")
				}
				if mot == "(low," && i != 0 {
					nbr := strings.TrimFunc(chaine[i+1], func(r rune) bool {
						return !unicode.IsDigit(r)
					})
					conv, err := strconv.Atoi(nbr)
					if err == nil {
						if conv <= i {
							for j := i - conv; j <= i-1; j++ {
								chaine[j] = strings.ToLower(chaine[j])
							}
						} else {
							for j := 0; j < i; j++ {
								chaine[j] = strings.ToLower(chaine[j])
							}
						}
					}
					chaine[i] = ""
					chaine[i+1] = ""
					nchaine = strings.Join(chaine, " ")
				}

				it := regexp.MustCompile(`'\s*([^']+?)\s*'`)
				nchaine = it.ReplaceAllString(nchaine, "'$1'")

				// ******************** gerer la ponctuation*********************

				ponctuation := []string{".", ":", ",", ";", "?", "!"}
				for _, ponc := range ponctuation {
					nchaine = strings.ReplaceAll(nchaine, " "+ponc, ponc+" ")
					nchaine = strings.ReplaceAll(nchaine, " "+ponc, ponc)

				}

				// ******************** supprimer le surplus d'espace entre les mots ******************************
				re := regexp.MustCompile(`\s+`)
				nchaine = re.ReplaceAllString(nchaine, " ")
				nchaine = strings.Trim(nchaine, " ")

				//Suppression
				nchaine = strings.ReplaceAll(nchaine, "(bin)", "")
				nchaine = strings.ReplaceAll(nchaine, "(hex)", "")
				nchaine = strings.ReplaceAll(nchaine, "(up)", "")
				nchaine = strings.ReplaceAll(nchaine, "(low)", "")
				nchaine = strings.ReplaceAll(nchaine, "(cap)", "")

			}
			SORTIE := os.WriteFile(os.Args[2], []byte(Aan(nchaine)), 0777)
			if SORTIE != nil {
				fmt.Println(Aan(nchaine))

			}
		}
	} else {
		log.Fatal("Verifie le nombre d'argument saisie ou le nom des fichiers")
	}

}
