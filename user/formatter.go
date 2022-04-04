package user

type UserFormatter struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

type UserFormatterV1 struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:         user.ID,
		Username:   user.Username,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
	}

	return formatter
}

func FormatUserV1(user User) UserFormatterV1 {
	userFormatter := UserFormatterV1{}
	userFormatter.ID = user.ID
	userFormatter.Username = user.Username
	userFormatter.Name = user.Name
	userFormatter.Occupation = user.Occupation
	userFormatter.Email = user.Email

	return userFormatter
}

func FormatUsers(users []User) []UserFormatterV1 {
	usersFormatter := []UserFormatterV1{}

	for _, user := range users {
		userFormatter := FormatUserV1(user)
		usersFormatter = append(usersFormatter, userFormatter)
	}

	return usersFormatter
}
