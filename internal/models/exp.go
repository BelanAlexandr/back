package models

type Expert struct {
	Id          int    `json:"id"`
	Name        string `json:"name" validate:"required,custom_name"`
	Second_Name string `json:"second_name" validate:"required,custom_name"`
	Patronymic  string `json:"patronymic,omitempty" validate:"omitempty,custom_name"`
}

type ExpListItem struct {
	Id        int      `json:"id"`
	Data_Post string   `json:"data_post"`
	Is_Closed bool     `json:"is_closed"`
	Fab       string   `json:"fab"`
	Experts   []Expert `json:"experts"`
}

type Exp struct {
	Id           int    `json:"id"`
	Creator_id   int    `json:"creator_id" validate:"required"`
	Data_Post    string `json:"data_post" validate:"required"`
	Fab          string `json:"fab" validate:"required"`
	Adm_Material int    `json:"adm_material" validate:"required"`
	Nom_Statyi   string `json:"nom_statyi" validate:"required"`
	Vid_Exp      int    `json:"vid_exp" validate:"required"`
	Organ        string `json:"organ" validate:"required"`
	Name_Organ   string `json:"name_organ" validate:"required"`

	Name_Naznch        string  `json:"name_naznch" validate:"required,custom_name"`
	Second_Name_Naznch string  `json:"second_name_naznch" validate:"required,custom_name"`
	Patronymic_Naznch  *string `json:"patronymic_naznch" validate:"omitempty,custom_name"`

	Experts []Expert `json:"experts" validate:"required,dive"`

	Question_Count int `json:"question_count" validate:"required"`
	Object_Count   int `json:"object_count" validate:"required"`

	Srok_Exp      *int    `json:"srok_exp"`
	Stop_Date     *string `json:"stop_date"`
	Stop_Reason   *string `json:"stop_reason"`
	Resuming_Date *string `json:"resuming_date"`
	Srok_Resuming *int    `json:"srok_resuming"`

	End_Date          *string  `json:"end_date" validate:"required_if=Is_Closed true"`
	Day_Count         *int     `json:"day_count" validate:"required_if=Is_Closed true"`
	Exp_Day_Count     *int     `json:"exp_day_count" validate:"required_if=Is_Closed true"`
	Cat_Vivod         *int     `json:"cat_vivod" validate:"required_if=Is_Closed true"`
	Possible_Vivod    *int     `json:"possible_vivod" validate:"required_if=Is_Closed true"`
	Impossible_Vivod  *int     `json:"impossible_vivod" validate:"required_if=Is_Closed true"`
	Hour_Count        *int     `json:"hour_count" validate:"required_if=Is_Closed true"`
	Exp_Res_Id        *int     `json:"exp_res_id" validate:"required_if=Is_Closed true"`
	Expert_Cost       *float64 `json:"expert_cost"`
	Material_Cost     *float64 `json:"material_cost"`
	Exploitation_Cost *float64 `json:"exploitation_cost"`
	Full_Cost         *float64 `json:"full_cost"`
	Full_Cost_Nds     *float64 `json:"full_cost_nds"`
	Descrip           *string  `json:"descrip"`

	Is_Closed bool `json:"is_closed"`

	Stat_Id     *int `json:"stat_id" validate:"required"`
	Category_Id *int `json:"category_id" validate:"required"`
	Region_Id   *int `json:"region_id" validate:"required"`
	Iz_Nix_Id   *int `json:"iz_nix_id" validate:"required"`
	Diff_Cat_Id *int `json:"diff_cat_id" validate:"required"`
}
