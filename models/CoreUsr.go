package models

import (
	"fmt"
	"time"
)

type CoreUsr struct {
	Oid       string    `gorm:"primaryKey"`
	Createdon time.Time `gorm:"autoUpdateTime:milli"`
	ABalance  float64   `gorm:"column:a_balance;'"`
}

func (tx *CoreUsr) GetString() string {
	return fmt.Sprintf("用户名为%s,创建时间为%s,余额为%f", tx.Oid, tx.Createdon.Format("2006-01-02 15:04:05"), tx.ABalance)
}
