package service

import "github.com/maximkz561/gin_tutorial/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FindALL() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindALL() []entity.Video {
	return service.videos
}
