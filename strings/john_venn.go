package strings

import (
	"errors"
	"strings"
)

// Time complexity O(m + n + k)
// Space complexity O(n + m)
func JohnVenn(lines string) (string, error) {
	input := strings.Split(strings.TrimSpace(lines), "\n")
	if len(input) != 2 {
		return "", errors.New("input must contain exactly two lines")
	}

	order := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz" // O(k)

	set1 := strings.Fields(input[0])
	set2 := strings.Fields(input[1])

	set1Map := make(map[string]bool) // O(n)
	set2Map := make(map[string]bool) // O(m)

	for _, elem := range set1 {
		set1Map[elem] = true
	}

	for _, elem := range set2 {
		set2Map[elem] = true
	}

	var result []string
	for _, char := range order {
		charStr := string(char)

		if set1Map[charStr] && set2Map[charStr] {
			result = append(result, charStr)
		}
	}

	return strings.Join(result, " "), nil
}
