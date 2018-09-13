package model

// Attachment - discord message attachment
type Attachment struct {
	MessageID string `gorm:"primary_key" json:"id"`
	URL       string `json:"url"`
	ProxyURL  string `json:"proxy_url"`
	Filename  string `json:"filename"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Size      int    `json:"size"`
}
