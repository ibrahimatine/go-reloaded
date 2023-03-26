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

// fonction pour les apostrophes
/*func guil(s string) string {
	// Trouver tous les apostrophes dans la phrase
	mots := strings.Fields(s)
	cpt := 0
	for _, c := range mots {
		if c == "'" {
			cpt++
		}
	}

	// Vérifier que le nombre d'apostrophes est pair
	if cpt%2 == 0 {
		// Modifier la phrase avec les apostrophes correctement placés
		nchaine := s
		for i := 0; i < cpt; i += 2 {
			deb := mots[i] + 1
			fin := mots[i+1]
			nchaine = nchaine[:deb] + strings.TrimSpace(nchaine[deb:fin]) + nchaine[fin:]
		}
	}

	return nchaine
}*/

// ******************** gerer la ponctuation*********************
func coP(chaine string) string {
	ponctuation := []string{".", ":", ",", ";", "?", "!"}
	for _, ponc := range ponctuation {
		chaine = strings.ReplaceAll(chaine, " "+ponc, ponc+" ")
		chaine = strings.ReplaceAll(chaine, " "+ponc, ponc)
		//chaine = strings.ReplaceAll(chaine, ponc+" ", ponc)
	}
	return chaine
}

func Aan(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words)-1; i++ {
		if (words[i] == "a") && (isVowel(words[i+1][0])) {
			words[i] = "an"
		} else if (words[i] == "A") && (isVowel(words[i+1][0])) {
			words[i] = "An"
		}
	}
	return strings.Join(words, " ")
}
func isVowel(c byte) bool {
	vowels := "aeiouhAEIOUH"
	return strings.ContainsRune(vowels, rune(c))
}

func main() {
	//arguments := os.Args[1:]
	//if len(arguments) == 2 {
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
				// Extraire le nombre à partir de la chaîne suivante
				nbr := strings.TrimFunc(chaine[i+1], func(r rune) bool {
					return !unicode.IsDigit(r)
				})
				conv, err := strconv.Atoi(nbr)
				if err == nil {
					if conv <= i {
						// Capitaliser les mots précédents
						for j := i - conv; j < i; j++ {
							cap := []rune(chaine[j])
							cap[0] = unicode.ToUpper(cap[0])
							for i := range cap {
								if (i) != 0 {
									cap[j] = unicode.ToLower(cap[j])
								}
							}
							caps := string(cap)
							chaine[j] = caps
						}

					} else {
						for j := 1; j < i; j++ {
							cap := []rune(chaine[j])
							cap[0] = unicode.ToUpper(cap[0])
							for i := range cap {
								if (i) != 0 {
									cap[j] = unicode.ToLower(cap[j])
								}
							}
							caps := string(cap)
							chaine[j] = caps
						}

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
					if conv <= i {
						for j := i - conv; j < i; j++ {
							chaine[i-j] = strings.ToUpper(chaine[i-j])
						}
					} else {
						for j := 0; j <= i; j++ {
							chaine[i-j] = strings.ToUpper(chaine[i-j])
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
					if i <= conv {
						for j := i - conv; j <= i; j++ {
							chaine[i-j] = strings.ToLower(chaine[i-j])
						}
					} else {
						for j := 0; j <= i; j++ {
							chaine[i-j] = strings.ToLower(chaine[i-j])
						}
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
		fmt.Println(Aan(nchaine)) //(Aan(esp(coP(esp(nchaine)))))
	}
}

//}
