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

func punct(line string) string {

	re := regexp.MustCompile(`(\S+)\s*([.,!?;:])\s*`)
	return re.ReplaceAllString(line, "$1$2 ")
}

func Apost(line string) string {

	re := regexp.MustCompile(`'\s*([^']+?)\s*'`)
	return re.ReplaceAllString(line, "'$1'")
}

func ponc(s string) string {

	ponctuation := []string{".", ":", ",", ";", "?", "!"}
	for _, ponc := range ponctuation {
		s = strings.ReplaceAll(s, " "+ponc, ponc+" ")
		s = strings.ReplaceAll(s, " "+ponc, ponc)
	}
	return s
}

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
			nchaine = punct(nchaine)
			chaine := strings.Fields(nchaine)

			for i, mot := range chaine {

				if chaine[i] == "(hex)" && i != 0 {
					dc, err := strconv.ParseInt(chaine[i-1], 16, 64)
					if err == nil {
						dci := strconv.FormatInt(dc, 10)
						chaine[i-1] = dci
						for j := i; j >= 1; j-- {
							vide := ""
							chaine[j] = chaine[j-1]
							chaine[j-1] = vide
						}
					}

				}

				if chaine[i] == "(bin)" && i != 0 {
					dc, err := strconv.ParseInt(chaine[i-1], 2, 64)
					if err == nil {
						dci := strconv.FormatInt(dc, 10)
						chaine[i-1] = dci
						for j := i; j >= 1; j-- {
							vide := ""
							chaine[j] = chaine[j-1]
							chaine[j-1] = vide
						}
					}

				}
				if mot == "(up)" && i != 0 {
					chaine[i-1] = strings.ToUpper(chaine[i-1])
					for j := i; j >= 1; j-- {
						vide := ""
						chaine[j] = chaine[j-1]
						chaine[j-1] = vide

					}
				}
				if mot == "(up)" && i == 0 {
					chaine[i] = strings.Trim(chaine[i], chaine[i])
				}

				if chaine[i] == "(low)" && i != 0 {
					chaine[i-1] = strings.ToLower(chaine[i-1])
					for j := i; j >= 1; j-- {
						vide := ""
						chaine[j] = chaine[j-1]
						chaine[j-1] = vide
					}
				} else if chaine[i] == "(low)" && i == 0 {
					chaine[i] = strings.Trim(chaine[i], chaine[i])
				}

				if chaine[i] == "(cap)" && i != 0 {
					up := strings.ToLower(chaine[i-1])
					chaine[i-1] = strings.Title(up)
					for j := i; j >= 1; j-- {
						vide := ""
						chaine[j] = chaine[j-1]
						chaine[j-1] = vide

					}
				} else if chaine[i] == "(cap)" && i == 0 {
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
						for j := i + 1; j >= 2; j-- {
							chaine[j] = chaine[j-2]
							chaine[j-2] = ""
						}

					}

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
						for j := i + 1; j >= 2; j-- {
							chaine[j] = chaine[j-2]
							chaine[j-2] = ""
						}
						if i == 1 {
							chaine[i] = strings.Trim(chaine[i], chaine[i])
						}
					}

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
						for j := i + 1; j >= 2; j-- {
							chaine[j] = chaine[j-2]
							chaine[j-2] = ""
						}
						if i == 1 {
							chaine[i] = strings.Trim(chaine[i], chaine[i])
						}
					}

					nchaine = strings.Join(chaine, " ")
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

				// ******************** gerer la ponctuation*********************

			}
			nchaine = strings.Join(chaine, " ")
			nchaine = Apost(nchaine)
			nchaine = Aan(nchaine)

			nchaine = ponc(nchaine)

			ap := regexp.MustCompile(`\.\s+'`)
			nchaine = ap.ReplaceAllString(nchaine, `.'`)
			ao := regexp.MustCompile(`\!\s+'`)
			nchaine = ao.ReplaceAllString(nchaine, `!'`)
			ak := regexp.MustCompile(`\:\s+'`)
			nchaine = ak.ReplaceAllString(nchaine, `:'`)
			ai := regexp.MustCompile(`\;\s+'`)
			nchaine = ai.ReplaceAllString(nchaine, `;'`)
			am := regexp.MustCompile(`\,\s+'`)
			nchaine = am.ReplaceAllString(nchaine, `,'`)
			ad := regexp.MustCompile(`\?\s+'`)
			nchaine = ad.ReplaceAllString(nchaine, `?'`)
			SORTIE := os.WriteFile(os.Args[2], []byte((nchaine)), 0777)
			if SORTIE != nil {
				fmt.Println(Apost(nchaine))

			}
		}
	} else {
		log.Fatal("Verifie le nombre d'argument saisie ou le nom des fichiers")
	}

}
