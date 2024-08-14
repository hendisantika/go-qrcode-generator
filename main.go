package main

import (
	"encoding/json"
	"fmt"
	"github.com/skip2/go-qrcode"
	"net/http"
	"strconv"
)

const MAX_UPLOAD_SIZE = 1024 * 1024 // 1MB
const WATERMARK_WIDTH = 64

func handleRequest(writer http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(10 << 20)
	var size, content = request.FormValue("size"), request.FormValue("content")
	var codeData []byte

	writer.Header().Set("Content-Type", "application/json")

	if content == "" {
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(
			"Could not determine the desired QR code content.",
		)
		return
	}

	qrCodeSize, err := strconv.Atoi(size)
	if err != nil || size == "" {
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode("Could not determine the desired QR code size.")
		return
	}

	qrCode := simpleQRCode{Content: content, Size: qrCodeSize}
	codeData, err = qrCode.Generate()
	if err != nil {
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(
			fmt.Sprintf("Could not generate QR code. %v", err),
		)
		return
	}

	writer.Header().Set("Content-Type", "image/png")
	writer.Write(codeData)
}

func main() {
	http.HandleFunc("/generate", handleRequest)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Listening on port 8080")
}

type simpleQRCode struct {
	Content string
	Size    int
}

func (code *simpleQRCode) Generate() ([]byte, error) {
	qrCode, err := qrcode.Encode(code.Content, qrcode.Medium, code.Size)
	if err != nil {
		return nil, fmt.Errorf("could not generate a QR code: %v", err)
	}
	return qrCode, nil
}

// GenerateWithWatermark generates a QR code using the value of simpleQRCode.Content
// and adds a watermark to it, centered in the middle of the QR code, using the
// supplied watermark image data
func (code *simpleQRCode) GenerateWithWatermark(watermark []byte) ([]byte, error) {
	qrCode, err := code.Generate()
	if err != nil {
		return nil, err
	}

	qrCode, err = code.addWatermark(qrCode, watermark, code.Size)
	if err != nil {
		return nil, fmt.Errorf("could not add watermark to QR code: %v", err)
	}

	return qrCode, nil
}
