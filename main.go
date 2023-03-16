package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func esp(s string) string {

	re := regexp.MustCompile(`\s+`)
	s = re.ReplaceAllString(s, " ")

	return s
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
					cap := strings.Title(chaine[i-1])
					chaine[i-1] = cap
					nchaine = strings.Join(chaine, " ")
				}

				if mot == "(cap," && i != 0 {
					// Extraire le nombre à partir de la chaîne suivante
					nbr := strings.TrimFunc(chaine[i+1], func(r rune) bool {
						return !unicode.IsDigit(r)
					})
					conv, err := strconv.Atoi(nbr)
					if err == nil {

						// Capitaliser les mots précédents
						for j := 1; j <= conv; j++ {
							chaine[i-j] = strings.Title(chaine[i-j])
						}
					}
					// Supprimer les mots actuel et suivant

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
						for j := 1; j <= conv; j++ {
							chaine[i-j] = strings.ToUpper(chaine[i-j])
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
						for j := 1; j <= conv; j++ {
							chaine[i-j] = strings.ToLower(chaine[i-j])
						}
					}
					chaine[i] = ""
					chaine[i+1] = ""

					nchaine = strings.Join(chaine, " ")
				}

				//Suppression
				nchaine = strings.ReplaceAll(nchaine, "(bin)", "")
				nchaine = strings.ReplaceAll(nchaine, "(hex)", "")
				nchaine = strings.ReplaceAll(nchaine, "(up)", "")
				nchaine = strings.ReplaceAll(nchaine, "(low)", "")
				nchaine = strings.ReplaceAll(nchaine, "(cap)", "")

			}
			fmt.Println(esp(nchaine))
		}

	}

}
