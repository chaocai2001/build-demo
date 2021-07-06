package gitops

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type CodeRepository interface {
	CloneWithBranch(repoURL string, targetDir string,
		branch string) error
	CloneWithTag(repoURL string, targetDir string,
		tag string) error
}

type CommitObj struct {
	Reference string
}

type GitConnector struct {
}

func NewGitConnector() CodeRepository {
	return &GitConnector{}
}

func (gitC *GitConnector) CloneWithBranch(repoURL string, targetDir string,
	branch string) error {
	refName := plumbing.NewBranchReferenceName(branch)
	return gitC.CloneWithRef(repoURL, targetDir, refName)
}

func (gitC *GitConnector) CloneWithTag(repoURL string, targetDir string,
	tag string) error {
	refName := plumbing.NewTagReferenceName(tag)
	return gitC.CloneWithRef(repoURL, targetDir, refName)
}

func (gitC *GitConnector) CloneWithRef(repoURL string, targetDir string,
	ref plumbing.ReferenceName) error {
	_, err := git.PlainClone(targetDir, false, &git.CloneOptions{
		URL:               repoURL,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		ReferenceName:     ref,
	})

	if err != nil {
		return err
	}
	return nil
}
