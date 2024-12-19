package vo

type GroupPeople struct {
	TargetType int64
	GroupCode  string
	Cnt        int64
}

func NewGroupPeople(targetType int64, groupCode string, cnt int64) GroupPeople {
	return GroupPeople{
		TargetType: targetType,
		GroupCode:  groupCode,
		Cnt:        cnt,
	}
}
