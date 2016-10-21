package persistent

import (
	"testing"

	"../../epaxos"
	"github.com/golang/leveldb/db"
	"github.com/stretchr/testify/assert"
)

func TestNewCloseAndDrop(t *testing.T) {
	l, err := NewLevelDB("/tmp/test", false)
	defer func() {
		assert.NoError(t, l.Close())
		assert.NoError(t, l.Drop())
	}()

	assert.NoError(t, err)
	assert.NotNil(t, l)
	assert.NotNil(t, l.ldb)
	assert.Equal(t, l.wsync, &db.WriteOptions{Sync: true})
}

func TestPutAndGet(t *testing.T) {
	l, err := NewLevelDB("/tmp/test", false)
	assert.NoError(t, err)

	defer func() {
		assert.NoError(t, l.Close())
		assert.NoError(t, l.Drop())
	}()

	assert.NoError(t, l.Put("hello", []byte("world")))

	v, err := l.Get("hello")
	assert.NoError(t, err)
	assert.Equal(t, v, []byte("world"))

	v, err = l.Get("world")
	assert.Equal(t, err, epaxos.ErrorNotFound)
}

func TestBatchPut(t *testing.T) {
	l, err := NewLevelDB("/tmp/test", false)
	assert.NoError(t, err)

	defer func() {
		assert.NoError(t, l.Close())
		assert.NoError(t, l.Drop())
	}()

	kvs := make([]*epaxos.KVpair, 2)
	kvs[0] = &epaxos.KVpair{
		Key:   "hello",
		Value: []byte("world"),
	}
	kvs[1] = &epaxos.KVpair{
		Key:   "epaxos",
		Value: []byte("rocks"),
	}

	assert.NoError(t, l.BatchPut(kvs))
	v, err := l.Get("hello")
	assert.NoError(t, err)
	assert.Equal(t, v, []byte("world"))

	v, err = l.Get("epaxos")
	assert.Equal(t, v, []byte("rocks"))
}
