/*
Copyright © 2021 Till Hoffmann <till.hoffmann@enforge.de>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"log"
	"os"
	"path"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/spf13/cobra"
	"github.com/thetillhoff/eat/pkg/docker"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Prepares all prerequisites for provisioning",
	Long: `Creates the live-os iso-file and builds the dhcp-container. For example:

eat init`,
	Run: func(cmd *cobra.Command, args []string) {

		if debug {
			// Set debug for imports
			docker.Debug()

			log.Println("INF debug:", debug)
		}

		var containerPath string = "live-os"

		docker.Init()
		docker.BuildImage(containerPath, "live-os-builder", true)

		pwd, err := os.Getwd()
		if err != nil {
			log.Fatalln("ERR Couldn't retrieve working directory:", err)
		}
		sourcePath := path.Join(pwd, containerPath, "/container")

		containerID := docker.StartWithAutoStop("live-os-builder", container.HostConfig{
			AutoRemove: true,
			Mounts: []mount.Mount{
				{
					Type:   mount.TypeBind,
					Source: sourcePath,
					Target: "/container",
				},
			},
		})
		docker.Wait(containerID)

		log.Println("SUC Created live-os iso-file at '" + sourcePath + "'.")

		// mv live-os/custom.iso dnsmasq/isos/
		err = os.Rename(path.Join(sourcePath, "debian-live-11.0.0-custom.iso"), "dnsmasq/isos/debian-live-11.0.0-custom.iso")
		if err != nil {
			log.Fatal(err)
		}

		// TODO
		// Build dnsmasq container image
		docker.BuildImage("dnsmasq", "dnsmasq", true)
		log.Println("SUC Created dnsmasq container image.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
