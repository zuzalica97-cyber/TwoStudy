package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UserListSelectRow(conn *pgx.Conn, ctx context.Context, tableName string) (error, []UserModel) {

	sqlQuery := `
	SELECT *
	FROM $1
	ORDER BY id ASC;
	`
	rows, err := conn.Query(ctx, sqlQuery, tableName) //rows это результат запроса, который может содержать несколько строк (если нам нужен только один результат,то используем QueryRow)

	if err != nil {
		return err, nil
	}
	defer rows.Close() //ОБЯЗАТЕЛЬНО!!!! закрываем результат запроса после того, как мы закончили с ним работать

	task := make([]UserModel, 0) //создаем слайс, в который мы будем записывать данные из каждой строки

	for rows.Next() { //делаем цыкл пока есть строки, которые можно прочитать(так нужно делеть если нужно прочитать несколько строк )

		var User UserModel //создаем переменную типа UserModel, в которую мы будем записывать данные из каждой строки

		err := rows.Scan(&User.ID, &User.Name, &User.Money) //сканируем текущую строку и записываем данные в переменные, которые мы создали выше
		if err != nil {
			return err, nil
		}

		task = append(task, User) //добавляем данные в слайс
	}

	return err, task
}

func ProdListSelectRow(conn *pgx.Conn, ctx context.Context) (error, []ProductModel) {

	sqlQuery := `
	SELECT *
	FROM products
	ORDER BY id ASC;
	`
	rows, err := conn.Query(ctx, sqlQuery) //rows это результат запроса, который может содержать несколько строк (если нам нужен только один результат,то используем QueryRow)

	if err != nil {
		return err, nil
	}
	defer rows.Close() //ОБЯЗАТЕЛЬНО!!!! закрываем результат запроса после того, как мы закончили с ним работать

	task := make([]ProductModel, 0) //создаем слайс, в который мы будем записывать данные из каждой строки

	for rows.Next() { //делаем цыкл пока есть строки, которые можно прочитать(так нужно делеть если нужно прочитать несколько строк )

		var Prod ProductModel //создаем переменную типа ProductModel, в которую мы будем записывать данные из каждой строки

		err := rows.Scan(&Prod.ID, &Prod.Name, &Prod.Description, &Prod.Cost, &Prod.Amount) //сканируем текущую строку и записываем данные в переменные, которые мы создали выше
		if err != nil {
			return err, nil
		}

		task = append(task, Prod) //добавляем данные в слайс
	}

	return err, task
}

func UserSelectRow(conn *pgx.Conn, ctx context.Context, idU int) (error, UserModel) {

	sqlQuery := `
	SELECT *
	FROM users
	WHERE id = $1;
	`
	row := conn.QueryRow(ctx, sqlQuery, idU) //QweryRow это результат запроса, который может содержать только одну строку (если нам нужно несколько строк, то используем Query)

	var User UserModel //создаем переменную типа UserModel, в которую мы будем записывать данные из cтроки

	err := row.Scan(&User.ID, &User.Name, &User.Money) //сканируем текущую строку и записываем данные в переменные, которые мы создали выше

	if err != nil {
		return err, User
	}

	return err, User
}

func ProdSelectRow(conn *pgx.Conn, ctx context.Context, idU int) (error, ProductModel) {

	sqlQuery := `
	SELECT *
	FROM products
	WHERE id = $1;
	`
	row := conn.QueryRow(ctx, sqlQuery, idU) //QweryRow это результат запроса, который может содержать только одну строку (если нам нужно несколько строк, то используем Query)

	var Prod ProductModel //создаем переменную типа UserModel, в которую мы будем записывать данные из cтроки

	err := row.Scan(&Prod.ID, &Prod.Name, &Prod.Description, &Prod.Cost, &Prod.Amount) //сканируем текущую строку и записываем данные в переменные, которые мы создали выше

	if err != nil {
		return err, Prod
	}

	return err, Prod
}
