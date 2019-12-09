package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")
var db *bolt.DB

// Task is struct which defines key value pair.
// Here key is serial number of Task and value is actual task.
type Task struct {
	Key   int
	Value string
}

//Init function is to create new bolt db connection and to create bolt db bucket
func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

//MockMyUpdate is variable which is used in mocking bolt's update function.
var MockMyUpdate = myUpdate

func myUpdate(task string) (int, error) {
	var id int

	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)
		return b.Put(key, []byte(task))
	})

	return id, err
}

//CreateTask is a function to insert new task in bolt db
func CreateTask(task string) (int, error) {
	id, err := MockMyUpdate(task)

	if err != nil {
		return -1, err
	}

	return id, nil
}

//FindAll is a function to list all the tasks.
func FindAll() ([]Task, error) {
	var tasks []Task

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})

	return tasks, err
}

//MockFindAll is mock variable used for go unit testing.
var MockFindAll = FindAll

//AllTasks is a function to list all tasks.
func AllTasks() ([]Task, error) {

	tasks, err := MockFindAll()

	if err != nil {
		return nil, err
	}
	return tasks, nil
}

//MyDeleteTask is a function for deleting object from bolt db.
func MyDeleteTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})
}

//MockDeleteTask is a mock variable used in go unit testing.
var MockDeleteTask = MyDeleteTask

//DeleteTask is a exported function used to delete a task from bolt db.
func DeleteTask(key int) error {
	return MockDeleteTask(key)
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
