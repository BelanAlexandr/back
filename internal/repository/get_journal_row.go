package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/BelanAlexandr/back/internal/models"
)

func GetJournalRow(id int) (exp models.Exp, err error) {
	ctx := context.Background()

	queryJournal := `
        SELECT 
            id, creator_id, data_post, fab, 
            "№adm_material" AS adm_material, 
            "№stati" AS nom_statyi, 
            vid_exp, organ, name_organ,
            name_naznch, second_name_naznch, patronymic_naznch,
            question_count, object_count, srok_exp, stop_date, stop_reason, resuming_date,
            srok_resuming, end_date, day_count, exp_day_count, cat_vivod, possible_vivod,
            impossible_vivod, hour_count, expert_cost, material_cost, exploitation_cost, full_cost,
            full_cost_nds, descrip, is_closed, stat_id, category_id, region_id, iz_nix_id,
            diff_cat_id, exp_res_id
        FROM electronic_journal
        WHERE id = $1;`

	err = db.QueryRowContext(ctx, queryJournal, id).Scan(
		&exp.Id,
		&exp.Creator_id,
		&exp.Data_Post,
		&exp.Fab,
		&exp.Adm_Material,
		&exp.Nom_Statyi,
		&exp.Vid_Exp,
		&exp.Organ,
		&exp.Name_Organ,
		&exp.Name_Naznch,
		&exp.Second_Name_Naznch,
		&exp.Patronymic_Naznch,
		&exp.Question_Count,
		&exp.Object_Count,
		&exp.Srok_Exp,
		&exp.Stop_Date,
		&exp.Stop_Reason,
		&exp.Resuming_Date,
		&exp.Srok_Resuming,
		&exp.End_Date,
		&exp.Day_Count,
		&exp.Exp_Day_Count,
		&exp.Cat_Vivod,
		&exp.Possible_Vivod,
		&exp.Impossible_Vivod,
		&exp.Hour_Count,
		&exp.Expert_Cost,
		&exp.Material_Cost,
		&exp.Exploitation_Cost,
		&exp.Full_Cost,
		&exp.Full_Cost_Nds,
		&exp.Descrip,
		&exp.Is_Closed,
		&exp.Stat_Id,
		&exp.Category_Id,
		&exp.Region_Id,
		&exp.Iz_Nix_Id,
		&exp.Diff_Cat_Id,
		&exp.Exp_Res_Id,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return exp, errors.New("expertise not found")
		}
		return exp, err
	}

	queryExperts := `
        SELECT e.id, e.name, e.second_name, e.patronymic
        FROM dict_expert e
        JOIN electronic_journal_experts eje ON e.id = eje.expert_id
        WHERE eje.journal_id = $1;`

	rows, err := db.QueryContext(ctx, queryExperts, id)
	if err != nil {
		return exp, err
	}
	defer rows.Close()

	exp.Experts = make([]models.Expert, 0)

	for rows.Next() {
		var expUser models.Expert
		err := rows.Scan(
			&expUser.Id,
			&expUser.Name,
			&expUser.Second_Name,
			&expUser.Patronymic,
		)
		if err != nil {
			return exp, err
		}
		exp.Experts = append(exp.Experts, expUser)
	}

	if err = rows.Err(); err != nil {
		return exp, err
	}

	return exp, nil
}
