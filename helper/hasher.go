package helper

import "golang.org/x/crypto/bcrypt"

type Hasher interface {
	Hash(string) string
}

type BcryptHasher struct {
	cost int
}

func NewBcryptHasher(cost int) *BcryptHasher {
	return &BcryptHasher{
		cost: cost,
	}
}

func (bh *BcryptHasher) Hash(s string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(s), bh.cost)
	return string(bytes)
}
