// Package data database types.
package data

import (
	"database/sql/driver"
	"gorm.io/gorm"
)

/* ****************************************************************************************************************** *
 *                                                  Type definitions                                                  *
 * ****************************************************************************************************************** */

// Advertisement database type.
type Advertisement struct {
	gorm.Model
	// ContainerID unique container ID.
	ContainerID uint `gorm:"index"`
	// Name advertisement name.
	Name string
	// URL advertisement URL.
	URL string
}

// AssetReference asset (advertisement or image) reference database type.
type AssetReference struct {
	gorm.Model
	// AssetID unique asset ID.
	AssetID uint
	// AssetType asset type.
	AssetType AssetType `gorm:"type:asset_type"`
	// VideoID unique video ID.
	VideoID uint `gorm:"index"`
}

// AssetType asset reference type (AD or IMAGE).
type AssetType string

const (
	// AdvertisementAsset type.
	AdvertisementAsset AssetType = "AD"
	// ImageAsset type.
	ImageAsset AssetType = "IMAGE"
)

// Image database type.
type Image struct {
	gorm.Model
	// ContainerID unique container ID.
	ContainerID uint `gorm:"index"`
	// Name image name.
	Name string
	// URL image URL.
	URL string
}

type Video struct {
	gorm.Model
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
 *                                                Type implementations                                                *
 * ****************************************************************************************************************** */

/* ************************************************* Advertisement ************************************************** */

// Create the advertisement in the database.
func (advertisement Advertisement) Create(db *gorm.DB) error {
	return db.Create(&advertisement).Error
}

// DeleteAdvertisement delete the advertisement matching advertisementID from the database.
func DeleteAdvertisement(db *gorm.DB, advertisementID uint) error {
	return db.Delete(&Advertisement{}, advertisementID).Error
}

// GetAdvertisements get all advertisements matching containerID.
func GetAdvertisements(db *gorm.DB, containerID uint) ([]Advertisement, error) {
	var advertisements []Advertisement
	result := db.Where("container_id = ?", containerID).Find(&advertisements)

	return advertisements, result.Error
}

// Update the advertisement in the database.
func (advertisement Advertisement) Update(db *gorm.DB) error {
	return db.Save(&advertisement).Error
}

/* ************************************************ Asset reference ************************************************* */

// Create the asset reference in the database.
func (assetReference AssetReference) Create(db *gorm.DB) error {
	return db.Create(&assetReference).Error
}

// DeleteAssetReference delete the asset reference matching assetReferenceID from the database.
func DeleteAssetReference(db *gorm.DB, assetReferenceID uint) error {
	return db.Delete(&AssetReference{}, assetReferenceID).Error
}

// GetAssetReferences get all asset references matching videoID.
func GetAssetReferences(db *gorm.DB, videoID uint) ([]AssetReference, error) {
	var assetReferences []AssetReference
	result := db.Where("video_id = ?", videoID).Find(&assetReferences)

	return assetReferences, result.Error
}

// Update the asset reference in the database.
func (assetReference AssetReference) Update(db *gorm.DB) error {
	return db.Save(&assetReference).Error
}

/* *************************************************** Asset type *************************************************** */

func (assetType *AssetType) Scan(value interface{}) error {
	*assetType = AssetType(value.([]byte))

	return nil
}

func (assetType AssetType) Value() (driver.Value, error) {
	return string(assetType), nil
}

/* ***************************************************** Image ****************************************************** */

// Create the image in the database.
func (image Image) Create(db *gorm.DB) error {
	return db.Create(&image).Error
}

// DeleteImage delete the image matching imageID from the database.
func DeleteImage(db *gorm.DB, imageID uint) error {
	return db.Delete(&Image{}, imageID).Error
}

// GetImages get all images matching containerID.
func GetImages(db *gorm.DB, containerID uint) ([]Image, error) {
	var images []Image
	result := db.Where("container_id = ?", containerID).Find(&images)

	return images, result.Error
}

// Update the image in the database.
func (image Image) Update(db *gorm.DB) error {
	return db.Save(&image).Error
}

/* ***************************************************** Video ****************************************************** */

// Create the video in the database.
func (video Video) Create(db *gorm.DB) error {
	return db.Create(&video).Error
}

// DeleteVideo delete the video matching videoID from the database.
func DeleteVideo(db *gorm.DB, videoID uint) error {
	return db.Delete(&Video{}, videoID).Error
}

// GetVideos get all videos matching containerID.
func GetVideos(db *gorm.DB, containerID uint) ([]Video, error) {
	var videos []Video
	result := db.Where("container_id = ?", containerID).Find(&videos)

	return videos, result.Error
}

// Update the video in the database.
func (video Video) Update(db *gorm.DB) error {
	return db.Save(&video).Error
}

/* *************************************************** Video type *************************************************** */

func (videoType *VideoType) Scan(value interface{}) error {
	*videoType = VideoType(value.([]byte))

	return nil
}

func (videoType VideoType) Value() (driver.Value, error) {
	return string(videoType), nil
}
