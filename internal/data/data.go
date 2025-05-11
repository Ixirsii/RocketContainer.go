// Package data database types.
package data

// Advertisement database type.
type Advertisement struct {
	// ContainerID unique container ID.
	ContainerId uint64
	// ID unique advertisement ID.
	ID uint64
	// Name advertisement name.
	Name string
	// URL advertisement URL.
	URL string
}

// AssetReference asset (advertisement or image) reference database type.
type AssetReference struct {
	// AssetID unique asset ID.
	AssetID uint64
	// AssetType asset type.
	AssetType AssetType
	// VideoID unique video ID.
	VideoID uint64
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
	// ContainerID unique container ID.
	ContainerId uint64
	// ID unique image ID.
	ID uint64
	// Name image name.
	Name string
	// URL image URL.
	URL string
}

type Video struct {
	// ContainerID unique container ID.
	ContainerId uint64
	// Description video description.
	Description string
	// ExpirationDate expiration date.
	ExpirationDate string
	// ID unique video ID.
	ID uint64
	// PlaybackURL video playback URL.
	PlaybackURL string
	// Title video title.
	Title string
	// VideoType video type (CLIP, EPISODE, or MOVIE).
	VideoType VideoType
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
