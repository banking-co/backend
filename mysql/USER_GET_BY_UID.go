package mysql

import (
	"github.com/go-sql-driver/mysql"
	"rabotyaga-go-backend/mysql/database"
	"rabotyaga-go-backend/structures"
)

func USER_GET_BY_UID(userId uint) (*structures.ResponseUserGet, *mysql.MySQLError) {
	request, err := database.MySQL.Exec("CALL USER_GET_BY_UID(?)", userId)
	if err != nil {
		return nil, err
	}
	defer request.Close()

	if request.Next() {
		user := new(structures.ResponseUserGet)

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
