package rgm

type SalesList struct {
	Id     int     `json:"id" db:"id"`
	Title  string  `json:"title" db:"title" binding:"required"`
	Price  float64 `json:"price" db:"price"`
	Amount int     `json:"amount" db:"amount"`
}

type UsersList struct {
	Id     int
	UserId int
	ListID int
}

type SalesRepo struct {
	Id     int
	Title  string  `json:"title" db:"title" binding:"required"`
	Price  float64 `json:"price" db:"price"`
	Amount int     `json:"amount" db:"amount"`
	Total  int     `json:"total" db:"total"`
}

type SalesPdf struct {
	Id     int
	Saller string
	Title  string  `json:"title" db:"title" binding:"required"`
	Price  float64 `json:"price" db:"price"`
	Amount int     `json:"amount" db:"amount"`
	Total  int     `json:"total" db:"total"`
}
