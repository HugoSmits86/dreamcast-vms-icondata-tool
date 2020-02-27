package icondata

import (
	"io"
	"encoding/binary"
	"image"
	"image/color"
    "errors"
)

func Decode(r io.Reader) (string, image.Image, error) {
	if r == nil {
		return "", nil, errors.New("reader is nil")
	}

	buf := make([]byte, 1024)
	_, err := io.ReadFull(r, buf[:24])
	if err != nil {
		return "", nil, errors.New("could not read header")
	}
	desc := string(buf[:16])
	monoOff := int(binary.LittleEndian.Uint32(buf[16:20]))

	if monoOff < 24 || monoOff - 24 >= 1024 {
		return "", nil, errors.New("mono offset out of bounds")
	}

	_, err = io.ReadFull(r, buf[:(monoOff - 24)])
	if err != nil {
		return "", nil, errors.New("mono offset out of bounds")
	}

	_, err = io.ReadFull(r, buf[:128])
	if err != nil {
		return "", nil, errors.New("could not read mono bitmap")
	}

	monoImage := image.NewRGBA(image.Rect(0, 0, 32, 32))
	for i := 0; i < 128; i++ {
		b := buf[i]
		for j := 0; j < 8; j++ {
			bit := (b & (1 << j)) >> j
			c := uint8(0xff)
			if bit == 1 {
				c = 0x00
			}
			x := i % 4 * 8 + (7 - j)
			y := i / 4
			monoImage.Set(x, y, color.RGBA{c, c, c, 0xff})
		}
	}

	return desc, monoImage, nil
}

func Encode(w io.Writer, img image.Image) error {
	if img == nil {
		return errors.New("image is nil")
	}

	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	if width < 32 || height < 32 {
		return errors.New("image too small")
	}

	dest := make([]byte, 152)
	copy(dest[0:16], []byte("made with govmu!"))

	off := uint32(24)
	tmp := make([]byte, 4)
	binary.LittleEndian.PutUint32(tmp, off)
	copy(dest[16:20], tmp)

	binary.LittleEndian.PutUint32(tmp, 0)
	copy(dest[20:24], tmp)

	i := off
	for y := 0; y < 32; y++ {
		for x := 0; x < 32; x+= 8 {
			b := uint8(0)
			for j := 0; j < 8; j++ {
				r, _, _, _ := img.At(x + j, y).RGBA()
				if r >> 8 < 128 {
					b |= uint8(1 << uint(7 - j))
				}
			}
			dest[i] = b
			i++
		}
	}

	w.Write(dest)

	return nil
}