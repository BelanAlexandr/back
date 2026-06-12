package repository

import (
	"context"

	"github.com/BelanAlexandr/back/internal/models"
)

func AddExpRepo(exp models.Exp) error {
	query := `
		INSERT INTO electronic_journal (
			creator_id, data_post, fab, adm_material, nom_statyi, vid_exp, organ, name_organ,
			name_naznch, second_name_naznch, patronymic_naznch, name_exp, second_name_exp, patronymic_exp,
			question_count, object_count, srok_exp, stop_date, stop_reason, resuming_date,
			srok_resuming, end_date, day_count, exp_day_count, cat_vivod, possible_vivod,
			impossible_vivod, hour_count, expert_cost, material_cost, exploitation_cost, full_cost,
			full_cost_nds, descrip, is_closed, stat_id, category_id, region_id, iz_nix_id,
			diff_cat_id, exp_res_id
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20,
			$21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38,
			$39, $40, $41
		);`

	_, err := db.ExecContext(
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
		exp.Patronymic_Naznch,  // $11 (указатель)
		exp.Name_Exp,           // $12
		exp.Second_Name_Exp,    // $13
		exp.Patronymic_Exp,     // $14 (указатель)
		exp.Question_Count,     // $15
		exp.Object_Count,       // $16
		exp.Srok_Exp,           // $17 (указатель)
		exp.Stop_Date,          // $18 (указатель)
		exp.Stop_Reason,        // $19 (указатель)
		exp.Resuming_Date,      // $20 (указатель)
		exp.Srok_Resuming,      // $21 (указатель)
		exp.End_Date,           // $22 (указатель) -> новые поля тут
		exp.Day_Count,          // $23 (указатель)
		exp.Exp_Day_Count,      // $24 (указатель)
		exp.Cat_Vivod,          // $25 (указатель)
		exp.Possible_Vivod,     // $26 (указатель)
		exp.Impossible_Vivod,   // $27 (указатель)
		exp.Hour_Count,         // $28 (указатель)
		exp.Expert_Cost,        // $29 (указатель)
		exp.Material_Cost,      // $30 (указатель)
		exp.Exploitation_Cost,  // $31 (указатель)
		exp.Full_Cost,          // $32 (указатель)
		exp.Full_Cost_Nds,      // $33 (указатель)
		exp.Descrip,            // $34 (указатель)
		exp.Is_Closed,          // $35 (bool, улетит false или true в зависимости от запроса)
		exp.Stat_Id,            // $36 (указатель)
		exp.Category_Id,        // $37 (указатель)
		exp.Region_Id,          // $38 (указатель)
		exp.Iz_Nix_Id,          // $39 (указатель)
		exp.Diff_Cat_Id,        // $40 (указатель)
		exp.Exp_Res_Id,         // $41 (указатель)
	)

	return err
}
