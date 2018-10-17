package main

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"strings"

	"syscall/js"

	"github.com/disintegration/gift"
)

const (
	jpegPrefix = "data:image/jpeg;base64,"
	pngPrefix  = "data:image/png;base64,"
)

var (
	TYPE = ""
)

func loadImageHandle(i []js.Value) {
	filtertype := i[0].String()
	str := i[1].String()
	transStr := trans(filtertype, str)
	js.Global().Get("document").Call("getElementById", "pic-target").Set("src", transStr)
}

func trans(filtertype, base64str string) string {

	src := base64toJpg(base64str)
	filters := map[string]gift.Filter{
		"resize":              gift.Resize(100, 0, gift.LanczosResampling),
		"crop_to_size":        gift.CropToSize(100, 100, gift.LeftAnchor),
		"rotate_180":          gift.Rotate180(),
		"rotate_30":           gift.Rotate(30, color.Transparent, gift.CubicInterpolation),
		"brightness_increase": gift.Brightness(30),
		"brightness_decrease": gift.Brightness(-30),
		"contrast_increase":   gift.Contrast(30),
		"contrast_decrease":   gift.Contrast(-30),
		"saturation_increase": gift.Saturation(50),
		"saturation_decrease": gift.Saturation(-50),
		"gamma_1.5":           gift.Gamma(1.5),
		"gamma_0.5":           gift.Gamma(0.5),
		"gaussian_blur":       gift.GaussianBlur(1),
		"unsharp_mask":        gift.UnsharpMask(1, 1, 0),
		"sigmoid":             gift.Sigmoid(0.5, 7),
		"pixelate":            gift.Pixelate(5),
		"colorize":            gift.Colorize(240, 50, 100),
		"grayscale":           gift.Grayscale(),
		"sepia":               gift.Sepia(100),
		"invert":              gift.Invert(),
		"mean":                gift.Mean(5, true),
		"median":              gift.Median(5, true),
		"minimum":             gift.Minimum(5, true),
		"maximum":             gift.Maximum(5, true),
		"hue_rotate":          gift.Hue(45),
		"color_balance":       gift.ColorBalance(10, -10, -10),
		"color_func": gift.ColorFunc(
			func(r0, g0, b0, a0 float32) (r, g, b, a float32) {
				r = 1 - r0   // invert the red channel
				g = g0 + 0.1 // shift the green channel by 0.1
				b = 0        // set the blue channel to 0
				a = a0       // preserve the alpha channel
				return r, g, b, a
			},
		),
		"convolution_emboss": gift.Convolution(
			[]float32{
				-1, -1, 0,
				-1, 1, 1,
				0, 1, 1,
			},
			false, false, false, 0.0,
		),
	}

	filter := filters[filtertype]
	g := gift.New(filter)
	dst := image.NewNRGBA(g.Bounds(src.Bounds()))
	g.Draw(dst, src)

	// In-memory buffer to store PNG image
	// before we base 64 encode it
	var buff bytes.Buffer

	// The Buffer satisfies the Writer interface so we can use it with Encode
	// In previous example we encoded to a file, this time to a temp buffer
	jpeg.Encode(&buff, dst, &jpeg.Options{88})

	// Encode the bytes in the buffer to a base64 string
	encodedString := base64.StdEncoding.EncodeToString(buff.Bytes())
	switch {
	case TYPE == "jpeg":
		encodedString = jpegPrefix + encodedString
	case TYPE == "png":
		encodedString = pngPrefix + encodedString
	default:
		encodedString = encodedString
	}

	return encodedString
}

//Given a base64 string of a JPEG, encodes it into an JPEG image test.jpg
func base64toJpg(base64str string) image.Image {

	switch {
	case strings.HasPrefix(base64str, jpegPrefix):
		TYPE = "jpeg"
		base64str = strings.TrimPrefix(base64str, jpegPrefix)
	case strings.HasPrefix(base64str, pngPrefix):
		TYPE = "png"
		base64str = strings.TrimPrefix(base64str, pngPrefix)
	default:
		println("unrecognized image format")
	}

	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(base64str))
	m, _, err := image.Decode(reader)

	if err != nil {
		log.Fatal(err)
	}

	return m
}

func registerCallbacks() {
	js.Global().Set("loadImageHandle", js.NewCallback(loadImageHandle))
}

func main() {
	c := make(chan struct{}, 0)
	println("WASM inited")

	registerCallbacks()
	<-c
}
