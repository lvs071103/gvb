package ctype

import "encoding/json"

type Role int

const (
	PermissionAdmin       Role = 1 // 管理员
	PermissionUser        Role = 2 // 普通人
	PermissionVisitor     Role = 3 // 游客
	PermissionDisableUser Role = 4 // 被禁用的用户
)

func (r Role) MarshalJson() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Role) String() string {
	var str string
	switch r {
	case PermissionAdmin:
		str = "管理员"
	case PermissionUser:
		str = "用户"
	case PermissionVisitor:
		str = "游客"
	case PermissionDisableUser:
		str = "被禁言的用户"
	default:
		str = "其他"
	}
	return str
}
