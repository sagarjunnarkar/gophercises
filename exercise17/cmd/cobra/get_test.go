package cobra

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
)

func TestGetSecret(t *testing.T) {
	os.Remove("/home/sagar/.secrets")
	var cmd *cobra.Command
	setCmd.Run(cmd, []string{"key_1", "value1"})
	getCmd.Run(cmd, []string{"key_1"})
	getCmd.Run(cmd, []string{"ksdfkjsfkds"})
}
