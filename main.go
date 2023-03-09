package main

import (
	"fmt"
	"strconv"
	"strings"
)

func conversion(s string) string {

	chaine := strings.Split(s, " ")
	for i, mot := range chaine {
		if strings.HasSuffix(mot, "(hex)") {
			dc, err := strconv.ParseInt(chaine[i-1], 16, 64)
			if err != nil {
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

}

func main() {

	s := "1E (hex) files were added. It has been 10 (bin) years."
	fmt.Println(conversion(s))
}
