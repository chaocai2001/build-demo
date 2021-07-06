package buildtest

import (
	"context"

	"os"
	"testing"

	"github.com/chaocai2001/build-demo/gitops"
	"github.com/docker/docker/api/types"
)

func TestBuilding(t *testing.T) {
	tmpDir := "./tmp"
	repoURL := "https://github.com/chaocai2001/build-demo.git"
	image := "chaocai/build-demo:v0.0.7"
	err := os.Mkdir(tmpDir, 0755)
	if err != nil {
		t.Error(err)
		return
	}
	defer os.RemoveAll(tmpDir)
	git := gitops.NewGitConnector()
	err = git.CloneWithBranch(repoURL, tmpDir, "main")
	if err != nil {
		t.Error(err)
		return
	}
	builder := gitops.NewBuilder(gitops.BuilderType_PAKETO_BASE)
	ctx := context.Background()
	builder.Build(ctx, tmpDir, image)
	var authConfig = types.AuthConfig{
		Username: "chaocai",
		Password: "superman",
	}
	rd, err := builder.PushToHub(ctx, authConfig, image)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(rd)
}
