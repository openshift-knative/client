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
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"knative.dev/client/pkg/kn/commands"
	"knative.dev/client/pkg/kn/config"
)

// NewPluginInstallCommand creates a new `kn plugin install` command
func NewPluginInstallCommand(p *commands.KnParams) *cobra.Command {

	var localPath string

	cmd := &cobra.Command{
		Use:     "install",
		Short:   "Install a plugin from URL or local file",
		Aliases: []string{"ls"},
		Long: `Install a plugin from URL.or local filesystem.

Plugin is installed from a provided URL to ${XDG_HOME}/kn/plugins.

Current: ` + config.GlobalConfig.PluginsDir(),
		RunE: func(cmd *cobra.Command, args []string) error {
			var fileName, urlFile string
			if localPath != "" {
				segments := strings.Split(localPath, string(os.PathSeparator))
				fileName = filepath.Join(config.GlobalConfig.PluginsDir(), segments[len(segments)-1])
			} else {
				if len(args) != 1 {
					return errors.New("'kn plugin install' requires URL")
				}
				urlFile = args[0]

				// Build fileName from fullPath
				fileURL, err := url.Parse(urlFile)
				if err != nil {
					return err
				}
				path := fileURL.Path
				segments := strings.Split(path, "/")
				fileName = filepath.Join(config.GlobalConfig.PluginsDir(), segments[len(segments)-1])
			}

			file, err := os.Create(fileName)
			if err != nil {
				return err
			}
			defer file.Close()

			if urlFile != "" {
				client := &http.Client{
					CheckRedirect: func(req *http.Request, via []*http.Request) error {
						return http.ErrUseLastResponse
					},
				}

				// Download file
				resp, err := client.Get(urlFile)
				if err != nil {
					println(err)
				}
				defer resp.Body.Close()
				return installPlugin(fileName, file, resp.Body)

			}
			localFile, err := os.Open(localPath)
			if err != nil {
				return err
			}
			return installPlugin(fileName, file, localFile)
		},
	}

	cmd.Flags().StringVarP(&localPath, "file", "f", "", "Path to install plugin from local filesystem")

	return cmd
}

func installPlugin(fileName string, dst io.Writer, src io.Reader) error {
	_, err := io.Copy(dst, src)
	if err != nil {
		return err
	}
	err = os.Chmod(fileName, os.FileMode(0755))
	if err != nil {
		return err
	}
	fmt.Printf("Plugin installed to '%s'", fileName)
	return nil
}
