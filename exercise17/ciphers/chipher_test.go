package ciphers

import (
	"crypto/cipher"
	"errors"
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Vault struct {
	encodingKey string
	filepath    string
	mutex       sync.Mutex
	keyValues   map[string]string
}

type fakeSetup struct {
	path string
	err  error
}

func (f *fakeSetup) setup(path string) (cipher.Block, error) {
	f.path = path
	return nil, f.err
}
func TestEncryptWriter(t *testing.T) {
	t.Run("it does not return error when filepath is right", func(t *testing.T) {
		var v = Vault{encodingKey: "", filepath: "/home/sagar/.secrets"}
		os.Remove(v.filepath)
		f, err := os.OpenFile(v.filepath, os.O_RDWR|os.O_CREATE, 0755)

		if err != nil {
			v.keyValues = make(map[string]string)
			t.Errorf(err.Error())
		}
		_, err = EncryptWriter(v.encodingKey, f)
		assert.Equal(t, nil, err)
		f.Close()
	})

	t.Run("it returns error when file is not opened in read/write mode", func(t *testing.T) {
		var v = Vault{encodingKey: "", filepath: "/home/sagar/.secrets"}
		f, err := os.Open(v.filepath)
		if err != nil {
			v.keyValues = make(map[string]string)
			t.Errorf(err.Error())
		}
		_, err = EncryptWriter(v.encodingKey, f)
		assert.NotEqual(t, nil, err)
		f.Close()
	})

	t.Run("it returns error when newCipherBlock does not work", func(t *testing.T) {
		var v = Vault{encodingKey: "", filepath: "/home/sagar/.secrets"}

		f, err := os.OpenFile(v.filepath, os.O_RDWR|os.O_CREATE, 0755)

		if err != nil {
			v.keyValues = make(map[string]string)
			t.Errorf(err.Error())
		}
		ff := &fakeSetup{err: errors.New("Failed")}
		MockNewCipherBlock = ff.setup
		_, err = EncryptWriter(v.encodingKey, f)
		assert.NotEqual(t, nil, err)
		f.Close()
		defer func() {
			MockNewCipherBlock = newCipherBlock // set back original func at end of test
		}()
	})
}

func TestDecryptReader(t *testing.T) {
	t.Run("it does not return error when filepath is right", func(t *testing.T) {
		var v = Vault{encodingKey: "", filepath: "/home/sagar/.secrets"}
		f, err := os.OpenFile(v.filepath, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			v.keyValues = make(map[string]string)
			t.Errorf(err.Error())
		}
		_, err = DecryptReader(v.encodingKey, f)
		assert.Equal(t, nil, err)
		f.Close()
	})

	t.Run("it returns error when file is not opened in read/write mode", func(t *testing.T) {
		var v = Vault{encodingKey: "", filepath: "/home/sagar/.secrets"}
		f, err := os.OpenFile(v.filepath, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			v.keyValues = make(map[string]string)
			t.Errorf(err.Error())
		}
		_, err = DecryptReader(v.encodingKey, f)
		assert.NotEqual(t, nil, err)
		f.Close()
	})

	t.Run("it returns error when newCipherBlock does not work", func(t *testing.T) {
		var v = Vault{encodingKey: "", filepath: "/home/sagar/.secrets"}

		f, err := os.OpenFile(v.filepath, os.O_RDWR|os.O_CREATE, 0755)

		if err != nil {
			v.keyValues = make(map[string]string)
			t.Errorf(err.Error())
		}
		ff := &fakeSetup{err: errors.New("Failed")}
		MockNewCipherBlock = ff.setup
		_, err = DecryptReader(v.encodingKey, f)
		assert.NotEqual(t, nil, err)
		f.Close()
		defer func() {
			MockNewCipherBlock = newCipherBlock // set back original func at end of test
		}()
	})
}
