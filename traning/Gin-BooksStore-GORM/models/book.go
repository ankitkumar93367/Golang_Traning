type Book struct {
	ID uint			`json:"id" gorm:"promary_key`
	Title string	`json:"title"`
	Author string	`json:"author"`
}