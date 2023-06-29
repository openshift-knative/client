package volume // import "github.com/docker/docker/api/types/volume"

import "github.com/docker/docker/api/types/filters"

// ListOptions holds parameters to list volumes.
type ListOptions struct {
	Filters filters.Args
}
