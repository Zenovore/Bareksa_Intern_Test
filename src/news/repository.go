package news

import (
	"bareksaIntern/src/models"
	"bareksaIntern/src/utils"
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	DB *sqlx.DB
}

func newRepository(db *sqlx.DB) *repository {
	return &repository{
		DB: db}
}

func (r *repository) SearchOneNews(ctx context.Context, payload RequestSearchNews) (res models.News, err error) {
	psq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	qb := psq.Select(
		"guid",
		"title",
		"content",
		"created_at",
		"updated_at",
		"deleted_at",
		"published_at").
		From("news")

	if payload.GUID != "" {
		qb = qb.Where("guid = ?", payload.GUID)
	}

	if payload.Title != "" {
		qb = qb.Where("title = ?", payload.Title)
	}

	if payload.IsPublishedValid == 1 {
		if payload.IsPublished == 1 {
			qb = qb.Where("published_at IS NOT NULL")
		} else {
			qb = qb.Where("published_at IS NULL")
		}
	}

	query, args, _ := qb.Where("deleted_at IS NULL").ToSql()
	err = r.DB.GetContext(ctx, &res, query, args...)

	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (r *repository) SearchManyNews(ctx context.Context, payload RequestSearchNews) (res []models.News, err error) {
	psq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	qb := psq.Select(
		"guid",
		"title",
		"content",
		"created_at",
		"updated_at",
		"deleted_at",
		"published_at").
		From("news")

	if payload.GUID != "" {
		qb = qb.Where("guid = ?", payload.GUID)
	}

	if payload.Title != "" {
		qb = qb.Where("title = ?", payload.Title)
	}

	if payload.IsPublishedValid == 1 {
		if payload.IsPublished == 1 {
			qb = qb.Where("published_at IS NOT NULL")
		} else {
			qb = qb.Where("published_at IS NULL")
		}
	}

	if payload.IsDeletedValid == 1 {
		if payload.IsDeleted == 1 {
			qb = qb.Where("deleted_at IS NOT NULL")
		} else {
			qb = qb.Where("deleted_at IS NULL")
		}
	} else {
		qb = qb.Where("deleted_at IS NULL")
	}

	query, args, _ := qb.ToSql()
	err = r.DB.SelectContext(ctx, &res, query, args...)

	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (r *repository) GetNewsTags(ctx context.Context, payload RequestSearchNews) (tags []models.Tag, err error) {
	psq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	qb := psq.Select(
		"t.guid",
		"t.name",
		"t.created_at",
		"t.deleted_at",
		"t.updated_at").
		From("news_tags nt").LeftJoin("tags t ON nt.guid_tag = t.guid")

	query, args, _ := qb.Where("guid_news = ?", payload.GUID).Where("nt.deleted_at IS NULL").ToSql()
	err = r.DB.SelectContext(ctx, &tags, query, args...)

	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (r *repository) UpdateNews(ctx context.Context, news models.News) (err error) {
	psq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	qb := psq.Update("news").
		Set("title", news.Title).
		Set("content", news.Content).
		Set("updated_at", news.UpdatedAt).
		Set("deleted_at", news.DeletedAt).
		Set("published_at", news.PublishedAt).
		Where("guid = ?", news.GUID)

	query, args, _ := qb.ToSql()

	tx, err := r.DB.BeginTxx(ctx, nil)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
	}

	return
}

func (r *repository) CreateNews(ctx context.Context, news models.News) (err error) {
	psq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	qb := psq.Insert("news").Columns(
		"guid",
		"title",
		"content",
		"created_at").Values(
		news.GUID,
		news.Title,
		news.Content,
		news.CreatedAt)

	query, args, _ := qb.ToSql()

	tx, err := r.DB.BeginTxx(ctx, nil)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
	}

	return
}

func (r *repository) FilterNewsByTopic(ctx context.Context, payload RequestSearchNews) (res []string, err error) {
	psq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	qb := psq.Select(
		"tn.guid_news").
		From("topic_news tn").
		LeftJoin("topics t ON tn.guid_topic=t.guid").
		Where("t.name = ?", payload.Topic).
		Where("tn.deleted_at IS NULL").
		Where("t.deleted_at IS NULL")

	query, args, _ := qb.ToSql()
	err = r.DB.SelectContext(ctx, &res, query, args...)

	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (r *repository) AddTagToNews(ctx context.Context, payload RequestAddTag) (err error) {
	psq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	qb := psq.Insert("news_tags").Columns(
		"guid",
		"guid_news",
		"guid_tag",
		"created_at").Values(
		uuid.New().String(),
		payload.GUIDNews,
		payload.GUIDTag,
		utils.GetTimeNow())

	query, args, _ := qb.ToSql()

	tx, err := r.DB.BeginTxx(ctx, nil)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
	}

	return
}

func (r *repository) DeleteTagFromNews(ctx context.Context, payload RequestAddTag) (err error) {
	psq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	qb := psq.Update("news_tags").
		Set("deleted_at", utils.GetTimeNow()).
		Where("guid_news = ?", payload.GUIDNews).
		Where("guid_tag = ?", payload.GUIDTag)

	query, args, _ := qb.ToSql()

	tx, err := r.DB.BeginTxx(ctx, nil)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
	}

	return
}
