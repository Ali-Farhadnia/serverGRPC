package book

type Book struct {
	ID        string `json="id"`
	Name      string `json="name"`
	Author    string `json="author"`
	Pagecount int32  `json="pagecount"`
	Inventory int64  `json="inventory"`
}
