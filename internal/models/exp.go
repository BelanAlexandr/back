package models

type Exp struct {
	Id                 int     `json:"id"`
	Creator_id         int     `json:"creator_id" validate:"required"`
	Data_Post          string  `json:"data_post" validate:"required"`
	Fab                string  `json:"fab" validate:"required"`
	Adm_Material       int     `json:"adm_material" validate:"required"`
	Nom_Statyi         string  `json:"nom_statyi" validate:"required"`
	Vid_Exp            int     `json:"vid_exp" validate:"required"`
	Organ              string  `json:"organ" validate:"required"`
	Name_Organ         string  `json:"name_organ" validate:"required"`
	Name_Naznch        string  `json:"name_naznch" validate:"required"`
	Second_Name_Naznch string  `json:"second_name_naznch" validate:"required"`
	Patronymic_Naznch  *string `json:"patronymic_naznch"`
	Name_Exp           string  `json:"name_exp" validate:"required"`
	Second_Name_Exp    string  `json:"second_name_exp" validate:"required"`
	Patronymic_Exp     *string `json:"patronymic_exp"`
	Question_Count     int     `json:"question_count" validate:"required"`
	Object_Count       int     `json:"object_count" validate:"required"`

	Srok_Exp          *int     `json:"srok_exp"`
	Stop_Date         *string  `json:"stop_date"`
	Stop_Reason       *string  `json:"stop_reason"`
	Resuming_Date     *string  `json:"resuming_date"`
	Srok_Resuming     *int     `json:"srok_resuming"`
	End_Date          *string  `json:"end_date"`
	Day_Count         *int     `json:"day_count"`
	Exp_Day_Count     *int     `json:"exp_day_count"`
	Cat_Vivod         *int     `json:"cat_vivod"`
	Possible_Vivod    *int     `json:"possible_vivod"`
	Impossible_Vivod  *int     `json:"impossible_vivod"`
	Hour_Count        *int     `json:"hour_count"`
	Expert_Cost       *float64 `json:"expert_cost"`
	Material_Cost     *float64 `json:"material_cost"`
	Exploitation_Cost *float64 `json:"exploitation_cost"`
	Full_Cost         *float64 `json:"full_cost"`
	Full_Cost_Nds     *float64 `json:"full_cost_nds"`
	Descrip           *string  `json:"descrip"`

	Is_Closed bool `json:"is_closed"`

	Stat_Id     *int `json:"stat_id"`
	Category_Id *int `json:"category_id"`
	Region_Id   *int `json:"region_id"`
	Iz_Nix_Id   *int `json:"iz_nix_id"`
	Diff_Cat_Id *int `json:"diff_cat_id"`
	Exp_Res_Id  *int `json:"exp_res_id"`
}
type AddExp struct {
	Id                 int     `json:"id"`
	Creator_id         int     `json:"creator_id" validate:"required"`
	Data_Post          string  `json:"data_post" validate:"required"`
	Fab                string  `json:"fab" validate:"required"`
	Adm_Material       int     `json:"adm_material" validate:"required"`
	Nom_Statyi         string  `json:"nom_statyi" validate:"required"`
	Vid_Exp            int     `json:"vid_exp" validate:"required"`
	Organ              string  `json:"organ" validate:"required"`
	Name_Organ         string  `json:"name_organ" validate:"required"`
	Name_Naznch        string  `json:"name_naznch" validate:"required"`
	Second_Name_Naznch string  `json:"second_name_naznch" validate:"required"`
	Patronymic_Naznch  *string `json:"patronymic_naznch"`
	Name_Exp           string  `json:"name_exp" validate:"required"`
	Second_Name_Exp    string  `json:"second_name_exp" validate:"required"`
	Patronymic_Exp     *string `json:"patronymic_exp"`
	Question_Count     int     `json:"question_count" validate:"required"`
	Object_Count       int     `json:"object_count" validate:"required"`
	Srok_Exp           *int    `json:"srok_exp"`
	Stop_Date          *string `json:"stop_date"`
	Stop_Reason        *string `json:"stop_reason"`
	Resuming_Date      *string `json:"resuming_date"`
	Srok_Resuming      *int    `json:"srok_resuming"`
	Stat_Id            *int    `json:"stat_id" validate:"required"`
	Category_Id        *int    `json:"category_id" validate:"required"`
	Region_Id          *int    `json:"region_id" validate:"required"`
	Iz_Nix_Id          *int    `json:"iz_nix_id" validate:"required"`
	Diff_Cat_Id        *int    `json:"diff_cat_id" validate:"required"`
}
type CloseExp struct {
	Id                int      `json:"id"`
	Creator_id        int      `json:"creator_id" validate:"required"`
	End_Date          *string  `json:"end_date" validate:"required"`
	Day_Count         *int     `json:"day_count" validate:"required"`
	Exp_Day_Count     *int     `json:"exp_day_count" validate:"required"`
	Cat_Vivod         *int     `json:"cat_vivod" validate:"required"`
	Possible_Vivod    *int     `json:"possible_vivod" validate:"required"`
	Impossible_Vivod  *int     `json:"impossible_vivod" validate:"required"`
	Hour_Count        *int     `json:"hour_count" validate:"required"`
	Expert_Cost       *float64 `json:"expert_cost"`
	Material_Cost     *float64 `json:"material_cost"`
	Exploitation_Cost *float64 `json:"exploitation_cost"`
	Full_Cost         *float64 `json:"full_cost"`
	Full_Cost_Nds     *float64 `json:"full_cost_nds"`
	Descrip           *string  `json:"descrip"`

	Is_Closed  bool `json:"is_closed"`
	Exp_Res_Id *int `json:"exp_res_id" validate:"required"`
}
