package main

import (
	"errors"
	"fmt"
	"gophercises/exercise_07/db"

	"os"
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/assert"
)

// type errorString struct {
// 	s string
// }

// func (e *errorString) Error() string {
// 	return e.s
// }

func TestPanic(t *testing.T) {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))
	err := must(errors.New("This is error"))
	assert.Equal(t, -1, err)
}

func TestCobra(t *testing.T) {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	noerrr := must(db.Init(dbPath))
	assert.Equal(t, 1, noerrr)
}
func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	setup()
	// must(cmd.RootCmd.Execute())
	runTests := m.Run()
	// spew.Dump(runTests)
	os.Exit(runTests)
}

func setup() {
	home, _ := homedir.Dir()
	fmt.Println(home)
	dbPath := filepath.Join(home, "test.db")
	os.Remove(dbPath)

	db.Init(dbPath)
}
