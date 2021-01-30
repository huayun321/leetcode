package p3_longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) (l int) {
	m := make(map[uint8]bool)
	for i := range s {
		for j := i; j < len(s); j++ {
			if _, ok := m[s[j]]; ok {
				if len(m) > l {
					l = len(m)
				}
				m = make(map[uint8]bool)
				break
			}
			m[s[j]] = true
		}
	}
	if l > len(m) {
		return l
	} else {
		return len(m)
	}
}
