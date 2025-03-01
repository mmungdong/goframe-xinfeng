// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"goframe-xinfeng/internal/dao/internal"
)

// internalUserCouponInfoDao is internal type for wrapping internal DAO implements.
type internalUserCouponInfoDao = *internal.UserCouponInfoDao

// userCouponInfoDao is the data access object for table user_coupon_info.
// You can define custom methods on it to extend its functionality as you wish.
type userCouponInfoDao struct {
	internalUserCouponInfoDao
}

var (
	// UserCouponInfo is globally public accessible object for table user_coupon_info operations.
	UserCouponInfo = userCouponInfoDao{
		internal.NewUserCouponInfoDao(),
	}
)

// Fill with you ideas below.
