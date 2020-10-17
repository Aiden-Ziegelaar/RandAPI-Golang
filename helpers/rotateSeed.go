package helpers

import (
	cRand "crypto/rand"
	"encoding/binary"
	mRand "math/rand"
)

func RotateSeed () {
	var seed int64
	_ = binary.Read(cRand.Reader, binary.BigEndian, &seed)
	mRand.Seed(seed)
}

