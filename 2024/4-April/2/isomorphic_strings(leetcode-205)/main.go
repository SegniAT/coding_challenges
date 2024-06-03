package main

func main() {
}

func isIsomorphic(s string, t string) bool {
	sTt := make(map[byte]byte)
	tTs := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		sChar, tChar := s[i], t[i]
		res1, ok1 := sTt[sChar]
		res2, ok2 := tTs[tChar]

		if ok1 || ok2 {
			if res1 != tChar || res2 != sChar {
				return false
			}
		} else {
			sTt[sChar] = tChar
			tTs[tChar] = sChar
		}

	}

	return true
}
