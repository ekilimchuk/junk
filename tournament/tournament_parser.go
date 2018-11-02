package tournament

import (
	"errors"
	"strings"
)

func isValidStat(s []string) (bool, error) {
	if len(s) <= 1 {
		return false, nil
	}
	if len(s) != 3 {
		return false, errors.New("Parse error")
	}
	for _, v := range s[:1] {
		if v == "" {
			return false, errors.New("Parse error")
		}
	}
	if !(s[2] == "win" || s[2] == "draw" || s[2] == "loss") {
		return false, errors.New("Parse error")
	}
	return true, nil
}

func GameParcer(stat string) (Game, bool, error) {
	s := strings.Split(stat, ";")
	ok, err := isValidStat(s)
	if !ok {
		return Game{}, false, err
	}
	return Game{s[0], s[1], s[2]}, true, nil
}
