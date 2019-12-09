package db

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	log.Println("Do stuff BEFORE the tests!")
	setup()

	exitVal := m.Run()
	log.Println("Do stuff AFTER the tests!")
	// spew.Dump(runTests)
	os.Exit(exitVal)
}

func setup() {
	log.Println("setup started")
	home, _ := homedir.Dir()
	fmt.Println(home)
	dbPath := filepath.Join(home, "test.db")
	os.Remove(dbPath)

	Init(dbPath)
	log.Println("setup done")
}
func TestInit(t *testing.T) {
	t.Run("It should not raise error when path exists", func(t *testing.T) {
		home, _ := homedir.Dir()
		dbPath := filepath.Join(home, "test.db")
		os.Remove(dbPath)
		assert.Equal(t, nil, Init(dbPath))
	})

	t.Run("It should raise error when path does not exists", func(t *testing.T) {
		dbPath := filepath.Join("blabla", "test.db")
		os.Remove(dbPath)
		assert.NotEqual(t, nil, Init(dbPath))
	})
}

func TestCreateTask(t *testing.T) {
	t.Run("creates new task successfully", func(t *testing.T) {
		home, _ := homedir.Dir()
		fmt.Println(home)
		dbPath := filepath.Join(home, "test.db")
		os.Remove(dbPath)

		Init(dbPath)
		id, _ := CreateTask("blabla")
		assert.Equal(t, 1, id)
	})
}

type fakeSetup struct {
	path string
	err  error
}

func (f *fakeSetup) myUpdate(path string) (int, error) {
	return 0, f.err
}

func TestCreateTaskNegative(t *testing.T) {
	f := &fakeSetup{err: errors.New("failed")}
	// this mocks out the function that myUpdate() calls
	MockMyUpdate = f.myUpdate

	home, _ := homedir.Dir()
	fmt.Println(home)
	dbPath := filepath.Join(home, "test.db")
	os.Remove(dbPath)

	Init(dbPath)
	id, err := CreateTask("blabla")
	assert.Equal(t, -1, id)
	assert.Equal(t, "failed", err.Error())
	defer func() {
		MockMyUpdate = myUpdate // set back original func at end of test
	}()
}

func TestAllTasks(t *testing.T) {
	t.Run("Lists all tasks successfully", func(t *testing.T) {
		home, _ := homedir.Dir()
		fmt.Println(home)
		dbPath := filepath.Join(home, "test.db")
		os.Remove(dbPath)

		Init(dbPath)
		CreateTask("Task 1")
		CreateTask("Task 2")
		tasks, _ := AllTasks()

		assert.Equal(t, 2, len(tasks))
	})
}

func (f *fakeSetup) findAll() ([]Task, error) {
	return nil, f.err
}

func TestAllTasksNegative(t *testing.T) {
	f := &fakeSetup{err: errors.New("failed")}
	// this mocks out the function that findAll() calls
	MockFindAll = f.findAll

	home, _ := homedir.Dir()
	fmt.Println(home)
	dbPath := filepath.Join(home, "test.db")
	os.Remove(dbPath)

	Init(dbPath)
	CreateTask("Task 1")
	CreateTask("Task 2")
	tasks, err := AllTasks()

	assert.Equal(t, 0, len(tasks))
	assert.Equal(t, "failed", err.Error())

	defer func() {
		MockFindAll = FindAll // set back original func at end of test
	}()
}

func TestDeleteTask(t *testing.T) {
	t.Run("Deletes task successfully", func(t *testing.T) {
		home, _ := homedir.Dir()
		fmt.Println(home)
		dbPath := filepath.Join(home, "test.db")
		os.Remove(dbPath)

		Init(dbPath)
		CreateTask("Task 1")
		err := DeleteTask(1)

		assert.Equal(t, nil, err)
	})
}
