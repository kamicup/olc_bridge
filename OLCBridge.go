package olc_bridge

import (
	"bytes"
	"encoding/json"
	olc "github.com/google/open-location-code/go"
)

/* Instruction:
 * $ gomobile bind [-target android|ios] [-o output] [build flags] [package]
 *
 * https://godoc.org/golang.org/x/mobile/cmd/gomobile
 */

func Encode(lat, lng float64, codeLen int) string {
	return olc.Encode(lat, lng, codeLen)
}

func Decode(code string) string {
	d, err := olc.Decode(code)
	if err == nil {
		buf := bytes.NewBufferString("")
		err = json.NewEncoder(buf).Encode(d)
		if err == nil {
			return buf.String()
		}
		return ""
	} else {
		return ""
	}
}
