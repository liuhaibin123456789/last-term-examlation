package mysql

import "last-homework/model"

func CreateRoom(maker model.RoomMaker) (err error) {
	return GDB.Model(model.RoomMaker{}).Create(maker).Error
}
