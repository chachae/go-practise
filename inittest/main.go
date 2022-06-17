package inittest

import (
	"fmt"
	"image"
	"os"
)

func main() {
	x, y, err := getImgSize(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("image bounds info x: %d y:%d\n", x, y)
	}
}

func getImgSize(filePath string) (int, int, error) {
	img, err := os.Open(filePath)
	defer func(img *os.File) {
		err := img.Close()
		if err != nil {
			return
		}
	}(img)
	decode, _, err := image.Decode(img)
	if err != nil {
		return 0, 0, err
	} else {
		bounds := decode.Bounds()
		return bounds.Dx(), bounds.Dy(), nil
	}
}
