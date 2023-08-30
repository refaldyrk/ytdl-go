package dto

// The VideoLinks type is a struct that contains maps of VideoInfo objects for mp4, mp3, and other
// video formats.
// @property Mp4 - A map of string keys to VideoInfo values, representing video links in the MP4
// format.
// @property Mp3 - The `Mp3` property is a map that stores video information for MP3 files. Each key in
// the map represents a unique identifier for a specific MP3 file, and the corresponding value is an
// instance of the `VideoInfo` struct that contains details about the MP3 file.
// @property Other - The "Other" property is a map that contains video links of formats other than MP4
// and MP3. It stores the video links as values, with a unique identifier as the key. The type of the
// values is VideoInfo, which likely contains additional information about the video, such as its title
type VideoLinks struct {
	Mp4   map[string]VideoInfo `json:"mp4"`
	Mp3   map[string]VideoInfo `json:"mp3"`
	Other map[string]VideoInfo `json:"other"`
}

// The above type represents information about a video, including its size, format, quality, text
// description, and key.
// @property {string} Size - The size of the video file.
// @property {string} Format - The "Format" property in the VideoInfo struct represents the format of
// the video file. It could be a file extension like "mp4", "avi", "mov", etc.
// @property {string} Quality - The "Quality" property represents the quality of the video. It can be
// used to indicate the resolution or bitrate of the video.
// @property {string} QText - The QText property in the VideoInfo struct represents the quality of the
// video in text format. It provides a textual description of the video quality, such as "High",
// "Medium", or "Low".
// @property {string} Key - The "Key" property in the VideoInfo struct represents the unique identifier
// or key associated with the video. This key can be used to retrieve or reference the video in a
// database or storage system.
type VideoInfo struct {
	Size    string `json:"size"`
	Format  string `json:"f"`
	Quality string `json:"q"`
	QText   string `json:"q_text"`
	Key     string `json:"k"`
}

// The RelatedVideo type represents a video with a title and a list of video entries.
// @property {string} Title - The Title property is a string that represents the title of the related
// video.
// @property {[]VideoEntry} Contents - The `Contents` property is a slice of `VideoEntry` structs. It
// represents a collection of video entries related to a specific video.
type RelatedVideo struct {
	Title    string       `json:"title"`
	Contents []VideoEntry `json:"contents"`
}

// The VideoEntry type is a struct with two fields, "v" and "t", which are both strings and are tagged
// with json annotations.
// @property {string} V - The "V" property in the VideoEntry struct represents the video ID or video
// URL. It is used to uniquely identify a video.
// @property {string} T - The "T" property in the VideoEntry struct represents the title of the video.
type VideoEntry struct {
	V string `json:"v"`
	T string `json:"t"`
}

// The VideoDetail type represents the details of a video, including its status, title, links, and
// related videos.
// @property {string} Status - The status of the video (e.g., "success", "error").
// @property {string} Mess - The "Mess" property in the VideoDetail struct is a string that represents
// a message or error description related to the video.
// @property {string} Page - The "Page" property in the VideoDetail struct represents the page number
// of the video. It is a string type.
// @property {string} Vid - The "Vid" property in the VideoDetail struct represents the unique
// identifier of the video.
// @property {string} Extractor - The "Extractor" property in the VideoDetail struct is a string that
// represents the source or method used to extract the video.
// @property {string} Title - The title of the video.
// @property {int} T - The property "T" in the VideoDetail struct is of type int.
// @property {string} A - The property "A" is of type string and represents a value associated with the
// video. The specific meaning of this property may vary depending on the context or the application
// using this struct.
// @property {VideoLinks} Links - The `Links` property is of type `VideoLinks` and represents the links
// associated with the video.
// @property {[]RelatedVideo} Related - The "Related" property is an array of "RelatedVideo" objects.
// Each "RelatedVideo" object represents a related video to the main video.
type VideoDetail struct {
	Status    string         `json:"status"`
	Mess      string         `json:"mess"`
	Page      string         `json:"page"`
	Vid       string         `json:"vid"`
	Extractor string         `json:"extractor"`
	Title     string         `json:"title"`
	T         int            `json:"t"`
	A         string         `json:"a"`
	Links     VideoLinks     `json:"links"`
	Related   []RelatedVideo `json:"related"`
}
