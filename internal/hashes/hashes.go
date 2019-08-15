package hashes

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"hash"
)

func GetHash(kind, text string) (*string, error) {
	var (
		h      hash.Hash
		hashed string
	)
	switch kind {
	case "Sha1":
		h = sha1.New()
	case "Sha256":
		h = sha256.New()
	case "Sha384":
		h = sha512.New384()
	case "Sha512":
		h = sha512.New()
	case "Md5":
		h = md5.New()
	default:
		return nil, errors.New("unsupported hashing algorithm")
	}
	hashed = hex.EncodeToString(h.Sum([]byte(text)))
	return &hashed, nil
}
