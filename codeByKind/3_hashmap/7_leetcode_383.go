package main

func canConstruct(ransomNote string, magazine string) bool {
	NoteMap := [26]int{}
	for _, v := range magazine {
		NoteMap[v-'a']++
	}
	for _, v := range ransomNote {
		NoteMap[v-'a']--
		if NoteMap[v-'a'] < 0 {
			return false
		}
	}
	return true
}
