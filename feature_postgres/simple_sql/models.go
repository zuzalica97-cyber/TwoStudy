package simple_sql

type UserModel struct { //структура, которая описывает модель данных, которую мы будем использовать для работы с базой данных
	ID    int
	Name  string
	Money int
}

type ProductModel struct { //структура, которая описывает модель данных, которую мы будем использовать для работы с базой данных
	ID          int
	Name        string
	Description string
	Cost        int
	Amount      int
}
