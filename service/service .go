package service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/refaldyrk/ytdl-go/constant"
	"github.com/refaldyrk/ytdl-go/dto"
)

// The YTDL type is a struct in the Go programming language.
type YTDL struct {
	vid string
}

// The function returns a new instance of the YTDL struct.
func NewYtdl() *YTDL {
	return &YTDL{vid: ""}
}

// The `DetailVideo` function is a method of the `YTDL` struct. It takes a `uri` string as a parameter
// and returns a `dto.VideoDetail` object and an error.
func (y *YTDL) DetailVideo(uri string) (dto.VideoDetail, error) {
	var result dto.VideoDetail
	formData := url.Values{}
	formData.Set("k_query", uri)

	response, err := http.PostForm(constant.BASE_API_YTDL_ANALYZE, formData)
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return result, err
	}

	y.vid = result.Vid
	return result, nil
}

func (y *YTDL) DownloadVid(k string) (dto.VideoDownload, error) {
	if y.vid == "" {
		return dto.VideoDownload{}, errors.New("vid not found")
	}
	var result dto.VideoDownload
	formData := url.Values{}
	formData.Set("vid", y.vid)
	formData.Set("k", k)

	response, err := http.PostForm(constant.BASE_API_YTDL_DOWNLOAD, formData)
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return result, err
	}

	return result, nil
}
