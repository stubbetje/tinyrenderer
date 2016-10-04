package model

type Vec3f struct {
	X, Y, Z float64;
}

func NewVec3f(x, y, z float64) Vec3f {
	return Vec3f{X:x, Y:y, Z:z}
}
