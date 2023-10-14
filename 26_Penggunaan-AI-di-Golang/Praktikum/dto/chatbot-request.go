package dto

type LaptopRequest struct {
	Budget  int    `json:"budget" form:"budget"`
	Purpose string `json:"purpose" form:"purpose"`
}
