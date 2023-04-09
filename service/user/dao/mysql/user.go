package mysql

import "user/model"

func GetUserInfo(uname string) (model.ParamUsers, error) {
	var user model.ParamUsers
	err := DB.Where("name =?", uname).Model(model.User{}).First(&user).Error
	return user, err
}

func UpdateUserName(oldName, newName string) error {
	err := DB.Model(model.User{}).Where("name =?", oldName).Update("name", newName).Error
	return err
}

func SaveUserAvatar(userName, avatarUrl string) error {
	return DB.Model(model.User{}).Where("name = ?", userName).Update("avatar_url", avatarUrl).Error
}

func UpdateAuth(name, realName, idCard string) error {
	err := DB.Model(model.User{}).Where("name =?", name).Updates(model.User{RealName: realName, IdCard: idCard}).Error
	return err
}
