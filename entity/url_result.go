package entity

type URLResult struct {
	URL          string `json:"url"`
	Availability bool   `json:"availability"`
	Allowed      bool   `json:"allowed"`
	Error        string `json:"error"`
	FileType     string `json:"fileType"`
	FilePath     string `json:"filePath"`
}

const DownloadFolder string = "test_downloads"
const ArchiveFolder string = "test_zip"
