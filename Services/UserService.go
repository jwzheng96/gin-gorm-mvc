package Services

import (
	"fmt"
	"gin-gorm-mvc/Config"
	"gin-gorm-mvc/Models"
)

// GetAllUsers Fetch all user data
func GetAllUsers(user *[]Models.User) (err error) {
	if err = Config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

// CreateUser ... Insert New data
func CreateUser(user *Models.User) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByID ... Fetch only one user by Id
func GetUserByID(user *Models.User, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

// UpdateUser ... Update user
func UpdateUser(user *Models.User, id string) (err error) {
	fmt.Println(user)
	Config.DB.Save(user)
	return nil
}

// DeleteUser ... Delete user
func DeleteUser(user *Models.User, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(user)
	return nil
}
