package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/tonymontanapaffpaff/apod/pkg/data"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type PictureAPI struct {
	pictureData *data.PictureData
	albumData   *data.AlbumData
}

func NewPictureAPI(pictureData *data.PictureData, albumData *data.AlbumData) PictureAPI {
	return PictureAPI{
		pictureData: pictureData,
		albumData:   albumData,
	}
}

func ServePictureResource(r *gin.Engine, pictureData data.PictureData, albumData data.AlbumData) {
	api := &PictureAPI{pictureData: &pictureData, albumData: &albumData}
	r.GET("/picture", api.GetPicture)
}

type PictureJSON struct {
	Copyright      string `json:"copyright"`
	Date           string `json:"date"`
	Explanation    string `json:"explanation"`
	HDUrl          string `json:"hdurl"`
	MediaType      string `json:"media_type"`
	ServiceVersion string `json:"service_version"`
	Title          string `json:"title"`
	Url            string `json:"url"`
}

func (a PictureAPI) GetPicture(c *gin.Context) {
	apiKey := c.Query("api_key")
	save := c.Query("save")
	currentDate := time.Now().Format("2006-01-02")
	picture, err := a.pictureData.FindByDate(currentDate)
	log.Debugf("FindByDate picture id: %d", picture.PictureId)
	if err != nil {
		log.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	if picture.PictureId == 0 {
		serverData, err := fetchServerData(c, apiKey)
		log.Debugf("FetchServerData result: %v", serverData)
		if err != nil {
			log.Error(err)
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		if save == "true" {
			pictureId, err := a.pictureData.Add(data.Picture{
				Copyright:      serverData.Copyright,
				Date:           serverData.Date,
				Explanation:    serverData.Explanation,
				HDUrl:          serverData.HDUrl,
				MediaType:      serverData.MediaType,
				ServiceVersion: serverData.ServiceVersion,
				Title:          serverData.Title,
				Url:            serverData.Url,
			})
			log.Debugf("Add picture result: %d", pictureId)
			if err != nil {
				log.Error(err)
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
			_, err = a.albumData.Add(currentDate, pictureId)
			if err != nil {
				log.Error(err)
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
		}
		c.JSON(http.StatusOK, serverData)
		return
	}
	c.JSON(http.StatusOK, picture)
}

func fetchServerData(c *gin.Context, key string) (*PictureJSON, error) {
	if key == "" {
		key = "DEMO_KEY"
	}
	req, err := http.NewRequestWithContext(
		c,
		http.MethodGet,
		fmt.Sprintf("https://api.nasa.gov/planetary/apod?api_key=%s", key),
		nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusForbidden {
		return nil, fmt.Errorf("an invalid api_key was supplied")
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	pictureJSON := PictureJSON{}
	err = json.Unmarshal(body, &pictureJSON)
	if err != nil {
		return nil, err
	}
	return &pictureJSON, nil
}
