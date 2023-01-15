package model

import "gorm.io/gorm"

type Cut struct {
	gorm.Model
	// Name string
	Type string
	File []byte
	Text string
}

func (c *Cut) Creat() error {
	return db.Omit("id").Create(c).Error
}

func (c *Cut) Updata() error {
	return db.Updates(c).Error
}
func (c *Cut) Save() error {
	return db.Save(c).Error
}

func (c *Cut) Delete() error {
	return db.Delete(c).Error
}

func (c *Cut) QueryById() error {
	return db.First(c, c.ID).Error
}

// func (t *Cut) WebList() (int, int, []Cut) {
// 	q := db.Model(t)
// 	var Count int64
// 	q.Count(&Count)
// 	q.Order("id desc").Limit(p.PageSize).Offset((p.PageNum - 1) * p.PageSize).Find(&list)
// 	return int(Count)
// }
