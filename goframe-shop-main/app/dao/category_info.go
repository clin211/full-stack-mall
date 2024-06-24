// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"shop/app/dao/internal"
)

// categoryInfoDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type categoryInfoDao struct {
	*internal.CategoryInfoDao
}

var (
	// CategoryInfo is globally public accessible object for table category_info operations.
	CategoryInfo categoryInfoDao
)

func init() {
	CategoryInfo = categoryInfoDao{
		internal.NewCategoryInfoDao(),
	}
}

// Fill with you ideas below.
