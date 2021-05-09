package main

func longestPalindrome(s string) string {
	var res string
	runes := []rune(s)
	if len(runes) <= 1 {
		return s
	}
	for i := 0; i < len(runes); i++ {
		// 以字母为中心
		rl, rr := -1, -2
		for l, r := i-1, i+1; l >= 0 && r < len(runes); l, r = l-1, r+1 {
			if runes[l] == runes[r] {
				rl, rr = l, r
				continue
			} else {
				break
			}
		}
		if rr-rl+1 > len(res) {
			res = string(runes[rl : rr+1])
		}

		// 以间隙为中心
		rl, rr = -1, -2
		for l, r := i, i+1; l >= 0 && r < len(runes); l, r = l-1, r+1 {
			if runes[l] == runes[r] {
				rl, rr = l, r
				continue
			} else {
				break
			}
		}
		if rr-rl+1 > len(res) {
			res = string(runes[rl : rr+1])
		}
	}
	if res == "" {
		return string(runes[:1])
	}
	return res
}
