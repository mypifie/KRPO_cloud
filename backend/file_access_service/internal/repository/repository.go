package repository 

import (
	"context"

	"gorm.io/gorm"

	"file-access-service/internal/model"
)

type Repository struct{
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository{
	return &Repository{
		db: db,
	}
}
func (r *Repository) CreateFile(ctx context.Context, fileID string, userID string) error {
    access := model.Access{
        FileID: fileID,
        UserID: userID,
    }

    if err := r.db.WithContext(ctx).Create(&access).Error; err != nil {
        return err
    }

    return nil
}

func (r *Repository) AddAccess(ctx context.Context, fileID string, userID string) error {
    access := model.Access{
        FileID: fileID,
        UserID: userID,
    }

    if err := r.db.WithContext(ctx).Create(&access).Error; err != nil {
        return err
    }

    return nil
}

func (r *Repository) CheckAccess(ctx context.Context, fileID string, userID string) error{
    var access model.Access
    res := r.db.Where("user_id = ? and file_id = ?", userID, fileID).First(&access)
    return res.Error
}

func (r *Repository) GetAccessList(ctx context.Context, fileID string) ([]string, error){
    var accesses []model.Access
    if err := r.db.Model(&model.Access{}).Where("file_id = ?", fileID).Find(&accesses).Error; err != nil{
        return nil, err
    }
    users := make([]string, len(accesses))
    for i, a := range accesses{
        users[i] = a.UserID
    }
    return users, nil
}