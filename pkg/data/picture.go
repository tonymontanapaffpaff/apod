package data

import (
	"fmt"
	"gorm.io/gorm"
)

type PictureData struct {
	db *gorm.DB
}

func NewPictureData(db *gorm.DB) *PictureData {
	return &PictureData{db: db}
}

type Picture struct {
	PictureId      int `gorm:"primary_key"`
	Copyright      string
	Date           string
	Explanation    string
	HDUrl          string
	MediaType      string
	ServiceVersion string
	Title          string
	Url            string
}

func (d PictureData) Add(picture Picture) (int, error) {
	var pictureId int
	result := d.db.Raw("INSERT INTO pictures VALUES (DEFAULT,?,?,?,?,?,?,?,?) RETURNING picture_id",
		picture.Copyright, picture.Date, picture.Explanation, picture.HDUrl, picture.MediaType,
		picture.ServiceVersion, picture.Title, picture.Url).Scan(&pictureId)
	if result.Error != nil {
		return -1, fmt.Errorf("can't add picture, error: %w", result.Error)
	}
	return pictureId, nil
}

func (d PictureData) FindByDate(date string) (Picture, error) {
	var findByResult Picture
	result := d.db.Where("date = ?", date).Find(&findByResult)
	if result.Error != nil {
		return Picture{}, fmt.Errorf("can't read picture with given date, error: %w", result.Error)
	}
	return findByResult, nil
}
