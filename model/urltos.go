package model

import (
	"bytes"
	"image/png"

	"github.com/qpliu/qrencode-go/qrencode"
)

//CreateQR ddd
func CreateQR(url string) *bytes.Buffer {
	grid, _ := qrencode.Encode(url, qrencode.ECLevelQ)
	buf := new(bytes.Buffer)
	png.Encode(buf, grid.Image(8))
	return buf
}
