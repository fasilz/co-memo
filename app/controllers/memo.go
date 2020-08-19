package controllers

import (
	"github.com/fasilz/co-memo/app/models"
	"github.com/revel/revel"
)

type Memo struct {
	App
}

func (m *Memo) Memos() revel.Result {

	if m.CurrentUser == nil {
		m.Flash.Error("Please log in first")
		return m.Redirect(User.Login)
	}

	var memos []models.Memo

	err := m.Txn.Preload("To").Where("ID=?", m.CurrentUser.UserID).Preload("ReadReceipt", "ID=? NOT IN userid", m.CurrentUser.UserID).Find(&memos).Error
	if err != nil {
		return nil
	}

	return m.Render(memos)

}
