package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// VideoArchive -
type VideoArchive struct {
	ID            uint      `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	YoutubeID     string    `gorm:"column:youtube_id" json:"youtube_id"`
	URL           string    `gorm:"column:url" json:"url"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	DatePublished time.Time `json:"date_published"`
	Author        string    `json:"author"`
	Duration      int       `json:"duration"`
	UploadedBy    string    `json:"uploaded_by"`
}

// VideoArchiveSave -
func VideoArchiveSave(conn *gorm.DB, v *VideoArchive) error {
	return conn.Save(v).Error
}

// VideoArchiveDelete -
func VideoArchiveDelete(conn *gorm.DB, id string) error {
	return conn.Unscoped().Delete(VideoArchive{}, "id = ?", id).Error
}

// VideoArchiveList - return list of all video archives
func VideoArchiveList(conn *gorm.DB) ([]VideoArchive, error) {
	v := []VideoArchive{}
	err := conn.Find(&v).Error
	return v, err
}
