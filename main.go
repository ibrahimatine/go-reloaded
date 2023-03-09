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
			hx := strings.TrimSuffix(chaine[i-1], "(hex)")
			dc, err := strconv.ParseInt(hx, 16, 64)
			if err != nil {
				return s
			}
			chaine[i-1] = strconv.FormatInt(dc, 10) // Pour remplacer le int en caractere

		}
	}

	return strings.Join(chaine, " ")

}

func main() {

	s := "1E (hex) files were added. It has been F (hex) years."
	fmt.Println(conversion(s))
}
