package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Davincible/goinsta/v3"
)

func main() {

	postVideo()

}

func postVideo() {
	insta := goinsta.New("Goober", "Fuck OFF")
	const max_day = 1000000
	if err := insta.Login(); err != nil {
		panic(err)
	}
	defer insta.Logout()

	filePath := "output.mp4"

	log.Printf("Attempting to open file: %s", filePath)

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	for i := 1; i < max_day; i++ {
		i2str := strconv.Itoa(i)
		var caption string = "Day: " + i2str + "\n\n\n" + "#smilingfriends #everydaypost #adultswilm"

		item, err := insta.Upload(&goinsta.UploadOptions{
			File:    file,
			Caption: caption,
		})
		if err != nil {
			log.Fatalf("Failed to upload video: %s", err)
		}
		log.Printf("Video posted with ID: %s", item.ID)
		time.Sleep(86400 * time.Second)
	}
}
