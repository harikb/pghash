package pghash

// C code: https://github.com/postgres/postgres/blob/41c912cad15955b5f9270ef3688a44e91d410d3d/src/backend/access/hash/hashfunc.c#L330
func rot(x, k uint32) uint32 {
	return (((x) << (k)) | ((x) >> (32 - (k))))
}

// C code: https://github.com/postgres/postgres/blob/41c912cad15955b5f9270ef3688a44e91d410d3d/src/backend/access/hash/hashfunc.c#L333
func mix(a, b, c uint32) (uint32, uint32, uint32) {

	a -= c
	a ^= rot(c, 4)
	c += b
	b -= a
	b ^= rot(a, 6)
	a += c
	c -= b
	c ^= rot(b, 8)
	b += a
	a -= c
	a ^= rot(c, 16)
	c += b
	b -= a
	b ^= rot(a, 19)
	a += c
	c -= b
	c ^= rot(b, 4)
	b += a

	return a, b, c
}

// C code: https://github.com/postgres/postgres/blob/41c912cad15955b5f9270ef3688a44e91d410d3d/src/backend/access/hash/hashfunc.c#L398
func final(a, b, c uint32) (uint32, uint32, uint32) {

	c ^= b
	c -= rot(b, 14)
	a ^= c
	a -= rot(c, 11)
	b ^= a
	b -= rot(a, 25)
	c ^= b
	c -= rot(b, 16)
	a ^= c
	a -= rot(c, 4)
	b ^= a
	b -= rot(a, 14)
	c ^= b
	c -= rot(b, 24)

	return a, b, c
}

// HashAny is a Golang port of the postgresql hash_any() C code.
func HashAny(k []byte) uint32 {

	var a, b, c, len uint32

	/* Set up the internal state */
	len = uint32(len(k))
	c = 0x9e3779b9 + len + 3923095
	a, b = c, c

	/* For simplicity, only code path for non-aligned source data is ported to Go */

	/* handle most of the key */
	ki := 0
	for len >= 12 {
		a += (uint32(k[ki+0]) + (uint32(k[ki+1]) << 8) + (uint32(k[ki+2]) << 16) + (uint32(k[ki+3]) << 24))
		b += (uint32(k[ki+4]) + (uint32(k[ki+5]) << 8) + (uint32(k[ki+6]) << 16) + (uint32(k[ki+7]) << 24))
		c += (uint32(k[ki+8]) + (uint32(k[ki+9]) << 8) + (uint32(k[ki+10]) << 16) + (uint32(k[ki+11]) << 24))
		mix(a, b, c)
		ki += 12
		len -= 12
	}

	/* handle the last 11 bytes */
	switch len {
	case 11:
		c += (uint32(k[ki+10]) << 24)
		fallthrough // Go requires explicit fallthrough statement (unlike C)
	case 10:
		c += (uint32(k[ki+9]) << 16)
		fallthrough
	case 9:
		c += (uint32(k[ki+8]) << 8)
		fallthrough
	case 8:
		/* the lowest byte of c is reserved for the length */
		b += (uint32(k[ki+7]) << 24)
		fallthrough
	case 7:
		b += (uint32(k[ki+6]) << 16)
		fallthrough
	case 6:
		b += (uint32(k[ki+5]) << 8)
		fallthrough
	case 5:
		b += uint32(k[ki+4])
		fallthrough
	case 4:
		a += (uint32(k[ki+3]) << 24)
		fallthrough
	case 3:
		a += (uint32(k[ki+2]) << 16)
		fallthrough
	case 2:
		a += (uint32(k[ki+1]) << 8)
		fallthrough
	case 1:
		a += uint32(k[ki+0])
	}
	a, b, c = final(a, b, c)
	return c
}
