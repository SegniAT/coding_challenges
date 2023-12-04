package main

func main() {

}

func canConstruct(ransomeNote string, magazine string) bool {
	ransomeNoteCharCount := make(map[rune]int)
	magazineCharCount := make(map[rune]int)

	for i := 0; i < len(ransomeNote); i++ {
		ransomeNoteChar := ransomeNote[i]
		ransomeNoteCharCount[rune(ransomeNoteChar)]++
	}

	for i := 0; i < len(magazine); i++ {
		magazineChar := magazine[i]
		magazineCharCount[rune(magazineChar)]++
	}

	for char, count := range ransomeNoteCharCount {
		if magazineCharCount[char] < count {
			return false
		}
	}

	return true
}
