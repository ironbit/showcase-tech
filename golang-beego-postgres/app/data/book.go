package data

// Book represents the "book" information.
// The relation with Author entity is many to many.
type Book struct {
	ID      int       `json:"id" orm:"pk;column(book_id)"`
	Name    string    `json:"name" orm:"size(128);column(book_name)"`
	Genre   string    `json:"genre" orm:"size(64);column(book_genre)"`
	Authors []*Author `json:"authors,omitempty" orm:"rel(m2m)"`
}

// TableName
func (*Book) TableName() string {
	return "book"
}
