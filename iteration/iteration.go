package iteration

import "strings"

func Repeat(character string, numtimes int) string {
	return strings.Repeat(character, numtimes)

	// var repeated string
	// strings.Repeat(character, numtimes)
	// for i := 0; i < numtimes; i++ {
	// 	repeated += character
	// }
	// return repeated
}
