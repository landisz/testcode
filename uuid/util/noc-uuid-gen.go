package util

import (
	"encoding/hex"

	uuid "github.com/satori/go.uuid"
)

// UuidGen generates a uuid
func UuidGen() string {
	u := uuid.NewV4()

	ustring := String(u)
	return ustring
}

// String function converts UUID to string
func String(u uuid.UUID) string {
	buf := make([]byte, 36)

	hex.Encode(buf[0:8], u[0:4])
	buf[8] = '-'
	hex.Encode(buf[9:13], u[4:6])
	buf[13] = '-'
	hex.Encode(buf[14:18], u[6:8])
	buf[18] = '-'
	hex.Encode(buf[19:23], u[8:10])
	buf[23] = '-'
	hex.Encode(buf[24:], u[10:])

	return string(buf)
}