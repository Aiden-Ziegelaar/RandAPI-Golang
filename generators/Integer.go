package generators

import (
	"math/rand"
	"randAPI/proto/grpcService"
)

func GenerateIntegerN(max int32) grpcService.IntegerResponse {
	return grpcService.IntegerResponse{Result: int32(rand.Int31n(max+1))}
}

func GenerateInteger() grpcService.IntegerResponse {
	return grpcService.IntegerResponse{Result: int32(rand.Int31())}
}
