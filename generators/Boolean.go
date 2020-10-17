package generators

import (
	mRand "math/rand"
	"randAPI/proto/grpcService"
)

func GenerateBoolean() grpcService.BooleanResponse {
	return grpcService.BooleanResponse{Result: mRand.Intn(2) > 0}
}

func GenerateWeightedBoolean(falseWeight int, trueWeight int) grpcService.BooleanResponse {
	return grpcService.BooleanResponse{Result: mRand.Intn(falseWeight+trueWeight) > falseWeight-1}
}
