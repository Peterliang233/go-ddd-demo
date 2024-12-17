package do

type TaskDo struct {
	Id   int64  `gorm:"id" json:"id"`
	Name string `gorm:"name" json:"name"`
	Desc string `gorm:"desc" json:"desc"`
}
