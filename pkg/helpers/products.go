package helpers

import (
	"image"
	"image/jpeg"
	"math/rand"
	"net/http"
	"os"
)

func GenrateRandomString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func DownloadImage(url string) (image.Image, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	img, _, err := image.Decode(response.Body)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func SaveAsJPEG(img image.Image, outputDir string, quality int) (*string, error) {
	directoryPath := outputDir

	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		err := os.Mkdir(directoryPath, 0755)
		if err != nil {
			return nil, err
		}
	}
	path := directoryPath + GenrateRandomString(15) + ".jpg"
	outputFile, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	defer outputFile.Close()

	jpegOptions := jpeg.Options{Quality: quality}
	err = jpeg.Encode(outputFile, img, &jpegOptions)
	if err != nil {
		return nil, err
	}
	return &path, nil
}
