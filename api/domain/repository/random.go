package repository

import (
	crand "crypto/rand"
	"math"
	"math/big"
	"math/rand"

	"github.com/ttakuya50/go-architecture-sample/api/infra/random"
)

//go:generate mockgen -source=$GOFILE -destination=random_mock.go -package=$GOPACKAGE -self_package=github.com/ttakuya50/go-architecture-sample/api/domain/$GOPACKAGE

type RandomRepo interface {
	Int63() int64
}

type randomRepo struct {
}

func NewRandom() RandomRepo {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())
	return &randomRepo{}
}

func (r *randomRepo) Int63() int64 {
	return random.Int63()
}
