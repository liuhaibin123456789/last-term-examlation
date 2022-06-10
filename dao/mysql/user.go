package mysql

import "last-homework/model"

func InsertUser(u *model.User) (err error) {
	err = GDB.Model(&model.User{}).Create(u).Error
	if err != nil {
		return err
	}

	return GDB.Model(&model.User{}).Where("phone=?", u.Phone).Find(u).Error
}

func SelectUserPwd(phone string) (user *model.User, err error) {
	user = new(model.User)
	err = GDB.Model(&model.User{}).Where("phone=?", phone).Find(user).Error
	return
}

func SelectUserName(userId int64) (userName string, err error) {
	err = GDB.Model(&model.User{}).Select("user_name").Where("user_id=?", userId).Find(&userName).Error
	return
}

func SelectUser(userId int64) (user *model.User, err error) {
	user = new(model.User)
	err = GDB.Model(&model.User{}).Where("user_id=?", userId).Find(user).Error
	return
}
