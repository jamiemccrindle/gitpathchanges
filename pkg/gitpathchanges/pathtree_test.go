package gitpathchanges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathTree1(t *testing.T) {
	p := NewPathTree()
	p.Insert([]string{"one"})
	p.Insert([]string{"one", "sub"})
	p.Insert([]string{"two", "sub"})

	result1 := p.FindParents([]string{"one"})

	if len(result1) == 0 {
		t.Errorf("result1 should not be empty")
		return
	}

	if len(result1) != 1 {
		t.Errorf("result1 should only have 1 value")
		return
	}

	assert.Equal(t, result1, [][]string{
		{"one"},
	})

	assert.Equal(t, p.FindParents([]string{"one", "sub"}), [][]string{{"one"}, {"one", "sub"}})
	assert.Equal(t, p.FindParents([]string{"two"}), [][]string{})
	assert.Equal(t, p.FindParents([]string{"two", "sub"}), [][]string{{"two", "sub"}})

}

func TestPathTree2(t *testing.T) {
	p := NewPathTree()
	p.Insert([]string{"cmd"})
	p.Insert([]string{"pkg"})

	assert.Equal(t, p.FindParents([]string{"cmd", "gitpathchanges"}), [][]string{{"cmd"}})
	assert.Equal(t, p.FindParents([]string{"pkg", "gitpathchanges"}), [][]string{{"pkg"}})

}
