package dao

import (
	"context"
	"time"

	"github.com/SsrCoder/leetwatcher/model"
)

func (d *Dao) GetAllUsers(ctx context.Context) (users []model.User, err error) {
	err = d.mysql.Find(&users).Error
	return
}

func (d *Dao) UpdateUserLastSubmitTime(ctx context.Context, id int64, submitTime int64) error {
	return d.mysql.Table(model.User{}.TableName()).Where("id = ?", id).UpdateColumn("last_submit_time", time.Unix(submitTime, 0)).Error
}
