package cmd

import (
	"errors"
	"gophercises/task/db"
	"testing"
)

func TestListCmd(t *testing.T) {
	t.Run("It should not raise error when path exists", func(t *testing.T) {
		executeCommand(RootCmd, "add", "sagar")
		executeCommand(RootCmd, "add", "sagar")
		executeCommand(RootCmd, "add", "sagar")
		executeCommand(RootCmd, "list")

		// assert.Equal(t, nil, err)
	})

	t.Run("It should not raise error when path exists", func(t *testing.T) {
		f := &fakeSetup{err: errors.New("failed")}
		db.MockFindAll = f.mockFindAll

		executeCommand(RootCmd, "add", "sagar")
		executeCommand(RootCmd, "list")

		defer func() {
			db.MockFindAll = db.FindAll // set back original func at end of test
		}()
	})
}
