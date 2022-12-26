package wave

import (
	"strings"
)

func Wave(s string) []string {
	w := []string{}

	for i, _ := range s {
		if string(s[i]) == " " {
			continue
		}

		upper := strings.ToUpper(string(s[i]))
		res := strings.Join([]string{s[0:i], upper, s[i+1:len(s)]}, "")
		w = append(w, res)
	}
	return w
}