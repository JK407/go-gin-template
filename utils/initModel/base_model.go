package initModel

import (
	"fmt"
	"gin-template/utils/ztime"
	"time"

	"gorm.io/gorm"
)

type ID struct {
	ID uint64 `gorm:"primarykey;column:id" db:"id" json:"id"`
}

func (o ID) GetId() uint64 {
	return o.ID
}

// Timestamps 时间兼容现有数据库
type Timestamps struct {
	CreateTime int64 `gorm:"autoCreateTime;column:create_time" db:"create_time" json:"create_time"`
	UpdateTime int64 `gorm:"autoUpdateTime;column:update_time" db:"update_time" json:"update_time"`
}

type SoftDeletes struct {
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at" db:"deleted_at" json:"deleted_at"`
}

func SetPage(PageReq, PageSizeReq int) (page, pageSize, offset int) {
	page = 1
	pageSize = 20
	if PageReq > 1 {
		page = PageReq
	}
	if PageSizeReq > 1 {
		pageSize = PageSizeReq
	}
	offset = (page - 1) * pageSize
	return page, pageSize, offset
}

type HTime time.Time

func (t *HTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format(ztime.DateTimeFormat))), nil
}

type HDate time.Time

func (t *HDate) MarshalJSON() ([]byte, error) {
	tDate := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tDate.Format(ztime.DateFormat))), nil
}

type Json []byte
