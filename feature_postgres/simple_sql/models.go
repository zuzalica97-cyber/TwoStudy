package simple_sql

type UserModel struct { //структура, которая описывает модель данных, которую мы будем использовать для работы с базой данных
	ID    int
	Name  string
	Money int
	IDU   int
}

type ProductModel struct { //структура, которая описывает модель данных, которую мы будем использовать для работы с базой данных
	ID                 int
	ProductName        string
	ProductDescription string
	Cost               int
	Amount             int
	IDP                int
}

type DataBaseModel struct { //структура, которая описывает модель данных, которую мы будем использовать для работы с базой данных
	ID        int
	UserID    int
	ProductID int
	Amount    int
	TotalCost int
	Canceled  bool
	BaseId    int
}
