package entity

type UserEntity struct {
	BaseEntity
	Username    string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
}

// TableName 指定模型的表名称
func (UserEntity) TableName() string {
	return TableUser
}
