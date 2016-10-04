package model

type Vec2i struct {
	X, Y int;
}

func ( v1 Vec2i ) Add(v2 Vec2i) Vec2i {
	return Vec2i{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
	}
}

func ( v1 Vec2i ) Subtract(v2 Vec2i) Vec2i {
	return Vec2i{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
	}
}

func ( v1 Vec2i ) Scale(f float64) Vec2i {
	return Vec2i{
		X: int(float64(v1.X) * f),
		Y: int(float64(v1.Y) * f),
	}
}


//template <class t> struct Vec2 {
//union {
//struct {t u, v;};
//struct {t x, y;};
//t raw[2];
//};
//Vec2() : u(0), v(0) {}
//Vec2(t _u, t _v) : u(_u),v(_v) {}
//inline Vec2<t> operator +(const Vec2<t> &V) const { return Vec2<t>(u+V.u, v+V.v); }
//inline Vec2<t> operator -(const Vec2<t> &V) const { return Vec2<t>(u-V.u, v-V.v); }
//inline Vec2<t> operator *(float f)          const { return Vec2<t>(u*f, v*f); }
//template <class > friend std::ostream& operator<<(std::ostream& s, Vec2<t>& v);
//};
