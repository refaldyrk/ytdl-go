package dto

// The VideoDownload type represents a video download with various properties such as status, message,
// conversion status, video ID, title, file type, file quality, and download link.
// @property {string} Status - The status of the video download process.
// @property {string} Mess - The "Mess" property is a string that represents any additional message or
// information related to the video download.
// @property {string} Conversion - The "Conversion" property in the VideoDownload struct represents the
// status of the video conversion process. It indicates whether the video has been successfully
// converted or not.
// @property {string} Vid - The "Vid" property represents the video ID or unique identifier of the
// video that is being downloaded.
// @property {string} Title - The title of the video being downloaded.
// @property {string} FileType - FileType is a string property that represents the type of the video
// file. It could be any valid file extension such as mp4, avi, mov, etc.
// @property {string} FileQuality - The FileQuality property represents the quality of the downloaded
// video file. It could be a value like "720p", "1080p", "480p", etc., indicating the resolution or
// quality of the video file.
// @property {string} DownloadLink - The DownloadLink property is a string that represents the URL or
// link to download the video.
type VideoDownload struct {
	Status       string `json:"status"`
	Mess         string `json:"mess"`
	Conversion   string `json:"c_status"`
	Vid          string `json:"vid"`
	Title        string `json:"title"`
	FileType     string `json:"ftype"`
	FileQuality  string `json:"fquality"`
	DownloadLink string `json:"dlink"`
}
