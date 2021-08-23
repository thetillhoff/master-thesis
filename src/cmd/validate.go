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
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Load CSAR and validate contents (== try to parse to TOSCA service template)",
	Long: `Load CSAR either from zip-archive (proper CSAR, detected by extension of provided path) OR from a folder containing the extracted contents of a CSAR.
	Usage examples:

	eat validate example-csar.zip
	eat validate example-csar/
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var (
			csarPath string = args[0]
			archive  csar.CSAR
		)

		if debug {
			// Set debug for imports
			csar.Debug = debug

			log.Println("INF debug:", debug)

			log.Println("INF csarPath:", csarPath)
		}

		archive = csar.LoadFromPath(csarPath)
		if debug {
			log.Println("SUC Loaded CSAR from '" + csarPath + "'.")
		}

		// for _, x := range archive.ServiceTemplate.DataTypes {
		// 	x.ValidateConstraints() // <- missing value to validate
		// }

		//archive.ServiceTemplate = archive.ResolveDerivations()

		// archive.Print()

		if debug {
			log.Println("len(nodeTypes):", len(archive.ServiceTemplate.NodeTypes))
		}
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// validateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// validateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
