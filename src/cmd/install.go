/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
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
	Long: `"Load CSAR, validate contents (== try to parse to TOSCA service template) and detect primary workflow in . For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			csarPath string = args[0]
			archive  csar.CSAR
		)

		if debug {
			// Set debug for imports
			csar.Debug()

			log.Println("INF debug:", debug)

			log.Println("INF csarPath:", csarPath)
		}

		archive = csar.LoadFromPath(csarPath)
		if debug {
			log.Println("SUC Loaded CSAR from '" + csarPath + "'.")
		}

		// allow named inputs -> add those to archive.ServiceTemplate.TopologyTemplate.Inputs

		// for _, x := range archive.ServiceTemplate.DataTypes {
		// 	x.ValidateConstraints() // <- missing value to validate
		// }

		toscaorchestrator.Install(archive.ServiceTemplate)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
