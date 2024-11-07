package rand

import (
	"fmt"
	"github.com/gofrs/uuid"
	"math/rand"
	"strconv"
	"time"
)

const alphaCharset = "abcdefghijklmnopqrstuvwyxz"
const emailCharset = "abcdefghijklmnopqrstuvwxyz0123456789"
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const base64Charset = "abcdefghijklmnopqrstuvwxyz0123456789"
const base64Int = "0123456789"
const base64Float = "0123456789."

const unambiguousCharset = "abcdefghjkmnprstvwxy2346789"

var seededRand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func StringWithIntSet(length int, intSet string) int {
	b := make([]byte, length)
	for i := range b {
		b[i] = intSet[seededRand.Intn(len(intSet))]
	}

	// Convert the byte slice to a string
	str := string(b)
	result, _ := strconv.Atoi(str)

	return result
}

func StringWithFloatSet(length int, floatSet string) float64 {
	b := make([]byte, length)
	for i := range b {
		b[i] = floatSet[seededRand.Intn(len(floatSet))]
	}

	// Convert the byte slice to a string
	str := string(b)
	result, _ := strconv.ParseFloat(str, 64)

	return result
}

func ApiKey() string {
	//4-4-4-4
	newUuid := fmt.Sprintf(
		"%v-%v-%v-%v",
		StringWithCharset(4, unambiguousCharset),
		StringWithCharset(4, unambiguousCharset),
		StringWithCharset(4, unambiguousCharset),
		StringWithCharset(4, unambiguousCharset),
	)
	return newUuid
}

func Uuid() string {
	//8-4-4-4-12
	newUuid := fmt.Sprintf(
		"%v-%v-%v-%v-%v",
		StringWithCharset(8, emailCharset),
		StringWithCharset(4, emailCharset),
		StringWithCharset(4, emailCharset),
		StringWithCharset(4, emailCharset),
		StringWithCharset(12, emailCharset),
	)
	return newUuid
}

func Email() string {
	name := StringWithCharset(10, emailCharset)
	return fmt.Sprintf("%v@purplevisits.com", name)
}

func StringFixed(length int) string {
	return StringWithCharset(length, charset)
}

func StringBase64Fixed(length int) string {
	return StringWithCharset(length, base64Charset)
}

func String() string {
	return StringWithCharset(10, charset)
}

func Int() int {
	return StringWithIntSet(3, base64Int)
}

func Float() float64 {
	return StringWithFloatSet(1, base64Float)
}

func StringAlpha() string {
	return StringWithCharset(10, alphaCharset)
}

func UidV4() string {
	newUuid, _ := uuid.NewV4()
	return newUuid.String()
}
