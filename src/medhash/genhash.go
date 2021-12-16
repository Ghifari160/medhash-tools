// MedHash Tools
// Copyright (c) 2021 GHIFARI160
// MIT License

package medhash

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"hash"
	"io"
	"os"
)

func bufferedGenHash(path string, hasher *hash.Hash) error {
	f, err := os.Open(path)
	if err != nil {
		return fmtError(err)
	}

	file := bufio.NewReader(f)

	_, err = io.Copy((*hasher), file)
	if err != nil {
		return fmtError(err)
	}

	return f.Close()
}

func GenHash(path string) (*Hash, error) {
	hash := Hash{}
	hasher := sha256.New()

	err := bufferedGenHash(path, &hasher)
	if err != nil {
		return nil, fmtError(err)
	}
	hash.SHA256 = hex.EncodeToString(hasher.Sum(nil))

	hasher = sha1.New()
	err = bufferedGenHash(path, &hasher)
	if err != nil {
		return nil, fmtError(err)
	}
	hash.SHA1 = hex.EncodeToString(hasher.Sum(nil))

	hasher = md5.New()
	err = bufferedGenHash(path, &hasher)
	if err != nil {
		return nil, fmtError(err)
	}
	hash.MD5 = hex.EncodeToString(hasher.Sum(nil))

	return &hash, nil
}