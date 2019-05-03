// +build js,wasm

package main

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/png"
	"log"
	"runtime"
	"time"

	"syscall/js"

	"github.com/disintegration/imaging"
	"github.com/vincent-petithory/dataurl"
)

var (
	canvas js.Value
	ctx    js.Value
)

func init() {
	canvas = js.Global().Get("document").Call("getElementById", canvasElementID)
	ctx = canvas.Call("getContext", "2d")

}

func logInfo(s interface{}) {
	log.Println(s)
	js.Global().Get("document").
		Call("getElementById", infoBoxElementID).
		Set("innerHTML", s)
}

func getModel() ([]byte, error) {
	files := js.Global().Get("document").Call("getElementById", knowledgeFileElementID).Get("files")
	if files.Length() == 1 {
		logInfo("Reading the model from the memory of the browser")
		//fmt.Println("Reading from uploaded file")
		reader := js.Global().Get("FileReader").New()
		reader.Call("readAsDataURL", files.Index(0))
		for reader.Get("readyState").Int() != 2 {
			//fmt.Println("Waiting for the file to be uploaded")
			time.Sleep(1 * time.Second)
		}
		content := reader.Get("result").String()
		dataURL, err := dataurl.DecodeString(content)
		return dataURL.Data, err
	}
	return nil, errors.New("too many file in the selector")
}

func getImage() (*image.Gray, error) {
	logInfo("Getting picture from the DOM")
	video := js.Global().Get("document").Call("getElementById", videoElementID)
	ctx.Call("drawImage", video, 0, 0)

	pic := canvas.Call("toDataURL")
	dataURL, err := dataurl.DecodeString(pic.String())
	if err != nil {
		return nil, err
	}
	if dataURL.ContentType() != "image/png" {
		return nil, errors.New("not a png image")
	}
	logInfo("Decoding the PNG file")
	m, err := png.Decode(bytes.NewReader(dataURL.Data))
	if err != nil {
		return nil, err
	}
	if m.Bounds().Dx() != m.Bounds().Dy() && m.Bounds().Dx() != width {
		// resize
		logInfo(fmt.Sprintf("image is %vx%v, resizing...", m.Bounds().Dx(), m.Bounds().Dy()))
		m = imaging.Resize(m, height, width, imaging.Lanczos)
	}

	var imgGray *image.Gray
	var ok bool
	imgGray, ok = m.(*image.Gray)
	if !ok {
		// convert to gray
		logInfo("convert picture to grayscale...")
		gray := imaging.Grayscale(m)
		imgGray = image.NewGray(gray.Bounds())
		for i := 0; i < len(imgGray.Pix); i++ {
			imgGray.Pix[i] = gray.Pix[i*4]
		}
	}
	return imgGray, nil
}

func displayResult(e emotions) {
	result := fmt.Sprintf("%v / %2.2f%%<br>%v / %2.2f%%",
		e[0].emotion, e[0].weight*100,
		e[1].emotion, e[1].weight*100,
	)
	logInfo(result)
}

func run() error {
	b, err := getModel()
	if err != nil {
		logInfo(err)
		return err
	}
	err = model.UnmarshalBinary(b)
	if err != nil {
		logInfo(err)
		return err
	}
	logInfo("Ready to process")
	// Declare callback
	cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// handle event
		// Get the picture
		img, err := getImage()
		runtime.GC()
		if err != nil {
			logInfo(err.Error())
			return nil

		}
		logInfo("displaying the result")

		err = displayPic(img)
		if err != nil {
			logInfo(err.Error())
		}
		logInfo("processing element")
		output, err := process(img)
		runtime.GC()
		if err != nil {
			logInfo(err.Error())
			return nil
		}

		displayResult(output)
		return nil
	})
	// Hook it up with a DOM event
	js.Global().Get("document").
		Call("getElementById", boutonSubmit).
		Call("addEventListener", "click", cb)
	c := make(chan struct{}, 0)
	<-c
	return nil
}

func displayPic(i *image.Gray) error {
	// encode in png
	var output bytes.Buffer
	err := png.Encode(&output, i)
	if err != nil {
		return err
	}

	processed := js.Global().Get("document").Call("createElement", "img")
	processed.Set("src", dataurl.EncodeBytes(output.Bytes()))

	ctx.Call("drawImage", processed, 0, 0)
	return nil
}
