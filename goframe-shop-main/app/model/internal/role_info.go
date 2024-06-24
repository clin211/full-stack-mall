// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// RoleInfo is the golang structure for table role_info.
type RoleInfo struct {
	Id        int         `orm:"id,primary" json:"id"`        //
	Name      string      `orm:"name"       json:"name"`      // 角色名称
	Desc      string      `orm:"desc"       json:"desc"`      // 描述
	CreatedAt *gtime.Time `orm:"created_at" json:"createdAt"` //
	UpdatedAt *gtime.Time `orm:"updated_at" json:"updatedAt"` //
	DeletedAt *gtime.Time `orm:"deleted_at" json:"deletedAt"` //
}