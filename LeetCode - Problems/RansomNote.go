/*
Given two strings ransomNote and magazine,
return true if ransomNote can be constructed by using the letters from magazine and false otherwise.
Each letter in magazine can only be used once in ransomNote.
*/

package RansomNote

func canConstruct(ransomNote string, magazine string) bool {
	counts := [26]int{}

	for i := 0; i < len(magazine); i++ {
		counts[magazine[i]-'a']++
	}

	for i := 0; i < len(ransomNote); i++ {
		counts[ransomNote[i]-'a']--
		if counts[ransomNote[i]-'a'] < 0 {
			return false
		}
	}
	return true
}
