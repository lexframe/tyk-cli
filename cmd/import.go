package cmd

import (
	"github.com/spf13/cobra"

	"github.com/TykTechnologies/tyk-cli/commands/importpkg"
)

var input string

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import API definitions",
	Long:  `This module lets you import previously exported API definitions from a JSON file.`,
	Run: func(cmd *cobra.Command, args []string) {
		args = []string{key, domain, port, input}
		importType := map[string]func([]string){
			"api": importpkg.APIs,
		}
		importType[cmd.Parent().Name()](args)
	},
}

func init() {
	apiCmd.AddCommand(importCmd)

	importCmd.Flags().StringVarP(&key, "key", "k", "", "Secret Key for the Dashboard API")
	importCmd.Flags().StringVarP(&domain, "domain", "d", "", "Domain name for your Dashboard")
	importCmd.Flags().StringVarP(&port, "port", "p", "", "Port number for your Dashboard")
	importCmd.Flags().StringVarP(&input, "input", "i", "", "Input file name for your JSON string")
}
