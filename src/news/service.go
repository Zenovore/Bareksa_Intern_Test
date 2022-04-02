package news

import (
	"bareksaIntern/src/models"
	"bareksaIntern/src/tags"
	"bareksaIntern/src/utils"
	"context"
	"log"

	"github.com/google/uuid"
)

type service struct {
	repository *repository
}

type Agent interface {
	SearchOneNews(ctx context.Context, payload RequestSearchNews) (data models.News, err error)
	SearchManyNews(ctx context.Context, payload RequestSearchNews) (data []models.News, err error)
	CreateNews(ctx context.Context, payload RequestCreateNews) (data models.News, err error)
	UpdateNews(ctx context.Context, payload RequestUpdateNews) (data models.News, err error)
	DeleteNews(ctx context.Context, payload RequestDeleteNews) (err error)
	FilterNewsByTopic(ctx context.Context, payload RequestSearchNews) (data []models.News, err error)
}

func NewService() Agent {
	db := utils.GetDB()
	repo := newRepository(db)
	return &service{
		repository: repo,
	}
}

func (s *service) SearchOneNews(ctx context.Context, payload RequestSearchNews) (data models.News, err error) {
	data, err = s.repository.SearchOneNews(ctx, payload)

	if err != nil {
		log.Println(err)
		return
	}

	reqSearch := RequestSearchNews{
		GUID: data.GUID,
	}

	data.Tags, err = s.repository.GetNewsTags(ctx, reqSearch)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (s *service) SearchManyNews(ctx context.Context, payload RequestSearchNews) (data []models.News, err error) {
	data, err = s.repository.SearchManyNews(ctx, payload)

	if err != nil {
		log.Println(err)
		return
	}

	for i := 0; i < len(data); i += 1 {
		reqSearch := RequestSearchNews{
			GUID: data[i].GUID,
		}

		data[i].Tags, err = s.repository.GetNewsTags(ctx, reqSearch)
		if err != nil {
			log.Println(err)
			return
		}
	}

	return
}

func (s *service) CreateNews(ctx context.Context, payload RequestCreateNews) (data models.News, err error) {
	news := models.News{
		GUID:      uuid.New().String(),
		Title:     payload.Title,
		Content:   payload.Content,
		CreatedAt: utils.GetTimeNow(),
	}

	err = s.repository.CreateNews(ctx, news)

	if err != nil {
		log.Println(err)
		return
	}

	data = news
	return
}

func (s *service) UpdateNews(ctx context.Context, payload RequestUpdateNews) (data models.News, err error) {
	search := RequestSearchNews{
		GUID: payload.GUID,
	}
	news, err := s.SearchOneNews(ctx, search)

	if err != nil {
		log.Println(err)
		return
	}

	if payload.Title != "" {
		news.Title = payload.Title
	}

	if payload.Content != "" {
		news.Content = payload.Content
	}

	if payload.IsPublishedValid == 1 {
		if payload.IsPublished == 1 {
			news.PublishedAt.Time = utils.GetTimeNow()
			news.PublishedAt.Valid = true
		} else {
			news.PublishedAt.Valid = false
		}
	}

	if payload.AddedTags != nil {
		for i := 0; i < len(payload.AddedTags); i += 1 {
			search := tags.RequestSearchTag{
				Name: payload.AddedTags[i],
			}

			var tag models.Tag
			tag, _ = tagsAgent.SearchOneTag(ctx, search)

			req := RequestAddTag{
				GUIDNews: payload.GUID,
				GUIDTag:  tag.GUID,
			}

			_, err = s.AddTagToNews(ctx, req)
			if err != nil {
				log.Println(err)
				return
			}
		}
	}

	if payload.DeletedTags != nil {
		for i := 0; i < len(payload.DeletedTags); i += 1 {
			search := tags.RequestSearchTag{
				Name: payload.DeletedTags[i],
			}

			var tag models.Tag
			tag, _ = tagsAgent.SearchOneTag(ctx, search)

			req := RequestAddTag{
				GUIDNews: payload.GUID,
				GUIDTag:  tag.GUID,
			}

			_, err = s.DeleteTagFromNews(ctx, req)
			if err != nil {
				log.Println(err)
				return
			}
		}
	}

	news.UpdatedAt.Time = utils.GetTimeNow()
	news.UpdatedAt.Valid = true

	err = s.repository.UpdateNews(ctx, news)
	if err != nil {
		log.Println(err)
		return
	}

	searchNews := RequestSearchNews{
		GUID: payload.GUID,
	}

	data, err = s.SearchOneNews(ctx, searchNews)
	return
}

func (s *service) DeleteNews(ctx context.Context, payload RequestDeleteNews) (err error) {
	search := RequestSearchNews{
		GUID: payload.GUID,
	}
	var news models.News
	news, err = s.SearchOneNews(ctx, search)

	if err != nil {
		log.Println(err)
		return
	}

	news.DeletedAt.Time = utils.GetTimeNow()
	news.DeletedAt.Valid = true

	err = s.repository.UpdateNews(ctx, news)
	return
}

func (s *service) FilterNewsByTopic(ctx context.Context, payload RequestSearchNews) (data []models.News, err error) {
	guid_news, err := s.repository.FilterNewsByTopic(ctx, payload)
	if err != nil {
		log.Println(err)
		return
	}

	for i := 0; i < len(guid_news); i += 1 {
		search := RequestSearchNews{
			GUID: guid_news[i],
		}
		var news models.News
		news, err = s.SearchOneNews(ctx, search)
		if err != nil {
			log.Println()
			return
		}
		data = append(data, news)
	}
	return
}

func (s *service) AddTagToNews(ctx context.Context, payload RequestAddTag) (data models.News, err error) {
	err = s.repository.AddTagToNews(ctx, payload)
	if err != nil {
		log.Println(err)
		return
	}
	search := RequestSearchNews{
		GUID: payload.GUIDNews,
	}
	data, err = s.SearchOneNews(ctx, search)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (s *service) DeleteTagFromNews(ctx context.Context, payload RequestAddTag) (data models.News, err error) {
	err = s.repository.DeleteTagFromNews(ctx, payload)
	if err != nil {
		log.Println(err)
		return
	}
	search := RequestSearchNews{
		GUID: payload.GUIDNews,
	}
	data, err = s.SearchOneNews(ctx, search)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
