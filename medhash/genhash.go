// MedHash Tools
// Copyright (c) 2023 GHIFARI160
// MIT License

package medhash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"hash"
	"io"
	"os"
	"path/filepath"

	"golang.org/x/crypto/sha3"
)

// Config configures the hasher.
type Config struct {
	// Dir is the path to the target directory.
	Dir string
	// Path is the path of the current media.
	Path string

	// SHA3 toggles the SHA3-256 hash generation.
	SHA3 bool
	// SHA256 toggles the SHA256 hash generation.
	SHA256 bool
	// SHA1 toggles the SHA1 hash generation.
	// Deprecated: SHA1 support is deprecated in spec v0.4.0.
	SHA1 bool
	// MD5 toggles the MD5 hash generation.
	// Deprecated: MD5 support is deprecated in spec v0.4.0.
	MD5 bool
}

// GenHash generates a hash for the media specified in the config path.
// Hashes for the media are generated at the same time.
func GenHash(config Config) (med Media, err error) {
	return genHash(config)
}

// genHash generates a hash for the media specified in the config path.
func genHash(config Config) (med Media, err error) {
	writers := make([]io.Writer, 0)
	hashers := make(map[string]hash.Hash)

	if config.SHA3 {
		hashers["sha3"] = sha3.New256()
		writers = append(writers, hashers["sha3"])
	}

	if config.SHA256 {
		hashers["sha256"] = sha256.New()
		writers = append(writers, hashers["sha256"])
	}

	if config.SHA1 {
		hashers["sha1"] = sha1.New()
		writers = append(writers, hashers["sha1"])
	}

	if config.MD5 {
		hashers["md5"] = md5.New()
		writers = append(writers, hashers["md5"])
	}

	writer := io.MultiWriter(writers...)

	f, err := os.Open(filepath.Join(config.Dir, config.Path))
	if err != nil {
		return
	}
	defer f.Close()

	_, err = io.Copy(writer, f)
	if err != nil {
		return
	}

	hash := Hash{}

	if h := hashers["sha3"]; h != nil {
		hash.SHA3_256 = hex.EncodeToString(h.Sum(nil))
	}

	if h := hashers["sha256"]; h != nil {
		hash.SHA256 = hex.EncodeToString(h.Sum(nil))
	}

	if h := hashers["sha1"]; h != nil {
		hash.SHA1 = hex.EncodeToString(h.Sum(nil))
	}

	if h := hashers["md5"]; h != nil {
		hash.MD5 = hex.EncodeToString(h.Sum(nil))
	}

	med.Path = filepath.ToSlash(config.Path)
	med.Hash = hash

	return
}
