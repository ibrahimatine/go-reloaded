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
				} else if chaine[i] == "(bin)" && i != 0 {
					dc, err := strconv.ParseInt(chaine[i-1], 2, 64)
					if err == nil {
						dci := strconv.FormatInt(dc, 10)
						chaine[i-1] = dci
						nchaine = strings.Join(chaine, " ")
					}
				}
				if chaine[i] == "(up)" {
					up := strings.ToUpper(chaine[i-1])
					chaine[i-1] = up
					nchaine = strings.Join(chaine, " ")
				}
				if chaine[i] == "(low)" && i != 0 {
					low := strings.ToLower(chaine[i-1])
					chaine[i-1] = low
					nchaine = strings.Join(chaine, " ")
				}
				if chaine[i] == "(cap)" && i != 0 {

					cap := []rune(chaine[i-1])
					cap[0] = unicode.ToUpper(cap[0])
					for i := range cap {
						if i != 0 {
							cap[i] = unicode.ToLower(cap[i])
						}
					}
					caps := string(cap)
					chaine[i-1] = caps
					nchaine = strings.Join(chaine, " ")
				}

				if mot == "(cap," && i != 0 {
					nbr := strings.TrimFunc(chaine[i+1], func(r rune) bool {
						return !unicode.IsDigit(r)
					})
					conv, err := strconv.Atoi(nbr)
					if err == nil {
						if conv <= i {
							for j := i - conv; j <= i-1; j++ {
								chaine[j] = strings.Title(chaine[j])
							}
						} else {
							for j := 0; j < i; j++ {
								chaine[j] = strings.Title(chaine[j])
							}
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
				tab := strings.Split(nchaine, " ")
				apos := 0
				c := 0

				for j := range tab {
					if tab[j] == "'" && apos == 0 {
						c = j
						apos = 1
					}
					if tab[j] == "'" && apos == 1 && j > c {
						tab[c+1] = strings.Trim(tab[c]+tab[c+1], " ")
						tab[c] = strings.Trim(tab[c], "'")
						tab[j-1] = strings.Trim(tab[j-1]+tab[j], " ")
						tab[j] = strings.Trim(tab[j], "'")
						nchaine = strings.Join(tab, " ")
						nchaine = Aan(nchaine)
						apos = 0
					}
				}

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

				//Suppression
				nchaine = strings.ReplaceAll(nchaine, "(bin)", "")
				nchaine = strings.ReplaceAll(nchaine, "(hex)", "")
				nchaine = strings.ReplaceAll(nchaine, "(up)", "")
				nchaine = strings.ReplaceAll(nchaine, "(low)", "")
				nchaine = strings.ReplaceAll(nchaine, "(cap)", "")

				nchaine = strings.ReplaceAll(nchaine, " "+"'", "'")

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
