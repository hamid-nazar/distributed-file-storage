package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func ContentAddressablePathTransformFunc(key string) string {
	hash := sha1.Sum([]byte(key))
	hashString := hex.EncodeToString(hash[:])

	blockSize := 5
	sliceLen := len(hashString) / blockSize
	paths := make([]string, sliceLen)

	for i := 0; i < sliceLen; i++ {
		from, to := i*blockSize, (i*blockSize)+blockSize
		paths[i] = hashString[from:to]
	}

	return strings.Join(paths, "/")
}

type PathTransformFunc func(string) string

var DefaultPathTransform = func(key string) string {
	return key
}

type StoreOptions struct {
	PathTransformFunc PathTransformFunc
}

type Store struct {
	StoreOptions
}

func NewStore(options StoreOptions) *Store {
	return &Store{
		StoreOptions: options,
	}
}

func (store *Store) writeStream(key string, r io.Reader) error {
	pathName := store.PathTransformFunc(key)

	if err := os.MkdirAll(pathName, os.ModePerm); err != nil {
		return err
	}

	fileName := "randdomFileName"

	pathAndFileName := pathName + "/" + fileName

	file, err := os.Create(pathAndFileName)

	if err != nil {
		return err
	}

	n, err := io.Copy(file, r)

	if err != nil {
		return err
	}

	log.Printf("Written (%d) bytes to disk: %s\n", n, pathAndFileName)

	return nil
}
