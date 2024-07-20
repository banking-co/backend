package mysql

import (
	"github.com/go-sql-driver/mysql"
	"rabotyaga-go-backend/mysql/database"
	"rabotyaga-go-backend/structures"
)

func USER_INIT(userId uint) (*structures.User, *mysql.MySQLError) {
	request, err := database.MySQL.Exec("CALL USER_INIT(?)", userId)
	if err != nil {
		return nil, err
	}
	defer request.Close()

	if request.Next() {
		user := new(structures.User)

		err := request.Scan(
			&user.Id,
			&user.UserId,
			&user.Username,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt,
		)

		if err != nil {
			return nil, &mysql.MySQLError{}
		}

		return user, nil
	}

	return nil, &mysql.MySQLError{}
}
