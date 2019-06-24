package main
import (
	"fmt"
	"strings"
)

type Node struct {
	Children map[rune]*Node
	IsWord   bool
}

type Trie struct {
	Root Node
}

func NewNode() *Node {
	return &Node{Children: make(map[rune]*Node)}
}

func NewTrie() *Trie {
	root := NewNode()
	return &Trie{Root: *root}
}

func (t *Trie) PushWord(word string) {
	root := &t.Root
	for _, char := range word {
		child, found := root.Children[char]
		if found {
			root = child
		} else {
			child := NewNode()
			root.Children[char] = child
			root = child
		}
	}
	root.IsWord = true
}

func (t *Trie) PushText(text string) {
	words := strings.Fields(text)
	for _, word := range words {
		t.PushWord(word)
	}
}

func (t *Trie) Complete(prefix string) []string{
	root := &t.Root
	var res = []string{}
	for _, char := range prefix {
		child, found := root.Children[char]
		if found {
			root = child
		} else {
			fmt.Printf("cannot find %v", char)
			return nil
		}
	}
	for char, node := range root.Children {
		if node.IsWord {
			res = append(res, prefix + string(char))
		}
		res = append(res, t.Complete(prefix + string(char))...)
	}
	return res
}

func repr(root Node, level int) {
	for char, node := range root.Children {
		for index := 0; index < level; index++ {
			fmt.Printf(" ")
		}
		fmt.Println(string(char))
		repr(*node, level+1)
	}
}

func (t *Trie) Repr() {
	repr(t.Root, 0)
}
