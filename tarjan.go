package main

import "fmt"

var (
	G            [][]int
	point_stack  []int
	marked       []bool
	marked_stack []int
)

func main() {
	G = [][]int{[]int{1}, []int{10}, []int{0}, []int{0}, []int{3}, []int{8}, []int{9}, []int{4, 5}, []int{2}, []int{6}, []int{7}}
	entry_tarjan(G)
}

func entry_tarjan(G [][]int) {
	marked = make([]bool, len(G))

	for i := 0; i < len(G); i++ {
		tarjan(i, i)
		for len(marked_stack) > 0 {
			u := marked_stack[len(marked_stack)-1]
			marked_stack = marked_stack[:len(marked_stack)-1]
			marked[u] = false
		}
	}
}

func tarjan(s int, v int) bool {
	f := false
	point_stack = append(point_stack, v)
	marked[v] = true
	marked_stack = append(marked_stack, v)

	for _, w := range G[v] {
		cb := make(chan bool, len(G[v]))
		go branch(s, v, w, cb)
		f = <-cb
	}

	if f == true {
		for marked_stack[len(marked_stack)-1] != v {
			u_ := marked_stack[len(marked_stack)-1]
			marked_stack = marked_stack[:len(marked_stack)-1]
			marked[u_] = false
		}
		marked_stack = marked_stack[:len(marked_stack)-1]
		marked[v] = false
	}

	point_stack = point_stack[:len(point_stack)-1]
	return f
}

func branch(s int, v int, w int, cb chan bool) {
	f_ := false
	if w < s {
		G[w] = []int{}
	} else if w == s {
		fmt.Println(point_stack)
		f_ = true
	} else if marked[w] == false {
		g_ := tarjan(s, w)
		f_ = f_ || g_
	}

	cb <- f_
}
