package hashid

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

import (
	"fmt"

	hashids "github.com/speps/go-hashids"

	appLog "pearshop_backend/pkg/log"
)

var singleton IDHasher

type IDHasher interface {
	Encode(id int) string
	Decode(hashed string) (int, error)
}

type idHasher struct {
	hashID *hashids.HashID
}

// InitIDHasher initials an ID hasher
func InitIDHasher(minLength int, salt string) error {
	if singleton != nil {
		return nil
	}

	hashID, err := hashids.NewWithData(&hashids.HashIDData{
		Alphabet:  hashids.DefaultAlphabet,
		MinLength: minLength,
		Salt:      salt,
	})
	if err != nil {
		return fmt.Errorf("error while init hash ID: %w", err)
	}

	singleton = &idHasher{
		hashID: hashID,
	}

	return nil
}

// GetIDHasher Get returns the singleton instance of ID hasher
func GetIDHasher() IDHasher {
	return singleton
}

// Encode encodes an integer based ID to a string
func (h *idHasher) Encode(id int) string {
	hashed, err := h.hashID.Encode([]int{id})
	if err != nil {
		appLog.Errorf("error while encoding id: %s", err.Error())
	}

	return hashed
}

// Decode decodes the hashed value to an integer
func (h *idHasher) Decode(hashed string) (int, error) {
	id, err := h.hashID.DecodeWithError(hashed)
	if err != nil {
		return 0, fmt.Errorf("error while decoding hashed id: %w", err)
	}

	if len(id) != 1 {
		return 0, fmt.Errorf("incorrect hash id format")
	}

	return id[0], nil
}
