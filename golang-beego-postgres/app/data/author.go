package data

// Author represents the "author" information.
// The relation with Book entity is many to many.
type Author struct {
	ID    int     `json:"id" orm:"pk;auto;column(author_id)"`
	Name  string  `json:"name" orm:"size(128);column(author_name)"`
	Books []*Book `json:"books,omitempty" orm:"reverse(many)"`
}

// TableName of the author structure.
func (*Author) TableName() string {
	return "author"
}
