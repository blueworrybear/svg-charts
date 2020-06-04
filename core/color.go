package core

import (
	"errors"
	"strconv"
)

type RGB struct {
	R int
	G int
	B int
}

type Hex string

var errHexFormat = errors.New("Error Hex format")

func Hex2RGB(hex Hex) (RGB, error) {
	chars := []rune(hex)
	if string(chars[0]) == "#" {
		hex = Hex(chars[1:])
	}
	if len(hex) != 6 {
		return RGB{}, errHexFormat
	}
	values, err := strconv.ParseUint(string(hex), 16, 32)
	if err != nil {
		return RGB{}, err
	}
	rgb := RGB{
		R: int(values >> 16),
		G: int((values >> 8) & 0xFF),
		B: int(values & 0xFF),
	}
	return rgb, nil
}
