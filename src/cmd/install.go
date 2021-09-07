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

	"github.com/spf13/cobra"
	"github.com/thetillhoff/eat/pkg/csar"
	toscaorchestrator "github.com/thetillhoff/eat/pkg/tosca_orchestrator"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Load CSAR and install contained topologyTemplate",
	Long: `"Load CSAR, validate contents (== try to parse to TOSCA service template) and deploy the topologyTemplate. For example:

eat install some-csar-file.zip
eat install some-csar-file.zip --input DOMAIN=example.tld`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var (
			csarPath string = args[0]
			archive  csar.CSAR
			err      error
			inputs   []string
			bindIp   string
		)

		inputs, err = cmd.PersistentFlags().GetStringSlice("input")
		if err != nil {
			log.Fatalln(err)
		}

		bindIp, err = cmd.PersistentFlags().GetString("bind-ip")
		if err != nil {
			log.Fatalln(err)
		}

		if debug {
			// Set debug for imports
			csar.Debug()
			toscaorchestrator.Debug()

			log.Println("INF debug:", debug)
			log.Println("INF csarPath:", csarPath)
			log.Println("INF inputs:", inputs)
			log.Println("INF bindIp:", bindIp)
		}

		archive = csar.LoadFromPath(csarPath)
		if debug {
			log.Println("SUC Loaded CSAR from '" + csarPath + "'.")
		}

		// allow named inputs -> add those to archive.ServiceTemplate.TopologyTemplate.Inputs

		// for _, x := range archive.ServiceTemplate.DataTypes {
		// 	x.ValidateConstraints() // <- missing value to validate
		// }

		toscaorchestrator.Install(archive, inputs, bindIp)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	installCmd.PersistentFlags().StringSliceP("input", "i", []string{}, "Define inputs for CSAR's TopologyTemplate, f.e. '-i port=443'. Multiple inputs are possible.")

	installCmd.PersistentFlags().String("bind-ip", "0.0.0.0", "Define the bind-ip for the dhcp-, tftp- and http-server, f.e. '--bind-ip=0.0.0.0'.")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
