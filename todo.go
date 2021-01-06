package todo

type List struct {
	Id          string `json:"id" db:"id"`
	Title       string `json:"title" binding:"required" db:"title"`
	Description string `json:"description" db:"description"`
}

type UsersList struct {
	Id     int `db:"id"`
	UserId int `db:"user_id"`
	ListId int `db:"list_id"`
}

type Item struct {
	Id          string `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}

type ListsItem struct {
	Id     int `db:"id"`
	ListId int `db:"list_id"`
	ItemId int `db:"item_id"`
}
