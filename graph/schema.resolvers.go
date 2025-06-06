package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.73

import (
	"context"
	"strconv"

	"RocketContainer.go/graph/model"
	"RocketContainer.go/internal/data"
)

/* ****************************************************************************************************************** *
 *                                                     Mutations                                                      *
 * ****************************************************************************************************************** */

// CreateAsset is the resolver for the createAsset field.
func (r *mutationResolver) CreateAsset(ctx context.Context, input model.NewAsset) (uint, error) {
	asset, err := data.CreateAsset(input)

	return asset.ID, err
}

// CreateVideo is the resolver for the createVideo field.
func (r *mutationResolver) CreateVideo(ctx context.Context, input model.NewVideo) (uint, error) {
	video, err := data.CreateVideo(input)

	return video.ID, err
}

// DeleteAsset is the resolver for the deleteAsset field.
func (r *mutationResolver) DeleteAsset(ctx context.Context, input uint) (bool, error) {
	err := data.DeleteAsset(input)

	return err != nil, err
}

// DeleteVideo is the resolver for the deleteVideo field.
func (r *mutationResolver) DeleteVideo(ctx context.Context, input uint) (bool, error) {
	err := data.DeleteVideo(input)

	return err != nil, err
}

// UpdateAsset is the resolver for the updateAsset field.
func (r *mutationResolver) UpdateAsset(ctx context.Context, input model.UpdateAsset) (bool, error) {
	err := data.UpdateAsset(input)

	return err != nil, err
}

// UpdateVideo is the resolver for the updateVideo field.
func (r *mutationResolver) UpdateVideo(ctx context.Context, input model.UpdateVideo) (bool, error) {
	err := data.UpdateVideo(input)

	return err != nil, err
}

/* ****************************************************************************************************************** *
 *                                                      Queries                                                       *
 * ****************************************************************************************************************** */

// Advertisements is the resolver for the advertisements field.
func (r *queryResolver) Advertisements(ctx context.Context, containerID uint) ([]*model.Asset, error) {
	assets, err := data.GetAssets(containerID, data.Advertisement)

	if err != nil {
		return []*model.Asset{}, err
	}

	results := make([]*model.Asset, 0, len(assets))

	for _, asset := range assets {
		results = append(
			results,
			&model.Asset{
				AssetType: model.AssetTypeAdvertisement,
				ID:        asset.ID,
				Name:      asset.Name,
				URL:       asset.URL,
			},
		)
	}

	return results, nil
}

// Container is the resolver for the container field.
func (r *queryResolver) Container(ctx context.Context, containerID uint) (*model.Container, error) {
	videos, err := data.GetVideosByContainer(containerID)

	if err != nil {
		return &model.Container{}, err
	}

	advertisements := make([]*model.Asset, 0, 16)
	images := make([]*model.Asset, 0, 16)
	modelVideos := make([]*model.Video, 0, len(videos))

	for _, video := range videos {
		assets := make([]uint, 0, len(video.Assets))

		for _, asset := range video.Assets {
			if asset.AssetType == data.Advertisement {
				advertisements = append(
					advertisements,
					&model.Asset{
						AssetType: model.AssetTypeAdvertisement,
						ID:        asset.ID,
						Name:      asset.Name,
						URL:       asset.URL,
					},
				)
			} else {
				images = append(
					images,
					&model.Asset{
						AssetType: model.AssetTypeImage,
						ID:        asset.ID,
						Name:      asset.Name,
						URL:       asset.URL,
					},
				)
			}

			assets = append(
				assets,
				asset.ID,
			)
		}

		modelVideos = append(
			modelVideos,
			&model.Video{
				Assets:         assets,
				Description:    video.Description,
				ExpirationDate: video.ExpirationDate,
				ID:             video.ID,
				PlaybackURL:    video.PlaybackURL,
				Title:          video.Title,
				VideoType:      model.VideoType(video.VideoType),
			},
		)
	}

	result := model.Container{
		Advertisements: advertisements,
		ID:             containerID,
		Images:         images,
		Name:           getContainerName(containerID, advertisements, images),
		Videos:         modelVideos,
	}

	return &result, nil
}

