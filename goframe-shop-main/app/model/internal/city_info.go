// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// CityInfo is the golang structure for table city_info.
type CityInfo struct {
	Id        int         `orm:"id,primary" json:"id"`        //
	Name      string      `orm:"name"       json:"name"`      //
	Pid       int         `orm:"pid"        json:"pid"`       //
	Status    int         `orm:"status"     json:"status"`    //
	UpdatedAt *gtime.Time `orm:"updated_at" json:"updatedAt"` //
}
