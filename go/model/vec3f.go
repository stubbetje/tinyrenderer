package model

import "math"

type Vec3f struct {
	X, Y, Z float64;
}

func NewVec3f(x, y, z float64) Vec3f {
	return Vec3f{X:x, Y:y, Z:z}
}

func ( v1 Vec3f ) Subtract(v2 Vec3f) Vec3f {
	return Vec3f{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
		Z: v1.Z - v2.Z,
	}
}

func ( v1 Vec3f ) Product(v2 Vec3f) Vec3f {
	return Vec3f{
		X: v1.Y * v2.Z - v1.Z * v2.Y,
		Y: v1.Z * v2.X - v1.X * v2.Z,
		Z: v1.X * v2.Y - v1.Y * v2.X,
	}
}

func ( v Vec3f ) Scale(f float64) Vec3f {
	return Vec3f{
		X: v.X * f,
		Y: v.Y * f,
		Z: v.Z * f,
	}
}

func ( v1 Vec3f ) Times(v2 Vec3f) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z * v2.Z
}

func ( v1 Vec3f ) norm() float64 {
	return math.Sqrt(v1.X * v1.X + v1.Y * v1.Y + v1.Z * v1.Z)
}

func ( v Vec3f ) Normalize(l int) Vec3f {
	return v.Scale(float64(l) / v.norm())
}

//inline Vec3<t> operator ^(const Vec3<t> &v) const { return Vec3<t>(y*v.z-z*v.y, z*v.x-x*v.z, x*v.y-y*v.x); }
//inline Vec3<t> operator +(const Vec3<t> &v) const { return Vec3<t>(x+v.x, y+v.y, z+v.z); }
//inline Vec3<t> operator -(const Vec3<t> &v) const { return Vec3<t>(x-v.x, y-v.y, z-v.z); }
//inline Vec3<t> operator *(float f)          const { return Vec3<t>(x*f, y*f, z*f); }
//inline t       operator *(const Vec3<t> &v) const { return x*v.x + y*v.y + z*v.z; }
//float norm () const { return std::sqrt(x*x+y*y+z*z); }
//Vec3<t> & normalize(t l=1) { *this = (*this)*(l/norm()); return *this; }
