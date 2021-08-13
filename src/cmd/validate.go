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
	Long: `Load CSAR from zip-archive (proper CSAR, therefore default) OR from a folder containing the extracted contents of a CSAR.
	Usage examples:

	eat validate example-csar.zip
	eat validate -d example-csar/
	eat validate --directory example-csar/
`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var (
			csarPath string
			archive  csar.CSAR
		)

		if debug {
			log.Println("debug:", debug)
			// Set debug for imports
			csar.Debug = debug
		}

		// Retrieve flag
		extractedCSAR, err := cmd.Flags().GetString("directory")
		if err != nil {
			log.Fatalln("ERR There was an error while reading the flag 'directory':", err)
		}

		// Validate usage and run CSAR validation
		if len(args) == 1 && extractedCSAR == "" { // Load CSAR from zip-archive
			log.Fatalln("// TODO: Implement LoadFromFile")
			csarPath = args[0]
			archive = csar.LoadFromFile(csarPath)
			if debug {
				log.Println("SUC Loaded CSAR from file at '" + csarPath + "'.")
			}
			csar.PrintCSAR(archive)
		} else if len(args) == 0 && extractedCSAR != "" { // Load extracted CSAR from directory
			archive = csar.LoadFromFolder(extractedCSAR)
			if debug {
				log.Println("SUC Loaded CSAR from folder at '" + extractedCSAR + "'.")
			}
			csar.PrintCSAR(archive)
		} else if len(args) == 0 && extractedCSAR == "" {
			log.Fatalln("ERR No CSAR source provided.")
		} else {
			log.Fatalln("ERR Concurrent loading of CSAR from zip-archive and from directory is not possible.")
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

	validateCmd.Flags().StringP("directory", "d", "", "Load extracted CSAR contents from directoy instead of zip-archive.")
}
