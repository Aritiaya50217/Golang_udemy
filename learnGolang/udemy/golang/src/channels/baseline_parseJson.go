package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type Photos []struct {
	AlbumId      int    `json:"albumId"`
	ID           int    `json:"id"`
	Title        string `json:"title"`
	URL          string `json:"url""`
	ThumbnailURL string `json:"thumbnailUrl"`
}

type Image struct {
	filePath string
	img      []byte
}

func getJson(url string, structType interface{}) error {
	//	https://jsonplaceholder.typicode.com/photos
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	switch v := structType.(type) {
	case *Photos:
		// รับค่า
		decoder := json.NewDecoder(res.Body)

		photos := structType.(*Photos)
		decoder.Decode(photos)
		return nil
	default:
		return fmt.Errorf("getJson : not support type %v ", v)
	}
}

func saveImage(fileName string, img []byte) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("saveImage - cannot create file : %v", err)
	}
	defer file.Close()

	// copy image to file
	_, err = io.Copy(file, bytes.NewReader(img))
	if err != nil {
		return fmt.Errorf("saveImage - cannot save file : %v", err)
	}
	log.Printf("Saved : %v\n", fileName)
	return err
}

func decodeImage(img []byte) (string, error) {
	_, format, err := image.Decode(bytes.NewReader(img))
	return format, err
}

func downloadImage(url string) ([]byte, error) {
	errMsg := func(err error) error {
		return fmt.Errorf("downloadImage : %v", err)
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, errMsg(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errMsg(err)
	}
	return body, nil
}

func main() {

	defer func() {
		fmt.Println("Main program exit successfully .")
	}()

	log.SetFlags(log.Ltime)
	photos := Photos{}
	err := getJson("https://jsonplaceholder.typicode.com/photos", &photos)
	fmt.Println(err)
	//fmt.Println("struct : ", photos[0:3])
	fmt.Println(len(photos))

	// check directory
	dir := "MyDownloadImage_" + time.Now().Format("02_01_2006")
	if _, err = os.Stat(dir); err != nil {
		os.Mkdir(dir, os.ModeDir)
	}
	chImg := make(chan Image, len(photos))
	for _, v := range photos {
		v := v
		go func() {
			// download image
			img, err := downloadImage(v.ThumbnailURL)
			if err != nil {
				log.Fatal(err)
			}

			// saveImage
			format, err := decodeImage(img)
			if err != nil {
				log.Fatal(err)
			}

			// directory
			log.Printf("Downloaded : %v\n", v.ID)
			absoluteFileName := filepath.Join(dir, fmt.Sprintf("%d.%s", v.ID, format))
			//err = saveImage(filepath.Join(dir, fileName), img)
			//if err != nil {
			//	log.Println(err)
			//}

			// send value to channel chImg
			chImg <- Image{filePath: absoluteFileName, img: img}
		}()
	}
	for range photos {
		// chImg receive value
		v := <-chImg
		saveImage(v.filePath, v.img)
		if err != nil {
			log.Println(err)
		}
	}
	//time.Sleep(10 * time.Second)
}
