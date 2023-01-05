package hard

type TrieNode struct {
	children map[byte]*TrieNode
	word     string
}

func (trie *TrieNode) Insert(word string) {
	for i := range word {
		c := word[i]
		if trie.children[c] == nil {
			trie.children[c] = &TrieNode{children: map[byte]*TrieNode{}}
		}
		trie = trie.children[c]
	}
	trie.word = word
}

func (trie *TrieNode) IsEmpty() bool {
	return len(trie.children) == 0
}

func (trie *TrieNode) Remove(word string) bool {
	if len(word) == 0 {
		trie.word = ""
		return trie.IsEmpty()
	}
	next := trie.children[word[0]]
	if next.Remove(word[1:]) {
		delete(trie.children, word[0])
		return trie.IsEmpty()
	}
	return false
}

func dfs(res *[]string, root, curr *TrieNode, board [][]byte, i, j int) {
	c := board[i][j]
	if c == 0 || curr.children[c] == nil {
		return
	}
	curr = curr.children[c]
	if len(curr.word) > 0 {
		*res = append(*res, curr.word)
		root.Remove(curr.word)
	}
	board[i][j] = 0
	if i+1 < len(board) {
		dfs(res, root, curr, board, i+1, j)
	}
	if i-1 >= 0 {
		dfs(res, root, curr, board, i-1, j)
	}
	if j+1 < len(board[0]) {
		dfs(res, root, curr, board, i, j+1)
	}
	if j-1 >= 0 {
		dfs(res, root, curr, board, i, j-1)
	}
	board[i][j] = byte(c)
}

func findWords(board [][]byte, words []string) []string {
	trie := &TrieNode{children: map[byte]*TrieNode{}}
	for _, w := range words {
		trie.Insert(w)
	}

	res := []string{}
	for i := range board {
		for j := range board[0] {
			dfs(&res, trie, trie, board, i, j)
		}
	}

	return res
}
