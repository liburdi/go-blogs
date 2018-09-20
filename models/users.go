package models

// UsersModel 模型
type UsersModel struct {
	ModelProperty
}

// UsersProperty 属性
type UsersProperty struct {
	User_id          int64
	Username         string
	Tel              string
	Password         string
	Taglist          string
	Headimg          string
	Position         string
	Company          string
	Selfintroduction string
	Homepage         string
	T_headimg        string
}

func (users *UsersModel) SetTableName() *UsersModel {
	users.ModelProperty.TableName = "php41_users"
	return users
}
