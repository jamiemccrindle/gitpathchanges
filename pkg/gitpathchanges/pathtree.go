package gitpathchanges

import "strings"

type PathTree struct {
	children *map[string]*PathTree
	terminal bool
}

func NewPathTree() PathTree {
	return PathTree{terminal: false}
}

func (t *PathTree) recursePaths(separator string, pathSoFar []string) []string {
	result := []string{}
	if t.terminal {
		result = append(result, strings.Join(pathSoFar, separator))
	}
	if t.children != nil && len(*t.children) > 0 {
		for k, v := range *t.children {
			result = append(result, v.recursePaths(separator, append(pathSoFar, k))...)
		}
	}
	return result
}

func (t *PathTree) Paths(separator string) []string {
	return t.recursePaths(separator, []string{})
}

func (t *PathTree) Insert(path []string) {
	current := t
	for _, p := range path {
		if current.children == nil {
			children := make(map[string]*PathTree)
			child := NewPathTree()
			children[p] = &child
			current.children = &children
			current = &child
		} else {
			if child, ok := (*current.children)[p]; ok {
				current = child
			} else {
				child := NewPathTree()
				(*current.children)[p] = &child
				current = &child
			}
		}
	}
	current.terminal = true
}

func (t *PathTree) FindParents(path []string) [][]string {
	current := t
	matched := [][]string{}
	pathSoFar := []string{}
	for _, p := range path {
		pathSoFar = append(pathSoFar, p)
		if child, ok := (*current.children)[p]; ok {
			current = child
		} else {
			return matched
		}
		if current.terminal {
			matched = append(matched, pathSoFar)
		}
		if current.children == nil {
			break
		}
	}
	return matched
}
