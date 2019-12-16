package cobra

import (
	"fmt"
	"gophercises/exercise17"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets a secret in your secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		v := exercise17.File(encodingKey, secretsPath())
		key, value := args[0], args[1]
		err := v.Set(key, value)
		if err != nil {
			fmt.Println("Something went wrong")
		}
		fmt.Println("Value set successfully!")
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}
