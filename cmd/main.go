package main

import "zavad/internal/app"

func main() {
	app.New().Run()
}

// func createAccessToken(user *entities.User) (string, error) {
// 	token := jwt.New(jwt.SigningMethodHS256)
// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["id"] = user.ID
// 	claims["user_token"] = user.UserToken
// 	claims["name"] = user.Name
// 	claims["phone"] = user.Phone
// 	claims["login"] = user.Login
// 	claims["email"] = user.Email
// 	claims["avatar"] = user.Avatar
// 	claims["exp"] = time.Now().Add(time.Hour * 24 * 365).Unix()

// 	tokenString, err := token.SignedString([]byte("lio2UbeLoKlYuJ7LDR+kSxUPQDHbaekq"))
// 	if err != nil {
// 		return "", err
// 	}

// 	return tokenString, nil
// }
