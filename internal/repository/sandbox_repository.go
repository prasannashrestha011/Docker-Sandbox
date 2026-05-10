package repository

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"main/internal/repository/model"
	sandbox_type "main/internal/sandbox/types"
)

// SandboxRepository defines persistence methods for sandbox sessions.
type SandboxRepository interface {
	Create(ctx context.Context, sandbox *model.Sandbox) error
	FindByID(ctx context.Context, id uuid.UUID) (*model.Sandbox, error)
	FindBySessionID(ctx context.Context, sessionID uuid.UUID) (*model.Sandbox, error)
	ListByUserID(ctx context.Context, userID string) ([]model.Sandbox, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status sandbox_type.SandboxState) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type gormSandboxRepository struct {
	db *gorm.DB
}

// NewSandboxRepository returns a GORM-backed SandboxRepository.
func NewSandboxRepository(db *gorm.DB) SandboxRepository {
	return &gormSandboxRepository{db: db}
}

func (r *gormSandboxRepository) Create(ctx context.Context, sandbox *model.Sandbox) error {
	return r.db.WithContext(ctx).Model(&model.Sandbox{}).Create(sandbox).Error
}

func (r *gormSandboxRepository) FindByID(ctx context.Context, id uuid.UUID) (*model.Sandbox, error) {
	var sandbox model.Sandbox
	err := r.db.WithContext(ctx).First(&sandbox, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &sandbox, nil
}

func (r *gormSandboxRepository) FindBySessionID(ctx context.Context, sessionID uuid.UUID) (*model.Sandbox, error) {
	var sandbox model.Sandbox
	err := r.db.WithContext(ctx).First(&sandbox, "session_id = ?", sessionID).Error
	if err != nil {
		return nil, err
	}
	return &sandbox, nil
}

func (r *gormSandboxRepository) ListByUserID(ctx context.Context, userID string) ([]model.Sandbox, error) {
	var sandboxes []model.Sandbox
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Order("created_at desc").Find(&sandboxes).Error
	if err != nil {
		return nil, err
	}
	return sandboxes, nil
}

func (r *gormSandboxRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status sandbox_type.SandboxState) error {
	return r.db.WithContext(ctx).Model(&model.Sandbox{}).Where("id = ?", id).Update("status", status).Error
}

func (r *gormSandboxRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&model.Sandbox{}, "id = ?", id).Error
}
