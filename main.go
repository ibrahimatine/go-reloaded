package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*func conversion(s string) string {

	chaine := strings.Split(s, " ")
	for i, mot := range chaine {
		if strings.HasSuffix(mot, "(hex)") {
			dc, err := strconv.ParseInt(chaine[i-1], 16, 64)
			if err == nil {
				return s
			}
			chaine[i-1] = strconv.FormatInt(dc, 10) //pour convertir les int en alpha
			chaine[i] = strings.TrimSuffix(mot, "(hex)")
		} else if strings.HasSuffix(mot, "(bin)") {
			dc, err := strconv.ParseInt(chaine[i-1], 2, 64)
			if err != nil {
				return s
			}
			chaine[i-1] = strconv.FormatInt(dc, 10)
			chaine[i] = strings.TrimSuffix(mot, "(bin)")

		}
	}

	return strings.Join(chaine, " ")

}*/

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

				if strings.HasSuffix(mot, "(up)") || strings.HasSuffix(mot, "(up),") {
					up := strings.ToUpper(chaine[i-1])
					chaine[i] = strings.TrimSuffix(mot, "(up)")
					if strings.HasSuffix(mot, "(up),") {
						chaine[i] = strings.TrimSuffix(mot, "(up)")
					}

					chaine[i-1] = up
					nchaine = strings.Join(chaine, " ")

				}

			}
			fmt.Println(nchaine)
		}

	}

}
