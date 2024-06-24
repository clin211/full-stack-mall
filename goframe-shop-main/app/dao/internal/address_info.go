// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// AddressInfoDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type AddressInfoDao struct {
	gmvc.M                     // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB             // DB is the raw underlying database management object.
	Table   string             // Table is the table name of the DAO.
	Columns addressInfoColumns // Columns contains all the columns of Table that for convenient usage.
}

// AddressInfoColumns defines and stores column names for table address_info.
type addressInfoColumns struct {
	Id        string //
	Name      string //
	Pid       string //
	Status    string //
	UpdatedAt string //
}

func NewAddressInfoDao() *AddressInfoDao {
	return &AddressInfoDao{
		M:     g.DB("default").Model("address_info").Safe(),
		DB:    g.DB("default"),
		Table: "address_info",
		Columns: addressInfoColumns{
			Id:        "id",
			Name:      "name",
			Pid:       "pid",
			Status:    "status",
			UpdatedAt: "updated_at",
		},
	}
}