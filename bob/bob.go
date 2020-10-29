package bob

import "strings"

// Hey should have a comment documenting it.
func Hey(remark string) string {

	answer := ""
	switch {
	case remark[len(remark)-1:] == "?":
		answer = "Sure."
	case strings.ToUpper(remark) == remark:
		answer = "Whoa, chill out!"
	default:
		answer = "Whatever."
	}

	return answer

}
