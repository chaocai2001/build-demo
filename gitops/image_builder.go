package gitops

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"

	"io"

	"github.com/buildpacks/pack"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	//	"github.com/docker/docker/pkg/archive"
)

type BuilderType string

const (
	BuilderType_GCR_Builder_V1 BuilderType = "gcr.io/buildpacks/builder:v1"
	BuilderType_HEROKU_20                  = "heroku/buildpacks:20"
	BuilderType_PAKETO_BASE                = "gcr.io/paketo-buildpacks/builder:base"
)

type Builder struct {
	builder BuilderType
}

func NewBuilder(buildType BuilderType) *Builder {
	return &Builder{buildType}
}

func (b *Builder) Build(ctx context.Context, appPath string, targetImage string) error {

	//initialize a pack client
	client, err := pack.NewClient()
	if err != nil {
		panic(err)
	}

	// initialize our options
	buildOpts := pack.BuildOptions{
		Image:        targetImage,
		Builder:      string(b.builder),
		AppPath:      appPath,
		TrustBuilder: true,
	}
	return client.Build(ctx, buildOpts)

}

func (b *Builder) PushToHub(ctx context.Context, authConfig types.AuthConfig,
	image string) (string, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err.Error())
	}

	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		panic(err)
	}
	authStr := base64.URLEncoding.EncodeToString(encodedJSON)
	var pushReader io.ReadCloser

	pushReader, err = cli.ImagePush(context.Background(), image, types.ImagePushOptions{
		All:           false,
		RegistryAuth:  authStr,
		PrivilegeFunc: nil,
	})
	if err != nil {
		return "", err
	}
	buf1 := new(bytes.Buffer)
	buf1.ReadFrom(pushReader)
	s1 := buf1.String()

	return s1, err

}
