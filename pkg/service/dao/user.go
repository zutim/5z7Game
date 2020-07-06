package dao

import (
	"5z7Game/pkg/service/entity"
	"github.com/jinzhu/gorm"
)

type userDao struct {
	BaseDao
}

func User(db *gorm.DB) *userDao {
	return &userDao{BaseDao{db: db}}
}

// GetByUsername 根据用户名获取记录
func (dao *userDao) GetByUsername(username string) (*entity.UserEntity, error) {
	query := dao.db.Table(entity.TableUser).
		Where("username = ?", username).
		Where(entity.SoftDeleteCondition)

	user := new(entity.UserEntity)
	if err := query.First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
