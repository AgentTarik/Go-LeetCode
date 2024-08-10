package problems

import "sort"

func LongestCommonPrefix(strs []string) string {
	length := len(strs)

	if length == 1 {
		return strs[0]
	}

	sort.Strings(strs)
	for i := range strs[0] {
		if strs[0][i] != strs[length-1][i] {
			return strs[0][:i]
		}

	}
	return strs[0]
}
