package utils

import "strings"

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func StartsWith(target string, command string) bool {
	return (strings.Fields(target)[0] == command)

}
