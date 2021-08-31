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
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/thetillhoff/eat/pkg/arp"
	"github.com/thetillhoff/eat/pkg/ssh"
	"github.com/thetillhoff/eat/pkg/wol"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test called")
		//
		//
		//

		if debug {
			// Set debug for imports
			wol.Debug()

			log.Println("INF debug:", debug)
		}

		var macAddress string = "00:15:5d:cf:8b:50"

		// // Test for dhcp
		// dhcp_and_tftp.Start()

		// Test for wake on lan
		wol.Wake(macAddress)
		log.Println("INF Waked machine with mac '" + macAddress + "'.")

		// Get ip for mac
		var ipAddress string = arp.GetIpForMac(macAddress)
		log.Println("INF IP-Address:", ipAddress)

		// Load private-ssh-key
		ssh.CheckKeys()
		ssh.RunCommandOnHost(ipAddress, "hostname")

		//
		//
		//
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
