package utils

type Trie struct {
	End   bool
	Child map[byte]*Trie
}

func Constructor() Trie {
	res := new(Trie)
	res.Child = make(map[byte]*Trie)
	return *res
}

func NewTire(str []string) Trie {
	constructor := Constructor()
	for _, s := range str {
		constructor.Insert(s)
	}
	return constructor
}

func (this *Trie) Insert(word string) {
	tmpTrie := this
	for i := 0; i < len(word); i++ {
		if _, ok := tmpTrie.Child[word[i]]; !ok {
			tmpNode := Constructor()
			tmpTrie.Child[word[i]] = &tmpNode
		}
		tmpTrie = tmpTrie.Child[word[i]]
	}
	tmpTrie.End = true
}

func (this *Trie) Search(word string) bool {
	tmpNode := this
	for i := 0; i < len(word); i++ {
		if val, ok := tmpNode.Child[word[i]]; ok {
			tmpNode = val
		} else {
			return false
		}
	}
	return tmpNode.End
}

func (this *Trie) StartsWith(prefix string) bool {
	tmpNode := this
	for i := 0; i < len(prefix); i++ {
		if val, ok := tmpNode.Child[prefix[i]]; ok {
			tmpNode = val
		} else {
			return false
		}
	}
	return true
}
