package model

type Vec2i struct {
	X, Y int;
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
