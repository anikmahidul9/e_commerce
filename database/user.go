package database

type User struct {
	ID          uint   `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

var users []User

func (user User) Store() (User, error) {
	if user.ID != 0 {
		return user, nil
	}
	user.ID = uint(len(users) + 1)
	users = append(users, user)
	return user, nil
}

func Find(email, password string) (*User, error) {
	for _, user := range users {
		if user.Email == email && user.Password == password {
			return &user, nil
		}
	}
	return nil, nil
}