package task

import (
	"context"

	"gorm.io/gorm"

	"github.com/CocaineCong/todolist-ddd/domain/task/entity"
	"github.com/CocaineCong/todolist-ddd/domain/task/repository"
	"github.com/CocaineCong/todolist-ddd/interfaces/types"
)

type RepositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repository.Task {
	return &RepositoryImpl{db: db}
}

func (r *RepositoryImpl) withDB(tx *gorm.DB) *gorm.DB {
	if tx != nil {
		return tx
	}
	return r.db
}

// Paginate returns a GORM scope that applies LIMIT/OFFSET with safe
// defaults when the caller did not provide a page size.
func Paginate(p types.Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := p.Page
		if page < 1 {
			page = 1
		}
		size := p.PageSize
		if size < 1 {
			size = 10
		}
		offset := (page - 1) * size
		return db.Offset(offset).Limit(size)
	}
}

func (r *RepositoryImpl) CreateTask(ctx context.Context, tx *gorm.DB, in *entity.Task) (*entity.Task, error) {
	po := Entity2PO(in)
	db := r.withDB(tx).WithContext(ctx)
	if err := db.Model(&Task{}).Create(&po).Error; err != nil {
		return nil, err
	}
	return PO2Entity(po), nil
}

func (r *RepositoryImpl) FindTaskByTid(ctx context.Context, tx *gorm.DB, taskId, userId uint) (*entity.Task, error) {
	task := &entity.Task{}
	err := r.withDB(tx).WithContext(ctx).Model(&Task{}).
		Where("id = ? AND uid = ?", taskId, userId).
		Find(&task).Error
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (r *RepositoryImpl) ListTaskByUid(ctx context.Context, tx *gorm.DB, uid uint, p types.Pagination) ([]*entity.Task, int64, error) {
	var tasks []*entity.Task
	var count int64
	db := r.withDB(tx).WithContext(ctx).Model(&Task{}).Where("uid = ?", uid)
	if err := db.Count(&count).Error; err != nil {
		return nil, count, err
	}
	if err := db.Scopes(Paginate(p)).Find(&tasks).Error; err != nil {
		return nil, count, err
	}
	return tasks, count, nil
}

func (r *RepositoryImpl) UpdateTask(ctx context.Context, tx *gorm.DB, task *entity.Task) error {
	update := make(map[string]any)
	if task.Content != "" {
		update["content"] = task.Content
	}
	if task.Status != 0 {
		update["status"] = task.Status
	}
	if task.Title != "" {
		update["title"] = task.Title
	}
	if len(update) == 0 {
		return nil
	}
	return r.withDB(tx).WithContext(ctx).Model(&Task{}).
		Where("id = ? AND uid = ?", task.Id, task.Uid).
		Updates(update).Error
}

func (r *RepositoryImpl) SearchTask(ctx context.Context, tx *gorm.DB, uid uint, keyword string, p types.Pagination) ([]*entity.Task, int64, error) {
	var tasks []*entity.Task
	var count int64
	db := r.withDB(tx).WithContext(ctx).Model(&Task{}).
		Where("uid = ?", uid).
		Where("title LIKE ? OR content LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	if err := db.Count(&count).Error; err != nil {
		return nil, count, err
	}
	if err := db.Scopes(Paginate(p)).Find(&tasks).Error; err != nil {
		return nil, count, err
	}
	return tasks, count, nil
}

func (r *RepositoryImpl) DeleteTask(ctx context.Context, tx *gorm.DB, uid, tid uint) error {
	return r.withDB(tx).WithContext(ctx).Model(&Task{}).
		Where("id = ? AND uid = ?", tid, uid).
		Delete(&Task{}).Error
}

func (r *RepositoryImpl) BumpUserName(ctx context.Context, tx *gorm.DB, uid uint, newName string) error {
	return r.withDB(tx).WithContext(ctx).Model(&Task{}).
		Where("uid = ?", uid).
		Update("user_name", newName).Error
}