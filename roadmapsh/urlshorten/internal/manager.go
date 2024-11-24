package internal

import (
	"context"
	"crypto/rand"
	"math/big"
)

type (
	manager struct {
		repo Repo
	}

	Manager interface {
		Create(ctx context.Context, req *CreateReq) (*UrlRecord, error)
		Get(ctx context.Context, shortCode string) (*UrlRecord, error)
		Update(ctx context.Context, req *UpdateReq) (*UrlRecord, error)
		Delete(ctx context.Context, shortCode string) error
		GetStats(ctx context.Context, shortCode string) (*UrlRecord, error)
	}
)

func NewManager(repo Repo) Manager {
	return &manager{
		repo: repo,
	}
}

type CreateReq struct {
	URL string `json:"url"`
}

func (m *manager) Create(ctx context.Context, req *CreateReq) (*UrlRecord, error) {
	var (
		err       error
		shortCode string
	)
	maxRetry := 3
	for i := 0; i < maxRetry; i++ {
		shortCode = generateShortCode()
		if _, err = m.repo.Create(ctx, req.URL, shortCode); err != nil {
			continue
		}
		break
	}

	return m.repo.Get(ctx, shortCode, false)
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateShortCode() string {
	randBytes := make([]byte, 6)

	for i := 0; i < 6; i++ {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			panic(err)
		}
		randBytes[i] = letters[idx.Int64()]
	}

	return string(randBytes)
}

func (m *manager) Get(ctx context.Context, shortCode string) (*UrlRecord, error) {
	result, err := m.repo.Get(ctx, shortCode, false)
	if err != nil {
		return nil, err
	}

	_ = m.repo.IncAccessCount(ctx, shortCode)

	return result, nil
}

type UpdateReq struct {
	ShortCode string `json:"shortCode"`
	URL       string `json:"url"`
}

func (m *manager) Update(ctx context.Context, req *UpdateReq) (*UrlRecord, error) {
	if err := m.repo.Update(context.Background(), req.ShortCode, req.URL); err != nil {
		return nil, err
	}

	return m.repo.Get(context.Background(), req.ShortCode, false)
}

func (m *manager) Delete(ctx context.Context, shortCode string) error {
	return m.repo.Delete(context.Background(), shortCode)
}

func (m *manager) GetStats(ctx context.Context, shortCode string) (*UrlRecord, error) {
	return m.repo.Get(context.Background(), shortCode, true)
}
