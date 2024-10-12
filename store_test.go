package main

import (
	"bytes"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "some random key"
	pathName := ContentAddressablePathTransformFunc(key)

	expectedValue := "ab30c/0f343/8434e/90956/b329b/642d6/757df/69695"

	if pathName != expectedValue {
		t.Errorf("got %s, want %s", pathName, "ab30c/0f343/8434e/90956/b329b/642d6/757df/69695")
	}
}

func TestStore(t *testing.T) {
	option := StoreOptions{
		PathTransformFunc: DefaultPathTransform,
	}

	store := NewStore(option)
	data := bytes.NewReader([]byte("some random data"))
	if err := store.writeStream("data", data); err != nil {
		t.Error(err)
	}
}
