package tags

import (
	"bareksaIntern/src/models"
	"bareksaIntern/src/utils"
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	DB *sqlx.DB
}

func newRepository(db *sqlx.DB) *repository {
	return &repository{
		DB: db}
}

func (r *repository) SearchOneTag(ctx context.Context, payload RequestSearchTag) (res models.Tag, err error) {
	psq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	qb := psq.Select(
		"guid",
		"name",
		"created_at",
		"updated_at",
		"deleted_at").
		From("tags")

	if payload.GUID != "" {
		qb = qb.Where("guid = ?", payload.GUID)
	}

	if payload.Name != "" {
		qb = qb.Where("name = ?", payload.Name)
	}

	query, args, _ := qb.Where("deleted_at IS NULL").ToSql()
	err = r.DB.GetContext(ctx, &res, query, args...)

	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (r *repository) CreateTag(ctx context.Context, tag models.Tag) (err error) {
	psq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	qb := psq.Insert("tags").Columns(
		"guid",
		"name",
		"created_at").Values(
		tag.GUID,
		tag.Name,
		tag.CreatedAt)

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

func (r *repository) UpdateTag(ctx context.Context, tag models.Tag) (err error) {
	psq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	qb := psq.Update("tags").
		Set("name", tag.Name).
		Set("updated_at", tag.UpdatedAt).
		Set("deleted_at", tag.DeletedAt).
		Where("guid = ?", tag.GUID)

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

func (r *repository) DeleteNewsTag(ctx context.Context, payload RequestDeleteTag) (err error) {
	psq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	qb := psq.Update("news_tags").
		Set("deleted_at", utils.GetTimeNow()).
		Where("guid = ?", payload.GUID)

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
