// Copyright © 2024 The Knative Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package plugin

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"knative.dev/client/pkg/kn/commands"
	"knative.dev/client/pkg/kn/config"
)

// NewPluginInstallCommand creates a new `kn plugin install` command
func NewPluginRemoveCommand(p *commands.KnParams) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "remove",
		Short:   "remove a plugin from local config directory",
		Aliases: []string{"ls"},
		Long: `Install a plugin from URL or local filesystem.

Plugin is removed by name from a local directory ${KN_CONFIG_DIR}/kn/plugins.

Current: ` + config.GlobalConfig.PluginsDir(),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("'kn plugin remove' requires plugin name")
			}
			name := args[0]
			filename := filepath.Join(config.GlobalConfig.PluginsDir(), name)
			if _, err := os.Stat(filename); os.IsNotExist(err) {
				return errors.New("plugin '" + name + "' does not exist")
			}
			println("Removed plugin: '" + filename + "'")
			return os.Remove(filename)
		},
	}
	return cmd
}
