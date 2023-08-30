package main

import (
	"encoding/json"
	"fmt"

	"github.com/refaldyrk/ytdl-go/service"
)

func main() {
	yt := service.NewYtdl()

	result, err := yt.DetailVideo("https://www.youtube.com/watch?v=yC6e7lY8BWQ")
	if err != nil {
		fmt.Println(err)
	}

	//Handle Mp4
	for _, v := range result.Links.Mp4 {
		data := map[string]string{}

		//Marshal
		byteJson, _ := json.Marshal(v)

		//Unmarshal
		err := json.Unmarshal(byteJson, &data)
		if err != nil {
			fmt.Println(err)
		}

		//Download All Key
		downloadVid, err := yt.DownloadVid(data["k"])
		if err != nil {
			fmt.Println(err)
		}

		//Print All Link Download
		fmt.Println(downloadVid.DownloadLink)
		fmt.Println()
	}
}
