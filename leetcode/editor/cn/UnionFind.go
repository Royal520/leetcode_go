package main

type UnionFind struct {
	size     []int
	parent   []int
	n        int
	setCount int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &UnionFind{
		n:        n,
		setCount: n,
		size:     size,
		parent:   parent,
	}
}

func (find *UnionFind) FindRoot(x int) int {
	if find.parent[x] == x {
		return x
	}
	return find.FindRoot(find.parent[x])
}

func (find *UnionFind) Unite(x, y int) bool {
	x = find.FindRoot(x)
	y = find.FindRoot(y)
	if x == y {
		return false
	}
	if find.size[x] < find.size[y] {
		temp := x
		x = y
		y = temp
	}
	find.size[x] += find.size[y]
	find.parent[y] = x
	find.setCount--
	return true
}

func (find *UnionFind) Connected(x, y int) bool {
	return find.FindRoot(x) == find.FindRoot(y)
}
