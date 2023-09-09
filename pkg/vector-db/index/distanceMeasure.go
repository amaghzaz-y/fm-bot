package index

import "math"

type DistanceMeasure interface {
	CalcDistance(v1, v2 []float32) float32
}

type cosineDistanceMeasure struct{}

func NewCosineDistanceMeasure() DistanceMeasure {
	return &cosineDistanceMeasure{}
}

func (cdm *cosineDistanceMeasure) CalcDistance(v1, v2 []float32) float32 {
	// calculates the cosine distance between two vectors
	if len(v1) != len(v2) || len(v1) == 0 {
		return 0.0
	}

	var dotProduct float32 = 0.0
	var magA float32 = 0.0
	var magB float32 = 0.0

	for i := 0; i < len(v1); i++ {
		dotProduct += v1[i] * v2[i]
		magA += v1[i] * v1[i]
		magB += v2[i] * v2[i]
	}

	magA = float32(math.Sqrt(float64(magA)))
	magB = float32(math.Sqrt(float64(magB)))

	if magA == 0 || magB == 0 {
		return 0.0
	}

	return -dotProduct / (magA * magB)
}

type euclideanDistanceMeasure struct{}

func NewEuclideanDistanceMeasure() DistanceMeasure {
	return &euclideanDistanceMeasure{}
}

func (cdm *euclideanDistanceMeasure) CalcDistance(v1, v2 []float32) float32 {
	// calculates the euclidean distance between two vectors
	if len(v1) != len(v2) || len(v1) == 0 {
		return 0.0
	}

	var sum float32 = 0.0

	for i := 0; i < len(v1); i++ {
		diff := v1[i] - v2[i]
		sum += float32(diff * diff)
	}

	return float32(math.Sqrt(float64(sum)))
}
