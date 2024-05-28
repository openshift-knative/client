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

var updateFlag bool

// NewPluginInstallCommand creates a new `kn plugin install` command
func NewPluginInstallCommand(p *commands.KnParams) *cobra.Command {

	var localPath string
	cmd := &cobra.Command{
		Use:     "install",
		Short:   "Install a plugin from URL or local file",
		Aliases: []string{"ls"},
		Long: `Install a plugin from URL or local filesystem.

Plugin is installed from a provided URL or local path to ${KN_CONFIG_DIR}/kn/plugins.

Current: ` + config.GlobalConfig.PluginsDir(),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				if localPath == "" {
					return errors.New("'kn plugin install' requires URL or local file parameter '-f' to be specified")
				}
				println("Copying plugin from local path: " + localPath)
				// Local plugin file to copy
				return fetchLocal(localPath)
			}
			println("Downloading from URL: " + args[0])
			// Download from provided external URL and copy to plugins dir
			return fetchFromURL(args[0])
		},
	}
	cmd.Flags().StringVarP(&localPath, "file", "f", "", "Path to install plugin from local filesystem")
	cmd.Flags().BoolVarP(&updateFlag, "update", "u", false, "Update installed plugin")
	return cmd
}

func parseFilePath(src, separator string) string {
	segments := strings.Split(src, separator)
	return filepath.Join(config.GlobalConfig.PluginsDir(), segments[len(segments)-1])
}

func fetchLocal(localPath string) error {
	localFile, err := os.Open(localPath)
	if err != nil {
		return err
	}
	fileName := parseFilePath(localPath, string(os.PathSeparator))

	if err := updatePluginCheck(fileName); err != nil {
		return err
	}

	dest, err := os.Create(fileName)
	if err != nil {
		return err
	}
	return installPlugin(fileName, dest, localFile)
}

func fetchFromURL(externalUrl string) error {
	// Build fileName from fullPath
	fileURL, err := url.Parse(externalUrl)
	if err != nil {
		return err
	}
	path := fileURL.Path
	// Assuming URL, hence forward-slash separator
	fileName := parseFilePath(path, "/")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	// Download file
	resp, err := client.Get(fileURL.String())
	if err != nil {
		println(err)
	}
	defer resp.Body.Close()

	if err := updatePluginCheck(fileName); err != nil {
		return err
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	return installPlugin(fileName, file, resp.Body)
}

func updatePluginCheck(path string) error {
	_, err := os.Stat(path)
	if !os.IsNotExist(err) {
		if !updateFlag {
			return errors.New("plugin already exists, please use `--update` flag to replace it at '" + path + "'")
		}
	}
	return nil
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
	println("Plugin installed to '" + fileName + "'")
	return nil
}
