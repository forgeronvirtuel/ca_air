package air

import (
	"strings"

	"github.com/forgeronvirtuel/ca_eau"
)

// Split splits a string into a list of substring for each match
// with sep.
func Split(txt, sep string) []string {
	r := rune(sep[0])
	pos := ca_eau.FindAllPosition(txt, r)

	last := 0
	var res []string

	for _, p := range pos {
		if !ca_eau.IsSubstringAt(txt, sep, p) {
			continue
		}

		res = append(res, txt[last:p])
		last = p + len(sep)
	}

	if last < len(txt) {
		res = append(res, txt[last:])
	}

	return res
}

// Concat concatenate a list of string, separated by the sep string
func Concat(txt []string, sep string) string {
	var builder strings.Builder
	for idx, s := range txt {
		builder.WriteString(s)
		if idx < len(txt)-1 {
			builder.WriteString(sep)
		}
	}
	return builder.String()
}
