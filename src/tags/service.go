package tags

import (
	"bareksaIntern/src/models"
	"bareksaIntern/src/utils"
	"context"
	"log"

	"github.com/google/uuid"
)

type service struct {
	repository *repository
}

type Agent interface {
	CreateTag(ctx context.Context, payload RequestCreateTag) (data models.Tag, err error)
	SearchOneTag(ctx context.Context, payload RequestSearchTag) (data models.Tag, err error)
	UpdateTag(ctx context.Context, payload RequestUpdateTag) (data models.Tag, err error)
	DeleteTag(ctx context.Context, payload RequestDeleteTag) (err error)
}

func NewService() Agent {
	db := utils.GetDB()
	repo := newRepository(db)
	return &service{
		repository: repo,
	}
}

func (s *service) CreateTag(ctx context.Context, payload RequestCreateTag) (data models.Tag, err error) {
	tag := models.Tag{
		GUID:      uuid.New().String(),
		Name:      payload.Name,
		CreatedAt: utils.GetTimeNow(),
	}
	err = s.repository.CreateTag(ctx, tag)
	if err != nil {
		log.Println(err)
		return
	}
	data = tag
	return
}

func (s *service) SearchOneTag(ctx context.Context, payload RequestSearchTag) (data models.Tag, err error) {
	data, err = s.repository.SearchOneTag(ctx, payload)
	return
}

func (s *service) UpdateTag(ctx context.Context, payload RequestUpdateTag) (data models.Tag, err error) {
	search := RequestSearchTag{
		GUID: payload.GUID,
	}

	tag, err := s.SearchOneTag(ctx, search)
	if err != nil {
		log.Println(err)
		return
	}

	if payload.Name != "" {
		tag.Name = payload.Name
	}

	tag.UpdatedAt.Time = utils.GetTimeNow()
	tag.UpdatedAt.Valid = true

	err = s.repository.UpdateTag(ctx, tag)
	if err != nil {
		log.Println(err)
		return
	}
	data = tag
	return
}

func (s *service) DeleteTag(ctx context.Context, payload RequestDeleteTag) (err error) {
	search := RequestSearchTag{
		GUID: payload.GUID,
	}

	tag, err := s.SearchOneTag(ctx, search)
	if err != nil {
		log.Println(err)
		return
	}

	tag.DeletedAt.Time = utils.GetTimeNow()
	tag.DeletedAt.Valid = true

	err = s.repository.UpdateTag(ctx, tag)
	if err != nil {
		log.Println(err)
		return
	}

	//delete this tag in all news
	err = s.repository.DeleteNewsTag(ctx, payload)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
