package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type Meme struct {
	Id      int64  `json:"id"`
	Name    string `json:"name" orm:"type(text)"`
	Url     string `json:"url" orm:"type(text)"`
	Caption string `json:"caption" orm:"type(text)"`
	Created int64  `json:"created"`
	Updated int64  `json:"updated"`
}

func (m *Meme) Insert() error {

	m.Created = time.Now().UnixNano()
	m.Updated = m.Created

	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Meme) Update(fields ...string) error {

	if len(fields) == 0 {
		m.Updated = time.Now().UnixNano()
	} else {
		fields = append(fields, "Updated")
		m.Updated = time.Now().UnixNano()
	}
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil

}

func (m *Meme) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func AllMemes() orm.QuerySeter {
	return orm.NewOrm().QueryTable("meme")
}

func init() {
	orm.RegisterModel(new(Meme))
}
