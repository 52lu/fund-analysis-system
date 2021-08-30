/**
 * @Description 重写
 **/
package global

import (
	"gorm.io/gorm"
	"time"
)

// 重写gorm.Model
type BaseModel struct {
	ID        uint           `json:"id" gorm:"primarykey;comment:主键"`
	CreatedAt time.Time      `json:"createdAt"  gorm:"comment:创建时间"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"comment:更新时间"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间" `
}
