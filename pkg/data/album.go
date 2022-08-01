package data

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AlbumData struct {
	db *gorm.DB
}

func NewAlbumData(db *gorm.DB) *AlbumData {
	return &AlbumData{db: db}
}

type (
	Album struct {
		Id          int `gorm:"primary_key"`
		RequestedAt string
		PictureId   int
		Picture     Picture `gorm:"references:PictureId"`
	}
	FindByResult struct {
		Id    int
		Title string
		Url   string
	}
	ReadAllResult struct {
		Id          int
		RequestedAt string
		Title       string
		Url         string
	}
)

func (Album) TableName() string {
	return "album"
}

func (d AlbumData) Add(requestedAt string, pictureId int) (int, error) {
	var returned Album
	album := Album{
		RequestedAt: requestedAt,
		PictureId:   pictureId,
	}
	result := d.db.Model(&returned).Clauses(
		clause.Returning{
			Columns: []clause.Column{clause.PrimaryColumn},
		}).Create(&album)
	if result.Error != nil {
		return -1, fmt.Errorf("can't add picture to album, error: %w", result.Error)
	}
	return returned.Id, nil
}

func (d AlbumData) FindByRequestedAt(date string) ([]FindByResult, error) {
	var findByResults []FindByResult
	result := d.db.Model(&Album{}).
		Select("album.id, pictures.title, pictures.url").
		Joins("left join pictures on pictures.picture_id = album.picture_id").
		Where("album.requested_at = ?", date).
		Scan(&findByResults)
	if result.Error != nil {
		return nil, fmt.Errorf("can't read picture with given date, error: %w", result.Error)
	}
	return findByResults, nil
}

func (d AlbumData) ReadAll() ([]ReadAllResult, error) {
	var readAllResults []ReadAllResult
	result := d.db.Model(&Album{}).
		Select("album.id, album.requested_at, pictures.title, pictures.url").
		Joins("left join pictures on pictures.picture_id = album.picture_id").
		Scan(&readAllResults)
	if result.Error != nil {
		return nil, fmt.Errorf("can't read album, error: %w", result.Error)
	}
	return readAllResults, nil
}
