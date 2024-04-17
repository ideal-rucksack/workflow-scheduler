package entities

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

// JSONFormatTime 是一个自定义时间类型，将 time.Time 格式化为 JSON 格式
type JSONFormatTime time.Time

type GlobalEntity struct {
	Id       *int64          `json:"id" gorm:"primary_key"`
	CreateAt *JSONFormatTime `json:"created_at" gorm:"autoCreateTime"`
	ModifyAt *JSONFormatTime `json:"modified_at" gorm:"autoUpdateTime"`
}

// TimeFormat 是 JSONFormatTime 的时间格式
const TimeFormat = "2006-01-02 15:04:05"

// MarshalJSON 实现了 JSONFormatTime 的 MarshalJSON 方法
func (t *JSONFormatTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", time.Time(*t).Format(TimeFormat))
	return []byte(formatted), nil
}

// UnmarshalJSON 实现了 JSONFormatTime 的 UnmarshalJSON 方法
func (t *JSONFormatTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse(TimeFormat, timeStr)
	*t = JSONFormatTime(t1)
	return err
}

// Scan UnmarshalJSON 实现了 JSONFormatTime 的 UnmarshalJSON 方法
func (t *JSONFormatTime) Scan(value any) error {
	switch valueType := value.(type) {
	case time.Time:
		*t = JSONFormatTime(valueType)
		break
	default:
		return fmt.Errorf("cannot scan type %T into JSONFormatTime", value)
	}
	return nil
}

// Value 实现了 JSONFormatTime 的 Value 方法
func (t *JSONFormatTime) Value() (driver.Value, error) {
	current := time.Time(*t)
	return current.Format(TimeFormat), nil
}
