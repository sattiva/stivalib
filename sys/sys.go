package sys

import (
	"crypto/rand"
	"io"
	"math/big"
	"os"
	"os/signal"
	"syscall"
)

func SecureRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := io.ReadFull(rand.Reader, b)
	return b, err
}

func SecureRandomInt(max int64) (int64, error) {
	nBig, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		return 0, err
	}
	return nBig.Int64(), nil
}

func ZeroBuffer(b []byte) {
	for i := range b {
		b[i] = 0
	}
}

func WaitForShutdown(cleanup func()) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig
	if cleanup != nil {
		cleanup()
	}
}
