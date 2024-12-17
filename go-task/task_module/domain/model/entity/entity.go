package entity

import (
	"github.com/Peterliang233/ddd/go-task/domain/model/vo"
)

type TaskEntity struct {
	Id        int64
	Name      string
	Desc      string
	AllTarget bool
	Target    vo.GroupPeople
}
