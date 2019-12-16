package cobra

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestSetSecret(t *testing.T) {
	var cmd *cobra.Command

	setCmd.Run(cmd, []string{"set", "key_2", "value2"})
	tmp := encodingKey
	encodingKey = "bla"
	setCmd.Run(cmd, []string{"set", "key_3", "value3"})
	defer func() {
		encodingKey = tmp
	}()
}
