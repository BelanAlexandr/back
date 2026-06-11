package models

type FilesNames struct {
	Id        int     `json:"id"`
	User_Id   int     `json:"user_id"`
	File_Name string  `json:"file_name"`
	Is_closed bool    `json:"is_closed"`
	Total     float64 `json:"total"`
	Total_NDS float64 `json:"total_nds"`
}
