package flagparser

import (
	"fmt"
	"strings"

	"github.com/docker/swarmkit/api"
	"github.com/spf13/pflag"
)

// parseNpipe only supports a very simple version of anonymous npipes for
// testing the most basic of data flows. Replace with a --mount flag, similar
// to what we have in docker service.
func parseNpipe(flags *pflag.FlagSet, spec *api.ServiceSpec) error {
	if flags.Changed("npipe") {
		npipes, err := flags.GetStringSlice("npipe")
		if err != nil {
			return err
		}

		container := spec.Task.GetContainer()

		for _, npipe := range npipes {
			if strings.Contains(npipe, ":") {
				return fmt.Errorf("npipe format %q not supported", npipe)
			}
			container.Mounts = append(container.Mounts, api.Mount{
				Type:   api.MountTypeNpipe,
				Target: npipe,
			})
		}
	}

	return nil
}
