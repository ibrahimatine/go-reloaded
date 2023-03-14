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
			chaine := strings.Split(string(data), " ")
			nchaine := string(data)

			for i, mot := range chaine {

				if strings.HasSuffix(mot, "(hex)") {

					dc, err := strconv.ParseInt(chaine[i-1], 16, 64)
					if err == nil {
						dci := strconv.FormatInt(dc, 10)
						chaine[i] = strings.TrimSuffix(mot, "(hex)")
						chaine[i-1] = dci //strings.ReplaceAll(nchaine, chaine[i-1], dci)

						nchaine = strings.Join(chaine, " ")
					}
				} else if strings.HasSuffix(mot, "(bin)") {
					dc, err := strconv.ParseInt(chaine[i-1], 2, 64)
					if err == nil {
						dci := strconv.FormatInt(dc, 10)
						chaine[i] = strings.TrimSuffix(mot, "(bin)")
						chaine[i-1] = dci

						nchaine = strings.Join(chaine, " ")

					}
				}

				if strings.HasSuffix(mot, "(up)") { //|| strings.HasSuffix(mot, "(up),")
					up := strings.ToUpper(chaine[i-1])
					chaine[i] = strings.TrimSuffix(mot, "(up)")
					chaine[i-1] = up
					nchaine = strings.Join(chaine, " ")

				}

				if strings.HasSuffix(mot, "(low)") {
					low := strings.ToLower(chaine[i-1])
					chaine[i] = strings.TrimSuffix(mot, "(low)")
					chaine[i-1] = low
					nchaine = strings.Join(chaine, " ")
				}

				if strings.HasSuffix(mot, "(cap)") {
					cap := strings.Title(chaine[i-1])
					chaine[i] = strings.TrimSuffix(mot, "(cap)")
					chaine[i-1] = cap
					nchaine = strings.Join(chaine, " ")
				}

				/*if mot == "(cap," {
					re := regexp.MustCompile("[0-9]+")
					indice := re.FindAllString(string(chaine[i+1]), -1)
					num := strings.Join(indice, "")
					conv, _ := strconv.Atoi(num)
					for j := 1; j <= conv; j++ {
						chaine[i-j] = strings.Title(chaine[i-j])
					}
					chaine[i] = ""
					chaine[i+1] = ""
					nchaine = strings.Join(chaine, " ")

				}*/
				if mot == "(cap," {
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
					//chaine = append(chaine[:i], chaine[i+2:]...)
					chaine[i] = ""
					chaine[i+1] = ""

					nchaine = strings.Join(chaine, " ")
				}

				if mot == "(up," {
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

				if mot == "(low," {
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

			}
			fmt.Println(esp(nchaine))
		}

	}

}
