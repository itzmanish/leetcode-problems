// 205. Isomorphic Strings
// Easy

// 3814

// 690

// Add to List

// Share
// Given two strings s and t, determine if they are isomorphic.

// Two strings s and t are isomorphic if the characters in s can be replaced to get t.

// All occurrences of a character must be replaced with another character while preserving
// the order of characters. No two characters may map to the same character, but a character
// may map to itself.

// Example 1:

// Input: s = "egg", t = "add"
// Output: true
// Example 2:

// Input: s = "foo", t = "bar"
// Output: false
// Example 3:

// Input: s = "paper", t = "title"
// Output: true

// Constraints:

// 1 <= s.length <= 5 * 104
// t.length == s.length
// s and t consist of any valid ascii character.

package leetcodepractice

func isIsomorphic(s string, t string) bool {
	m1 := make([]int, 256)
	m2 := make([]int, 256)

	for i := 0; i < len(s); i++ {
		if m1[int(s[i])] != m2[int(t[i])] {
			return false
		}

		i2 := i + 1
		m1[int(s[i])] = i2
		m2[int(t[i])] = i2
	}

	return true
}
