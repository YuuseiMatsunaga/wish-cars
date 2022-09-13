package repositories

import (
	"app/database"
	"app/entities"
	// "time"
)

func GetUsers() []entities.User {
	users := []entities.User{}
	database.DB.Find(&users)
	return users
}

func GetUser(id int) entities.User {
	user := entities.User{}
	database.DB.First(&user, id)

	return user
}

// func CreateUser(params echo.Context) (*entities.User, error) {
// 	user := new(entities.User)
// 	if err := params.Bind(&user); err != nil {
// 		return nil, err
// 	}
// 	user.CreatedAt = time.Now()
// 	database.DB.Create(&user)
// 	return user, nil
// }

// func UpdateUser(params echo.Context) (entities.User, error) {
// 	user := entities.User{}
// 	database.DB.First(&user, params.Param("id"))
// 	if err := params.Bind(&user); err != nil {
// 		return entities.User{}, err
// 	}
// 	database.DB.Save(&user)
// 	return user, nil
// }

func DeleteUser(id int) bool {
	if err := database.DB.Delete(&entities.User{}, id); err != nil {
		return false
	}
	return true
}
