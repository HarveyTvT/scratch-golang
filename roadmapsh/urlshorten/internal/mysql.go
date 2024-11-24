package internal

import (
	"context"
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

type UrlRecord struct {
	bun.BaseModel `bun:"url_records"`

	ID          uint64     `bun:"id,pk" json:"id"`
	URL         string     `bun:"url,notnull" json:"url"`
	ShortCode   string     `bun:"short_code,notnull" json:"shortCode"`
	AccessCount uint64     `bun:"access_count,notnull,default:0" json:"accessCount,omitempty"`
	CreatedAt   *time.Time `bun:"created_at,notnull,default:current_timestamp" json:"createdAt"`
	UpdatedAt   *time.Time `bun:"updated_at,notnull,default:current_timestamp" json:"updatedAt"`
}

type repo struct {
	db *bun.DB
}

type Repo interface {
	Create(ctx context.Context, url, shortCode string) (*UrlRecord, error)
	Get(ctx context.Context, shortCode string, withAccessCount bool) (*UrlRecord, error)
	Update(ctx context.Context, shortCode string, url string) error
	Delete(ctx context.Context, shortCode string) error
	IncAccessCount(ctx context.Context, shortCode string) error
}

func NewRepo(db *bun.DB) *repo {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, url, shortCode string) (*UrlRecord, error) {
	record := &UrlRecord{
		URL:       url,
		ShortCode: shortCode,
	}

	_, err := r.db.NewInsert().Model(record).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return record, nil
}

func (r *repo) Get(ctx context.Context, shortCode string, withAccessCount bool) (*UrlRecord, error) {
	record := &UrlRecord{}
	err := r.db.NewSelect().Model(record).Where("short_code = ?", shortCode).Scan(ctx)
	if err != nil {
		return nil, err
	}
	if !withAccessCount {
		record.AccessCount = 0
	}

	return record, nil
}

func (r *repo) Update(ctx context.Context, shortCode string, url string) error {
	record := &UrlRecord{
		URL: url,
	}

	_, err := r.db.NewUpdate().Model(record).Where("short_code = ?", shortCode).OmitZero().Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Delete(ctx context.Context, shortCode string) error {
	sqlResult, err := r.db.NewDelete().Model(&UrlRecord{}).Where("short_code = ?", shortCode).Exec(ctx)
	if err != nil {
		return err
	}
	rows, err := sqlResult.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *repo) IncAccessCount(ctx context.Context, shortCode string) error {
	_, err := r.db.NewUpdate().Model(&UrlRecord{}).Set("access_count = access_count + 1").Where("short_code = ?", shortCode).Exec(ctx)
	return err
}
