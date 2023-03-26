package entities

// image100, image75, etc. - name of images in different sizes
type Image struct {
	Id       int    `json:"id"`
	Image100 string `json:"image100"`
	Image75  string `json:"image75"`
	Image50  string `json:"image50"`
	Image25  string `json:"image25"`
}
