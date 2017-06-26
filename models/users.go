package models

type User struct {
	Username  string
	FirstName string
	LastName  string
}

func AllUsers() ([]User, error) {
	rows, err := DB.Query("SELECT username, first_name, last_name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.Username, &user.FirstName, &user.LastName)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
