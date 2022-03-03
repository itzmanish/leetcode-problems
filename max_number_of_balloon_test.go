package leetcodeplayground

import (
	"log"
	"testing"
)

// Not completed yet.

// Given a string text, you want to use the characters of text to form as many instances of the word "balloon" as possible.
// You can use each character in text at most once. Return the maximum number of instances that can be formed.

// Example 1:
// Input: text = "nlaebolko"
// Output: 1
// Example 2:
// Input: text = "loonbalxballpoon"
// Output: 2
// Example 3:
// Input: text = "leetcode"
// Output: 0
// example 4:
// Input: text = "balllllllllllloooooooooon"
// output: 1

// Constraints:
// 1 <= text.length <= 104
// text consists of lower case English letters only.

func TestMaxNumberOfBalloon(t *testing.T) {
	input := "krhizmmgmcrecekgyljqkldocicziihtgpqwbticmvuyznragqoyrukzopfmjhjjxemsxmrsxuqmnkrzhgvtgdgtykhcglurvppvcwhrhrjoislonvvglhdciilduvuiebmffaagxerjeewmtcwmhmtwlxtvlbocczlrppmpjbpnifqtlninyzjtmazxdbzwxthpvrfulvrspycqcghuopjirzoeuqhetnbrcdakilzmklxwudxxhwilasbjjhhfgghogqoofsufysmcqeilaivtmfziumjloewbkjvaahsaaggteppqyuoylgpbdwqubaalfwcqrjeycjbbpifjbpigjdnnswocusuprydgrtxuaojeriigwumlovafxnpibjopjfqzrwemoinmptxddgcszmfprdrichjeqcvikynzigleaajcysusqasqadjemgnyvmzmbcfrttrzonwafrnedglhpudovigwvpimttiketopkvqw"
	expected := 10
	actual := maxNumOfBalloon(input)
	log.Println(actual)
	if actual != expected {
		t.Fail()
	}
}

func maxNumOfBalloon(text string) int {
	// b=1,a=1,l=2,o=2,n=1
	// sum for one balloon = 1+1+2+2+1 = 7
	// total balloon word = 7n
	smap := map[rune]int{}
	sum := 0
	for _, letter := range text {
		if letter == 'b' || letter == 'a' || letter == 'l' || letter == 'o' || letter == 'n' {
			sum += 1
			smap[letter] += 1
		}
	}
	if len(smap) != 5 || smap['l'] < 2 {
		return 0
	}
	i := 1
	limit := smap['b']
	for ; i <= limit; i++ {
		valid := false
		for k, v := range smap {
			if (v == 'l' || v == 'o') && v >= 2 {
				smap[k] -= 2
				valid = true
			} else if v >= 1 && !(v == 'l' || v == 'o') {
				smap[k]--
				valid = true
			} else {
				valid = false
			}
		}
		if !valid {
			return i
		}
	}

	return i - 1
}
