package gitpathchanges

import (
	"fmt"
	"strings"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func getCommit(r *git.Repository, ref string) (*object.Commit, error) {
	hash := plumbing.NewHash(ref)
	commit, err := r.CommitObject(hash)
	if err != nil {
		return nil, fmt.Errorf("commit could not be found")
	}
	if commit == nil {
		return nil, fmt.Errorf("commit could not be found")
	}
	return commit, nil
}

func Files(path string, pathsToMatch []string, commitRef1 string, commitRef2 string) (*[]string, error) {
	r, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}
	commit1, err := getCommit(r, commitRef1)
	if err != nil {
		return nil, err
	}
	commit2, err := getCommit(r, commitRef2)
	if err != nil {
		return nil, err
	}
	patch, err := commit1.Patch(commit2)
	if err != nil {
		return nil, err
	}
	filePatches := patch.FilePatches()
	changedPaths := make(map[string]struct{})
	for _, filePatch := range filePatches {
		from, to := filePatch.Files()
		if from != nil {
			changedPaths[from.Path()] = struct{}{}
		}
		if to != nil {
			changedPaths[to.Path()] = struct{}{}
		}
	}
	result := []string{}
	if len(pathsToMatch) == 0 {
		for k := range changedPaths {
			result = append(result, k)
		}
	} else {
		matched := make(map[string]struct{})
		p := NewPathTree()
		for _, d := range pathsToMatch {
			p.Insert(strings.Split(d, "/"))
		}
		for k := range changedPaths {
			for _, parent := range p.FindParents(strings.Split(k, "/")) {
				matched[strings.Join(parent, "/")] = struct{}{}
			}
		}
		for k := range matched {
			result = append(result, k)
		}
	}
	return &result, nil
}
