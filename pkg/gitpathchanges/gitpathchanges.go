package gitpathchanges

import (
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func Files(path string, directories []string, commitRef1 string, commitRef2 string) (*[]string, error) {
	r, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}
	commit1, err := r.CommitObject(plumbing.NewHash(commitRef1))
	if err != nil {
		return nil, err
	}
	commit2, err := r.CommitObject(plumbing.NewHash(commitRef2))
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
		changedPaths[from.Path()] = struct{}{}
		changedPaths[to.Path()] = struct{}{}
	}
	result := []string{}
	for k, _ := range changedPaths {
		result = append(result, k)
	}
	return &result, nil
}
