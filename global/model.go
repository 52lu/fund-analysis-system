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
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
