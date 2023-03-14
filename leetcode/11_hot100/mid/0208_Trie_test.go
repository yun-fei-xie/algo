package mid

/*
https://leetcode.cn/problems/implement-trie-prefix-tree/?favorite=2cktkvj
实现字典树,这里只要求存储英文单词26个字母
整个过程逻辑还是比较简单的，这里的代码直接抄了官方题解
*/

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func ConstructorTrie() Trie {
	return Trie{}
}

/*
体会一下插入过程，还是比较简单的。不需要进行递归，一条线下去
*/
func (this *Trie) Insert(word string) {
	node := this
	for _, ch := range word {
		ch -= 'a' // 拿到偏移量
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this.SearchPrefix(word)
	return node != nil && node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.SearchPrefix(prefix) != nil
}

/*
返回prefix最后一个字符代表的trie节点 ,如果不存在这种前缀，则返回nil
*/
func (this *Trie) SearchPrefix(prefix string) *Trie {
	node := this
	for _, ch := range prefix {
		ch -= 'a'
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
