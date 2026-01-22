package errs

import (
	"fmt"
	"maps"
	"slices"
	"sort"
	"strings"
)

func stringifyMap[V any, T map[string][]V](m T) string {
	if len(m) == 0 {
		return ""
	}

	// Get sorted keys
	keys := slices.Sorted(maps.Keys(m))

	// res is the outer result slice
	// based on the sorted keys
	var res []string

	// Iterate over sorted keys
	for _, k := range keys {

		// lines is the inner result slice
		var lines []string
		for _, v := range m[k] {
			// Append each error message with a tab and hyphen
			lines = append(lines, fmt.Sprintf("\t- %v", v))
		}

		// Sort the inner lines
		sort.Strings(lines)

		// Combine key and its lines
		s := fmt.Sprintf("%s:\n%s", k, strings.Join(lines, "\n"))

		// Append to the outer result
		res = append(res, s)
	}

	// Join all key sections with newlines
	return strings.Join(res, "\n")
}
