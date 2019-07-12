package controllers

import (
	"context"
	"fmt"
	"github.com/hitman99/hashing/gen/hashing"
	"github.com/hitman99/hashing/internal/hashes"
	"log"
	"regexp"
)

// hashing service example implementation.
// The example methods log the requests and return zero values.
type hashingsrvc struct {
	logger *log.Logger
	hashTypeRegex *regexp.Regexp
}

// NewHashing returns the hashing service implementation.
func NewHashing(logger *log.Logger) hashing.Service {
	return &hashingsrvc{
		logger: logger,
		hashTypeRegex: regexp.MustCompile("^\\*hashing\\.([a-zA-Z0-9]+)Payload$"),
	}
}

func (s *hashingsrvc) Sha1(ctx context.Context, p *hashing.Sha1Payload) (res *hashing.Hashingresult, err error) {
	s.logger.Print("hashing.sha1")
	kind := s.getHashType(p)
	hashed, err := hashes.GetHash(*kind, p.String)
	res = &hashing.Hashingresult{
		Type: kind,
		Hash: hashed,
	}
	return
}

func (s *hashingsrvc) Sha256(ctx context.Context, p *hashing.Sha256Payload) (res *hashing.Hashingresult, err error) {
	s.logger.Print("hashing.sha256")
	kind := s.getHashType(p)
	hashed, err := hashes.GetHash(*kind, p.String)
	res = &hashing.Hashingresult{
		Type: kind,
		Hash: hashed,
	}
	return
}

func (s *hashingsrvc) Sha384(ctx context.Context, p *hashing.Sha384Payload) (res *hashing.Hashingresult, err error) {
	s.logger.Print("hashing.sha384")
	kind := s.getHashType(p)
	hashed, err := hashes.GetHash(*kind, p.String)
	res = &hashing.Hashingresult{
		Type: kind,
		Hash: hashed,
	}
	return
}

func (s *hashingsrvc) Sha512(ctx context.Context, p *hashing.Sha512Payload) (res *hashing.Hashingresult, err error) {
	s.logger.Print("hashing.sha512")
	kind := s.getHashType(p)
	hashed, err := hashes.GetHash(*kind, p.String)
	res = &hashing.Hashingresult{
		Type: kind,
		Hash: hashed,
	}
	return
}

func (s *hashingsrvc) Md5(ctx context.Context, p *hashing.Md5Payload) (res *hashing.Hashingresult, err error) {
	s.logger.Print("hashing.md5")
	kind := s.getHashType(p)
	hashed, err := hashes.GetHash(*kind, p.String)
	res = &hashing.Hashingresult{
		Type: kind,
		Hash: hashed,
	}
	return
}

func (s *hashingsrvc) getHashType( p interface{}) *string {
	t := s.hashTypeRegex.FindStringSubmatch(fmt.Sprintf("%T", p))
	hashType := "unknown"
	if len(t) == 2 {
		hashType = t[1]
	}
	return &hashType
}
