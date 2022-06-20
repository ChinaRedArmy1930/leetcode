package main

import (
	"fmt"
	"strconv"
	"strings"
)

//TODO
func traversal(board [][]byte, word string, x, y int, acrossWord *string, visited *map[string]bool, found *bool) {
	if *found {
		return
	}

	if _, ok := (*visited)[strconv.Itoa(x)+strconv.Itoa(y)]; ok {
		return
	}

	if x >= len(board) || y >= len(board[0]) || x < 0 || y < 0 {
		return
	}
	*acrossWord += string(board[x][y])

	if !strings.Contains(word, *acrossWord) {
		*acrossWord = (*acrossWord)[:len(*acrossWord)-1]
		return
	}

	if (*acrossWord) == word {
		*found = true
		return
	}

	if len(*acrossWord) > len(word) {
		*acrossWord = (*acrossWord)[:len(*acrossWord)-1]
		return
	}

	(*visited)[strconv.Itoa(x)+strconv.Itoa(y)] = true
	if x+1 < len(board) {
		traversal(board, word, x+1, y, acrossWord, visited, found)
	}

	if y+1 < len(board[0]) {
		traversal(board, word, x, y+1, acrossWord, visited, found)
	}

	if x-1 >= 0 {
		traversal(board, word, x-1, y, acrossWord, visited, found)
	}

	if y-1 >= 0 {
		traversal(board, word, x, y-1, acrossWord, visited, found)
	}

	*acrossWord = (*acrossWord)[:len(*acrossWord)-1]
	delete(*visited, strconv.Itoa(x)+strconv.Itoa(y))
}

func exist(board [][]byte, word string) bool {
	acrossWord := ""
	found := false
	visited := make(map[string]bool)
	first := string(word[0])

	for i, bytes := range board {
		for j, b := range bytes {
			if string(b) == first && !found {
				traversal(board, word, i, j, &acrossWord, &visited, &found)
			}
		}
	}

	return found
}

func main() {
	fmt.Println(exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCCED"))
	//fmt.Println(exist([][]byte{
	//	{'A', 'B', 'C', 'E'},
	//	{'S', 'F', 'C', 'S'},
	//	{'A', 'D', 'E', 'E'},
	//	{'A', 'B', 'C', 'E'},
	//}, "ADEEA"))
	//fmt.Println(exist([][]byte{
	//	{'A', 'B', 'C', 'E'},
	//	{'S', 'F', 'C', 'S'},
	//	{'A', 'D', 'E', 'E'},
	//	{'A', 'B', 'C', 'E'},
	//}, "SEE"))
	//fmt.Println(exist([][]byte{
	//	{'A', 'B', 'C', 'E'},
	//	{'S', 'F', 'C', 'S'},
	//	{'A', 'D', 'E', 'E'},
	//	{'A', 'B', 'C', 'E'},
	//}, "ABCB"))
	//fmt.Println(exist([][]byte{
	//	{'A'},
	//}, "A"))

	//[["b","b","a","a","b","a"],["b","b","a","b","a","a"],["b","b","b","b","b","b"],["a","a","a","b","a","a"],["a","b","a","a","b","b"]]
	//now := time.Now()
	//fmt.Println(exist([][]byte{
	//	{'b', 'b', 'a', 'a', 'b', 'a'},
	//	{'b', 'b', 'a', 'b', 'a', 'a'},
	//	{'b', 'b', 'b', 'b', 'b', 'b'},
	//	{'a', 'b', 'a', 'a', 'b', 'b'},
	//}, "abbbababaa"))
	//fmt.Println(time.Now().Sub(now).Seconds())
}
