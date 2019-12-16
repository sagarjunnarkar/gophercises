package exercise17

import (
	"crypto/cipher"
	"errors"
	"os"

	"github.com/davecgh/go-spew/spew"

	ciphers "gophercises/exercise17/ciphers"
	"testing"

	"github.com/stretchr/testify/assert"
)

type fakeSetup struct {
	path string
	err  error
}

func (f *fakeSetup) setup(path string) (cipher.Block, error) {
	f.path = path
	return nil, f.err
}

func TestSet(t *testing.T) {
	t.Run("sets keys successfully", func(t *testing.T) {
		v := File("", "/home/sagar/.secrets")
		os.Remove(v.filepath)
		key, value := "key_1", "value_1"
		err := v.Set(key, value)
		spew.Dump(err)
		assert.Equal(t, nil, err)
	})

	t.Run("throws error when key is invalid", func(t *testing.T) {
		v := File("key", "/home/sagar/.secrets")
		key, value := "key_1", "value_1"
		err := v.Set(key, value)
		assert.NotEqual(t, nil, err)
	})

	t.Run("throws error when path is invalid", func(t *testing.T) {
		v := File("", "/home/invalid/.s")
		key, value := "key_1", "value_1"
		err := v.Set(key, value)
		assert.NotEqual(t, nil, err)
	})

	t.Run("returns error when save() doesn't work", func(t *testing.T) {
		f := &fakeSetup{err: errors.New("Failed")}
		v := File("", "/home/sagar/.secrets")
		ciphers.MockNewCipherBlock = f.setup
		key, value := "key_1", "value_1"
		err := v.Set(key, value)
		assert.NotEqual(t, nil, err)
		defer func() {
			ciphers.MockNewCipherBlock = ciphers.NewCipherBlock // set back original func at end of test
		}()
	})
}

func TestGet(t *testing.T) {
	t.Run("Get key successfully", func(t *testing.T) {
		v := File("", "/home/sagar/.secrets")
		key, value := "key_1", "value_1"
		err := v.Set(key, value)
		_, errr := v.Get(key)

		assert.Equal(t, nil, err)
		assert.Equal(t, nil, errr)
	})

	t.Run("Get key successfully", func(t *testing.T) {
		v := File("", "/home/sagar/.secrets")
		_, errr := v.Get("kkkk")

		// assert.Equal(t, nil, err)
		assert.NotEqual(t, nil, errr)
	})

	t.Run("returns error when load does not work", func(t *testing.T) {
		v := File("", "/home/sagar/.secrets")
		key, value := "key_1", "value_1"
		err := v.Set(key, value)
		f := &fakeSetup{err: errors.New("Failed")}
		ciphers.MockNewCipherBlock = f.setup
		_, errr := v.Get(key)

		assert.Equal(t, nil, err)
		assert.NotEqual(t, nil, errr)
		defer func() {
			ciphers.MockNewCipherBlock = ciphers.NewCipherBlock // set back original func at end of test
		}()
	})
}
