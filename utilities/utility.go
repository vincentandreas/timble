package utilities

import (
	"crypto/sha256"
	"encoding/base64"
	"log"
	"time"
)

func HashParams(input string) string {
	h := sha256.New()
	h.Write([]byte(input))
	sha := base64.URLEncoding.EncodeToString(h.Sum(nil))
	return sha
}

func DateEqual(date1, date2 time.Time) bool {
    y1, m1, d1 := date1.Date()
    y2, m2, d2 := date2.Date()
    return y1 == y2 && m1 == m2 && d1 == d2
}

func DateBefore(date1, date2 time.Time) bool {
	date1 = time.Date(date1.Year(), date1.Month(), date1.Day(), 0, 0, 0, 0, date1.Location())
	date2 = time.Date(date2.Year(), date2.Month(), date2.Day(), 0, 0, 0, 0, date2.Location())

	return date1.Before(date2)
}


func ExtractId(data any) uint64 {
    var userIdUint64 uint64

    switch v := data.(type) {
    case int:
        userIdUint64 = uint64(v)
    case int64:
        userIdUint64 = uint64(v)
    case uint:
        userIdUint64 = uint64(v)
    case uint64:
        userIdUint64 = v
    default:
        userIdUint64 = 0
    }
	return userIdUint64
}

func CheckError(err error){
    if err != nil {
        log.Fatal(err.Error())
    }
}