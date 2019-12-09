package cmd

import (
	"errors"
	"fmt"
	"gophercises/task/db"
	"os"
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
)

func (f *fakeSetup) mockFindAll() ([]db.Task, error) {
	return nil, f.err
}
func (f *fakeSetup) mockDeleteAll(i int) error {
	return f.err
}

func TestDoCmd(t *testing.T) {
	t.Run("It should not raise error when path exists", func(t *testing.T) {
		executeCommand(RootCmd, "add", "sagar")
		executeCommand(RootCmd, "do", "2")
	})

	t.Run("It should raise error when argument is not number", func(t *testing.T) {
		home, _ := homedir.Dir()
		fmt.Println(home)
		dbPath := filepath.Join(home, "test.db")
		os.Remove(dbPath)

		db.Init(dbPath)

		executeCommand(RootCmd, "add", "sagar")
		executeCommand(RootCmd, "do", "two")
	})

	t.Run("It should raise error for db.findall", func(t *testing.T) {
		f := &fakeSetup{err: errors.New("failed")}
		db.MockFindAll = f.mockFindAll
		executeCommand(RootCmd, "add", "sagar")
		executeCommand(RootCmd, "do", "1")

		defer func() {
			db.MockFindAll = db.FindAll
		}()
	})

	t.Run("It should raise error for db.DeleteAll", func(t *testing.T) {
		f := &fakeSetup{err: errors.New("failed")}
		db.MockDeleteTask = f.mockDeleteAll

		executeCommand(RootCmd, "add", "sagar")
		executeCommand(RootCmd, "do", "1")
		defer func() {
			db.MockDeleteTask = db.MyDeleteTask
		}()
	})
}
