package repository

import (
	"context"

	"github.com/BelanAlexandr/back/internal/models"
)

func IndexGetRepo(lastID int, limit int) ([]models.Exp, error) {
	query := `
		SELECT 
			id, creator_id, data_post, fab, adm_material, nom_statyi, vid_exp, organ, name_organ,
			name_naznch, second_name_naznch, patronymic_naznch, name_exp, second_name_exp, patronymic_exp,
			question_count, object_count, srok_exp, stop_date, stop_reason, resuming_date,
			srok_resuming, end_date, day_count, exp_day_count, cat_vivod, possible_vivod,
			impossible_vivod, hour_count, expert_cost, material_cost, exploitation_cost, full_cost,
			full_cost_nds, descrip, is_closed, stat_id, category_id, region_id, iz_nix_id,
			diff_cat_id, exp_res_id
		FROM electronic_journal
		WHERE id > $1
		ORDER BY id ASC
		LIMIT $2;`

	rows, err := db.QueryContext(context.Background(), query, lastID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var exps []models.Exp

	for rows.Next() {
		var exp models.Exp
		err := rows.Scan(
			&exp.Id, &exp.Creator_id, &exp.Data_Post, &exp.Fab, &exp.Adm_Material, &exp.Nom_Statyi,
			&exp.Vid_Exp, &exp.Organ, &exp.Name_Organ, &exp.Name_Naznch, &exp.Second_Name_Naznch,
			&exp.Patronymic_Naznch, &exp.Name_Exp, &exp.Second_Name_Exp, &exp.Patronymic_Exp,
			&exp.Question_Count, &exp.Object_Count, &exp.Srok_Exp, &exp.Stop_Date, &exp.Stop_Reason,
			&exp.Resuming_Date, &exp.Srok_Resuming, &exp.End_Date, &exp.Day_Count, &exp.Exp_Day_Count,
			&exp.Cat_Vivod, &exp.Possible_Vivod, &exp.Impossible_Vivod, &exp.Hour_Count,
			&exp.Expert_Cost, &exp.Material_Cost, &exp.Exploitation_Cost, &exp.Full_Cost,
			&exp.Full_Cost_Nds, &exp.Descrip, &exp.Is_Closed, &exp.Stat_Id, &exp.Category_Id,
			&exp.Region_Id, &exp.Iz_Nix_Id, &exp.Diff_Cat_Id, &exp.Exp_Res_Id,
		)
		if err != nil {
			return nil, err
		}
		exps = append(exps, exp)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return exps, nil
}
func IndexGetEmployeeRepo(creator_id, lastID int, limit int) ([]models.Exp, error) {
	query := `
		SELECT 
			id, creator_id, data_post, fab, adm_material, nom_statyi, vid_exp, organ, name_organ,
			name_naznch, second_name_naznch, patronymic_naznch, name_exp, second_name_exp, patronymic_exp,
			question_count, object_count, srok_exp, stop_date, stop_reason, resuming_date,
			srok_resuming, end_date, day_count, exp_day_count, cat_vivod, possible_vivod,
			impossible_vivod, hour_count, expert_cost, material_cost, exploitation_cost, full_cost,
			full_cost_nds, descrip, is_closed, stat_id, category_id, region_id, iz_nix_id,
			diff_cat_id, exp_res_id
		FROM electronic_journal
		WHERE id > $1 AND creator_id=$2
		ORDER BY id ASC
		LIMIT $3;`

	rows, err := db.QueryContext(context.Background(), query, lastID, creator_id, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var exps []models.Exp

	for rows.Next() {
		var exp models.Exp
		err := rows.Scan(
			&exp.Id, &exp.Creator_id, &exp.Data_Post, &exp.Fab, &exp.Adm_Material, &exp.Nom_Statyi,
			&exp.Vid_Exp, &exp.Organ, &exp.Name_Organ, &exp.Name_Naznch, &exp.Second_Name_Naznch,
			&exp.Patronymic_Naznch, &exp.Name_Exp, &exp.Second_Name_Exp, &exp.Patronymic_Exp,
			&exp.Question_Count, &exp.Object_Count, &exp.Srok_Exp, &exp.Stop_Date, &exp.Stop_Reason,
			&exp.Resuming_Date, &exp.Srok_Resuming, &exp.End_Date, &exp.Day_Count, &exp.Exp_Day_Count,
			&exp.Cat_Vivod, &exp.Possible_Vivod, &exp.Impossible_Vivod, &exp.Hour_Count,
			&exp.Expert_Cost, &exp.Material_Cost, &exp.Exploitation_Cost, &exp.Full_Cost,
			&exp.Full_Cost_Nds, &exp.Descrip, &exp.Is_Closed, &exp.Stat_Id, &exp.Category_Id,
			&exp.Region_Id, &exp.Iz_Nix_Id, &exp.Diff_Cat_Id, &exp.Exp_Res_Id,
		)
		if err != nil {
			return nil, err
		}
		exps = append(exps, exp)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return exps, nil
}
