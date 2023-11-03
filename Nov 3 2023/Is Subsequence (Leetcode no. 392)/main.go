package main

func main() {

}

func isSubsequence(s, t string) bool {
	ptr1, ptr2 := 0, 0
	lenS, lenT := len(s), len(t)

	for ptr1 < lenS && ptr2 < lenT {
		if s[ptr1] == t[ptr2] {
			ptr1++
			if ptr1 >= lenS {
				return true
			}
		}
		ptr2++

		if ptr1 < lenS && ptr2 >= lenT && s[ptr1] != t[ptr2-1] {
			return false
		}
	}

	if ptr1 < lenS {
		return false
	}

	return true
}
