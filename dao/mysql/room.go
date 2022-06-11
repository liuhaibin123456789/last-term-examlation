package mysql

import "last-homework/model"

func CreateRoom(maker model.RoomMaker) (err error) {
	return GDB.Model(&model.RoomMaker{}).Create(maker).Error
}

func SelectRoom(userId int64) (roomId int64, err error) {
	err = GDB.Model(&model.RoomMaker{}).Select("room_id").Where("user_id=?", userId).Find(&roomId).Error
	return
}
