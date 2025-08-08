package entity

type Target struct {
	URL []string `json:"urls"`
}

var AllowedExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".pdf":  true,
}