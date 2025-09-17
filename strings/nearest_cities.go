package strings

import (
	"errors"
	"slices"
	"strconv"
	"strings"
)

// Time Complexity: O(n log n)
// Space Complexity: O(n)

// cities (Accra, 121; Kumasi, 349; Takoradi, 718)
func NearestCities(lines string) ([]int32, error) {
	input := strings.Split(strings.TrimSpace(lines), ";")

	if len(input) == 0 {
		return nil, errors.New("input must not be empty")
	}

	var distances []int32
	for _, item := range input {
		parts := strings.Split(strings.TrimSpace(item), ",")

		if len(parts) != 2 {
			continue
		}

		distance, err := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 32)

		if err != nil {
			continue
		}

		distances = append(distances, int32(distance))

	}

	if len(distances) == 0 {
		return nil, errors.New("no valid distances found")
	}

	slices.Sort(distances)

	result := make([]int32, 0, len(distances))
	result = append(result, distances[0])

	for i := 1; i < len(distances); i++ {
		result = append(result, abs(distances[i], distances[i-1]))
	}

	return result, nil
}

func abs(a, b int32) int32 {
	if a < b {
		return b - a
	}
	return a - b
}
