package repository

import (
	"context"
	"fmt"

	"github.com/BelanAlexandr/back/internal/models"
)

func UpdateExpRepo(exp models.Exp, closed bool) error {
	if closed {
		exp.Is_Closed = true
	} else {
		exp.Is_Closed = false
	}
	query := `
		UPDATE electronic_journal 
		SET 
			creator_id = $1, data_post = $2, fab = $3, adm_material = $4, nom_statyi = $5, 
			vid_exp = $6, organ = $7, name_organ = $8, name_naznch = $9, second_name_naznch = $10, 
			patronymic_naznch = $11, name_exp = $12, second_name_exp = $13, patronymic_exp = $14, 
			question_count = $15, object_count = $16, srok_exp = $17, stop_date = $18, 
			stop_reason = $19, resuming_date = $20, srok_resuming = $21, end_date = $22, 
			day_count = $23, exp_day_count = $24, cat_vivod = $25, possible_vivod = $26, 
			impossible_vivod = $27, hour_count = $28, expert_cost = $29, material_cost = $30, 
			exploitation_cost = $31, full_cost = $32, full_cost_nds = $33, descrip = $34, 
			is_closed = $35, stat_id = $36, category_id = $37, region_id = $38, 
			iz_nix_id = $39, diff_cat_id = $40, exp_res_id = $41
		WHERE id = $42;`

	result, err := db.ExecContext(
		context.Background(),
		query,
		exp.Creator_id,         // $1
		exp.Data_Post,          // $2
		exp.Fab,                // $3
		exp.Adm_Material,       // $4
		exp.Nom_Statyi,         // $5
		exp.Vid_Exp,            // $6
		exp.Organ,              // $7
		exp.Name_Organ,         // $8
		exp.Name_Naznch,        // $9
		exp.Second_Name_Naznch, // $10
		exp.Patronymic_Naznch,  // $11
		exp.Name_Exp,           // $12
		exp.Second_Name_Exp,    // $13
		exp.Patronymic_Exp,     // $14
		exp.Question_Count,     // $15
		exp.Object_Count,       // $16
		exp.Srok_Exp,           // $17
		exp.Stop_Date,          // $18
		exp.Stop_Reason,        // $19
		exp.Resuming_Date,      // $20
		exp.Srok_Resuming,      // $21
		exp.End_Date,           // $22
		exp.Day_Count,          // $23
		exp.Exp_Day_Count,      // $24
		exp.Cat_Vivod,          // $25
		exp.Possible_Vivod,     // $26
		exp.Impossible_Vivod,   // $27
		exp.Hour_Count,         // $28
		exp.Expert_Cost,        // $29
		exp.Material_Cost,      // $30
		exp.Exploitation_Cost,  // $31
		exp.Full_Cost,          // $32
		exp.Full_Cost_Nds,      // $33
		exp.Descrip,            // $34
		exp.Is_Closed,          // $35
		exp.Stat_Id,            // $36
		exp.Category_Id,        // $37
		exp.Region_Id,          // $38
		exp.Iz_Nix_Id,          // $39
		exp.Diff_Cat_Id,        // $40
		exp.Exp_Res_Id,         // $41
		exp.Id,                 // $42 (ID для условия WHERE)
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("экспертиза с id %d не найдена для обновления", exp.Id)
	}

	return nil
}