// Containers is the resolver for the containers field.
func (r *queryResolver) Containers(ctx context.Context) ([]*model.Container, error) {
	videos, err := data.GetVideos()

	if err != nil {
		return []*model.Container{}, err
	}

	advertisementMap := make(map[uint][]*model.Asset, 16)
	imageMap := make(map[uint][]*model.Asset, 16)
	videoMap := make(map[uint][]*model.Video, 16)

	for _, video := range videos {
		advertisements := make([]*model.Asset, 0, 16)
		assets := make([]uint, 0, len(video.Assets))
		images := make([]*model.Asset, 0, 16)

		for _, asset := range video.Assets {
			if asset.AssetType == data.Advertisement {
				advertisements = append(
					advertisements,
					&model.Asset{
						AssetType: model.AssetTypeAdvertisement,
						ID:        asset.ID,
						Name:      asset.Name,
						URL:       asset.URL,
					},
				)
			} else {
				images = append(
					images,
					&model.Asset{
						AssetType: model.AssetTypeImage,
						ID:        asset.ID,
						Name:      asset.Name,
						URL:       asset.URL,
					},
				)
			}

			assets = append(
				assets,
				asset.ID,
			)
		}

		advertisementMap[video.ContainerID] = append(advertisementMap[video.ContainerID], advertisements...)
		imageMap[video.ContainerID] = append(imageMap[video.ContainerID], images...)
		videoMap[video.ContainerID] = append(
			videoMap[video.ContainerID],
			&model.Video{
				Assets:         assets,
				Description:    video.Description,
				ExpirationDate: video.ExpirationDate,
				ID:             video.ID,
				PlaybackURL:    video.PlaybackURL,
				Title:          video.Title,
				VideoType:      model.VideoType(video.VideoType),
			},
		)
	}

	results := make([]*model.Container, 0, len(videoMap))

	for containerID, videos := range videoMap {
		advertisements := advertisementMap[containerID]
		images := imageMap[containerID]

		results = append(
			results,
			&model.Container{
				Advertisements: advertisements,
				ID:             containerID,
				Images:         images,
				Name:           getContainerName(containerID, advertisements, images),
				Videos:         videos,
			},
		)
	}

	return results, nil
}

// Images is the resolver for the images field.
func (r *queryResolver) Images(ctx context.Context, containerID uint) ([]*model.Asset, error) {
	assets, err := data.GetAssets(containerID, data.Image)

	if err != nil {
		return []*model.Asset{}, err
	}

	results := make([]*model.Asset, 0, len(assets))

	for _, asset := range assets {
		results = append(
			results,
			&model.Asset{
				AssetType: model.AssetTypeImage,
				ID:        asset.ID,
				Name:      asset.Name,
				URL:       asset.URL,
			},
		)
	}

	return results, nil
}

// Videos is the resolver for the videos field.
func (r *queryResolver) Videos(ctx context.Context, containerID uint) ([]*model.Video, error) {
	videos, err := data.GetVideosByContainer(containerID)

	if err != nil {
		return []*model.Video{}, err
	}

	results := make([]*model.Video, 0, len(videos))

	for _, video := range videos {
		assets := make([]uint, 0, len(video.Assets))

		for _, asset := range video.Assets {
			assets = append(
				assets,
				asset.ID,
			)
		}

		results = append(
			results,
			&model.Video{
				Assets:         assets,
				Description:    video.Description,
				ExpirationDate: video.ExpirationDate,
				ID:             video.ID,
				PlaybackURL:    video.PlaybackURL,
				Title:          video.Title,
				VideoType:      model.VideoType(video.VideoType),
			},
		)
	}

	return results, nil
}

/* ****************************************************************************************************************** *
 *                                                     Resolvers                                                      *
 * ****************************************************************************************************************** */

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

/* ****************************************************************************************************************** *
 *                                                 Private functions                                                  *
 * ****************************************************************************************************************** */

func getContainerName(containerID uint, advertisements []*model.Asset, images []*model.Asset) string {
	adsName := ""
	imagesName := ""

	if len(advertisements) > 0 {
		adsName = "_ads"
	}

	if len(images) > 0 {
		imagesName = "_images"
	}

	return "container-" + strconv.FormatUint(uint64(containerID), 10) + adsName + imagesName + "_videos"
}
