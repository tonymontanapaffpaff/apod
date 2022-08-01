package api

import (
	"net/http"

	"github.com/tonymontanapaffpaff/apod/pkg/data"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type AlbumAPI struct {
	albumData *data.AlbumData
}

func NewAlbumAPI(albumData *data.AlbumData) AlbumAPI {
	return AlbumAPI{
		albumData: albumData,
	}
}

func ServeAlbumResource(r *gin.Engine, albumData data.AlbumData) {
	api := &AlbumAPI{albumData: &albumData}
	r.GET("/album", api.GetPictureFromAlbum)
}

func (a AlbumAPI) GetPictureFromAlbum(c *gin.Context) {
	date := c.Query("date")
	if date != "" {
		picturesByDate, err := a.albumData.FindByRequestedAt(date)
		log.Debugf("FindByRequestedAt result: %v", picturesByDate)
		if err != nil {
			log.Error(err)
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, picturesByDate)
		return
	}
	pictures, err := a.albumData.ReadAll()
	log.Debugf("ReadAll result: %v", pictures)
	if err != nil {
		log.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, pictures)
}
