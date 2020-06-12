package random

import (
	"context"
	"crypto/md5"
	"encoding/binary"
	"math/rand"
)

type key int

func New(seed string) *rand.Rand {
	h := md5.New()
	_, _ = h.Write([]byte(seed))
	return rand.New(rand.NewSource(int64(binary.BigEndian.Uint64(h.Sum(nil)))))
}

func WithSeed(ctx context.Context, seed string) context.Context {
	return context.WithValue(ctx, key(0), New(seed))
}

func FromContext(ctx context.Context) *rand.Rand {
	v := ctx.Value(key(0))
	if r, ok := v.(*rand.Rand); ok {
		return r
	}
	return nil
}

func Intn(ctx context.Context, n int) int {
	intn := rand.Intn
	if r := FromContext(ctx); r != nil {
		intn = r.Intn
	}
	return intn(n)
}

func Int63n(ctx context.Context, n int) int64 {
	intn := rand.Int63n
	if r := FromContext(ctx); r != nil {
		intn = r.Int63n
	}
	return intn(int64(n))
}

func Perm(ctx context.Context, n int) []int {
	perm := rand.Perm
	if r := FromContext(ctx); r != nil {
		perm = r.Perm
	}
	return perm(n)
}
