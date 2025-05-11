// Package container container/controller types.
package container

import "RocketContainer.go/internal/data"

/* ************************************************ Type definitions ************************************************ */

// AdvertisementContainer advertisement container/controller type.
type AdvertisementContainer struct {
	ID   uint64
	Name string
	URL  string
}

// AssetReference asset (advertisement or image) reference container/controller type.
type AssetReference struct {
	AssetId   uint64
	AssetType data.AssetType
}

// Container asset container/controller type.
type Container struct {
	Ads    []AdvertisementContainer
	ID     uint64
	Images []ImageContainer
	Name   string
	Videos []VideoContainer
}

// ImageContainer image container/controller type.
type ImageContainer struct {
	ID   uint64
	Name string
	URL  string
}

// VideoContainer video container/controller type.
type VideoContainer struct {
	Assets         []AssetReference
	Description    string
	ExpirationDate string
	ID             uint64
	PlaybackUrl    string
	Title          string
	VideoType      data.VideoType
}

/* ********************************************** Type implementations ********************************************** */

// NewAdvertisementContainer construct an AdvertisementContainer from data.Advertisement.
func NewAdvertisementContainer(advertisement data.Advertisement) AdvertisementContainer {
	return AdvertisementContainer{advertisement.ID, advertisement.Name, advertisement.URL}
}

// NewAssetReference construct an AssetReference from data.AssetReference.
func NewAssetReference(reference data.AssetReference) AssetReference {
	return AssetReference{reference.AssetID, reference.AssetType}
}

// NewImageContainer construct an ImageContainer from data.Image.
func NewImageContainer(image data.Image) ImageContainer {
	return ImageContainer{image.ID, image.Name, image.URL}
}

// NewVideoContainer construct a VideoContainer from data.Video and data.AssetReference.
func NewVideoContainer(video data.Video, assets []data.AssetReference) VideoContainer {
	var references = make([]AssetReference, 0, len(assets))

	for _, asset := range assets {
		references = append(references, NewAssetReference(asset))
	}

	return VideoContainer{
		references, video.Description, video.ExpirationDate, video.ID, video.PlaybackURL, video.Title, video.VideoType,
	}
}
