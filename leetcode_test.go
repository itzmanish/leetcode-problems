package leetcodeplayground

import (
	"fmt"
	"log"
	"testing"
)

func TestCloneGraph(t *testing.T) {
	node1 := Node{
		Val:       1,
		Neighbors: []*Node{},
	}
	node2 := Node{
		Val:       2,
		Neighbors: []*Node{},
	}
	node3 := Node{
		Val:       3,
		Neighbors: []*Node{},
	}
	node4 := Node{
		Val:       4,
		Neighbors: []*Node{},
	}
	node1.Neighbors = append(node1.Neighbors, &node2, &node4)
	node2.Neighbors = append(node2.Neighbors, &node1, &node3)
	node3.Neighbors = append(node3.Neighbors, &node2, &node4)
	node4.Neighbors = append(node4.Neighbors, &node1, &node3)

	cloned1 := CloneGraph(&node1)
	log.Printf("%v", node1)
	log.Printf("%v", cloned1)

}

func TestReverseLetter(t *testing.T) {
	// ex 1:
	// Input: s = "ab-cd"
	// Output: "dc-ba"
	// eg-2:
	// Input: s = "a-bC-dEf-ghIj"
	// Output: "j-Ih-gfE-dCba"
	// eg-3:
	// Input: s = "Test1ng-Leet=code-Q!"
	// Output: "Qedo1ct-eeLg=ntse-T!"
	for i := 33; i <= 122; i++ {
		if (i >= 65 && i <= 90) || (i >= 97 && i <= 122) {
			continue
		}
		fmt.Printf("%c", i)
	}
	testcase := "a-bC-dEf-ghIj"
	expected := "j-Ih-gfE-dCba"
	actual := ReverseOnlyLetters(testcase)
	log.Println(actual)
	if actual != expected {
		t.Fail()
	}
}
