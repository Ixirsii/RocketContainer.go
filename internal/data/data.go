// Package data database types.
package data

import (
	"RocketContainer.go/graph/model"
	"database/sql/driver"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
	"os"
)

var database *gorm.DB
var logger *zap.Logger

/* ****************************************************************************************************************** *
 *                                                  Type definitions                                                  *
 * ****************************************************************************************************************** */

// Asset database type.
type Asset struct {
	gorm.Model
	// AssetType asset type.
	AssetType AssetType `gorm:"type:asset_type"`
	// ContainerID unique container ID.
	ContainerID uint `gorm:"index"`
	// Name asset name.
	Name string
	// URL asset URL.
	URL string
	// VideoID video ID foreign key.
	VideoID uint `gorm:"index"`
}

// AssetType asset reference type (ADVERTISEMENT or IMAGE).
type AssetType string

const (
	// Advertisement type.
	Advertisement AssetType = "ADVERTISEMENT"
	// Image type.
	Image AssetType = "IMAGE"
)

type Video struct {
	gorm.Model
	// Assets that belong to the video.
	Assets []Asset
	// ContainerID unique container ID.
	ContainerID uint `gorm:"index"`
	// Description video description.
	Description string
	// ExpirationDate expiration date.
	ExpirationDate string
	// PlaybackURL video playback URL.
	PlaybackURL string
	// Title video title.
	Title string
	// VideoType video type (CLIP, EPISODE, or MOVIE).
	VideoType VideoType `gorm:"type:video_type"`
}

// VideoType video type (CLIP, EPISODE, or MOVIE).
type VideoType string

const (
	// Clip video clip.
	Clip VideoType = "CLIP"
	// Episode television episode.
	Episode VideoType = "EPISODE"
	// Movie full-length film.
	Movie VideoType = "MOVIE"
)

/* ****************************************************************************************************************** *
 *                                                     Functions                                                      *
 * ****************************************************************************************************************** */

// InitDb initialize database.
func InitDb() {
	logger = zap.L().Named("database")
	gormLogger := zapgorm2.New(logger)
	gormLogger.SetAsDefault()

	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")

	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPass + " dbname=" + dbName + " port=" + dbPort

	db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: gormLogger})
	if dbErr != nil {
		logger.Fatal("Failed to connect to database", zap.Error(dbErr))
	}

	migrationErr := database.AutoMigrate(&Asset{}, &Video{})
	if migrationErr != nil {
		logger.Fatal("Failed to migrate database", zap.Error(migrationErr))
	}

	database = db
}

/* ************************************************* Asset ************************************************** */

// CreateAsset create the asset in the database.
func CreateAsset(new model.NewAsset) (Asset, error) {
	logger.Debug(
		"Creating asset",
		zap.String("assetType", string(new.AssetType)),
		zap.Uint("containerID", new.ContainerID),
		zap.String("name", new.Name),
		zap.String("url", new.URL),
		zap.Uint("videoID", new.VideoID),
	)

	asset := Asset{
		AssetType:   AssetType(new.AssetType),
		ContainerID: new.ContainerID,
		Name:        new.Name,
		URL:         new.URL,
		VideoID:     new.VideoID,
	}
	result := database.Create(&asset)

	return asset, result.Error
}

// DeleteAsset delete the asset matching assetID from the database.
func DeleteAsset(assetID uint) error {
	logger.Debug("Deleting asset", zap.Uint("assetID", assetID))

	return database.Delete(&Asset{}, assetID).Error
}

// GetAssets get all assets matching containerID and assetType.
func GetAssets(containerID uint, assetType AssetType) ([]Asset, error) {
	logger.Debug(
		"Getting assets",
		zap.Uint("containerID", containerID),
		zap.String("assetType", string(assetType)),
	)

	var assets []Asset
	result := database.Where("container_id = ? AND asset_type = ?", containerID, assetType).Find(&assets)

	return assets, result.Error
}

// UpdateAsset update the asset in the database.
func UpdateAsset(update model.UpdateAsset) error {
	logger.Debug(
		"Updating asset",
		zap.String("assetType", string(update.AssetType)),
		zap.Uint("containerID", update.ContainerID),
		zap.Uint("id", update.ID),
		zap.String("name", update.Name),
		zap.String("url", update.URL),
		zap.Uint("videoID", update.VideoID),
	)

	asset := Asset{
		Model:       gorm.Model{ID: update.ID},
		AssetType:   AssetType(update.AssetType),
		ContainerID: update.ContainerID,
		Name:        update.Name,
		URL:         update.URL,
		VideoID:     update.VideoID,
	}

	return database.Save(&asset).Error
}

/* *************************************************** Asset type *************************************************** */

func (assetType *AssetType) Scan(value interface{}) error {
	*assetType = AssetType(value.([]byte))

	return nil
}

func (assetType AssetType) Value() (driver.Value, error) {
	return string(assetType), nil
}

/* ***************************************************** Video ****************************************************** */

// CreateVideo create the video in the database.
func CreateVideo(new model.NewVideo) (Video, error) {
	logger.Debug(
		"Creating video",
		zap.Uint("containerID", new.ContainerID),
		zap.String("description", new.Description),
		zap.String("expirationDate", new.ExpirationDate),
		zap.String("playbackUrl", new.PlaybackURL),
		zap.String("title", new.Title),
		zap.String("videoType", string(new.VideoType)),
	)

	video := Video{
		ContainerID:    new.ContainerID,
		Description:    new.Description,
		ExpirationDate: new.ExpirationDate,
		PlaybackURL:    new.PlaybackURL,
		Title:          new.Title,
		VideoType:      VideoType(new.VideoType),
	}
	result := database.Create(&video)

	return video, result.Error
}

// DeleteVideo delete the video matching videoID from the database.
func DeleteVideo(videoID uint) error {
	logger.Debug("Deleting video", zap.Uint("videoID", videoID))

	return database.Delete(&Video{}, videoID).Error
}

// GetVideos get all videos.
func GetVideos() ([]Video, error) {
	logger.Debug("Getting all videos")

	var videos []Video
	result := database.Model(&Video{}).Preload("Assets").Find(&videos)

	return videos, result.Error
}

// GetVideosByContainer get all videos matching containerID.
func GetVideosByContainer(containerID uint) ([]Video, error) {
	logger.Debug("Getting videos", zap.Uint("containerID", containerID))

	var videos []Video
	result := database.Model(&Video{}).Preload("Assets").Where("container_id = ?", containerID).Find(&videos)

	return videos, result.Error
}

// UpdateVideo update the video in the database.
func UpdateVideo(update model.UpdateVideo) error {
	logger.Debug(
		"Updating video",
		zap.Uint("containerID", update.ContainerID),
		zap.Uint("id", update.ID),
		zap.String("description", update.Description),
		zap.String("expirationDate", update.ExpirationDate),
		zap.String("playbackUrl", update.PlaybackURL),
		zap.String("title", update.Title),
		zap.String("videoType", string(update.VideoType)),
	)

	video := Video{
		Model:          gorm.Model{ID: update.ID},
		ContainerID:    update.ContainerID,
		Description:    update.Description,
		ExpirationDate: update.ExpirationDate,
		PlaybackURL:    update.PlaybackURL,
		Title:          update.Title,
		VideoType:      VideoType(update.VideoType),
	}

	return database.Save(&video).Error
}

/* *************************************************** Video type *************************************************** */

func (videoType *VideoType) Scan(value interface{}) error {
	*videoType = VideoType(value.([]byte))

	return nil
}

func (videoType VideoType) Value() (driver.Value, error) {
	return string(videoType), nil
}
