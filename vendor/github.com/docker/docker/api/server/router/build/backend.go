package build

import (
	"io"

	"github.com/docker/docker/builder"
	"github.com/docker/engine-api/types"
	"golang.org/x/net/context"
)

// Backend abstracts an image builder whose only purpose is to build an image referenced by an imageID.
type Backend interface {
	// Build builds a Docker image referenced by an imageID string.
	//
	// Note: Tagging an image should not be done by a Builder, it should instead be done
	// by the caller.
	//
	// TODO: make this return a reference instead of string
	Build(clientCtx context.Context, config *types.ImageBuildOptions, context builder.Context, stdout io.Writer, stderr io.Writer, out io.Writer, clientGone <-chan bool) (string, error)
}
