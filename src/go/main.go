package main

import (
	"bytes"
	"encoding/base64"
	"image"
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
	str := i[0].String()
	transStr := trans(str)
	js.Global().Get("document").Call("getElementById", "pic-target").Set("src", transStr)
}

func trans(base64str string) string {
	src := base64toJpg(base64str)

	filter := gift.Rotate180()
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
