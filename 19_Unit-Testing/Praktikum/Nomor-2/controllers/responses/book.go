package responses

type BookResponse struct {
	ID        uint      `json:"id"`
	Title 	  string 	`json:"title" form:"title"`
	Writer 	  string 	`json:"writer" form:"writer"`
	Publisher string 	`json:"publisher" form:"publisher"`
}