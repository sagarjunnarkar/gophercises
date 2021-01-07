package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"gophercises/exercise_07/db"

	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

func setup() {
	home, _ := homedir.Dir()
	fmt.Println(home)
	dbPath := filepath.Join(home, "test.db")
	os.Remove(dbPath)

	db.Init(dbPath)
}

func TestMain(m *testing.M) {
	log.Println("Do stuff BEFORE the tests!")
	setup()
	exitVal := m.Run()
	log.Println("Do stuff AFTER the tests!")

	os.Exit(exitVal)
}
func TestAddCmd(t *testing.T) {
	executeCommand(RootCmd, "add", "sagar")
}

type fakeSetup struct {
	path string
	err  error
}

func (f *fakeSetup) myUpdate(path string) (int, error) {
	return 0, f.err
}
func TestAddCmdNegative(t *testing.T) {
	f := &fakeSetup{err: errors.New("failed")}

	db.MockMyUpdate = f.myUpdate

	executeCommand(RootCmd, "add", "sagar")
}

func emptyRun(*cobra.Command, []string) {}

func executeCommand(root *cobra.Command, args ...string) (output string, err error) {
	_, output, err = executeCommandC(root, args...)
	return output, err
}

func executeCommandC(root *cobra.Command, args ...string) (c *cobra.Command, output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOutput(buf)
	root.SetArgs(args)

	c, err = root.ExecuteC()

	return c, buf.String(), err
}
