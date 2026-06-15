package repository

import (
	"context"

	"github.com/BelanAlexandr/back/internal/models"
)

func AddExpRepo(exp models.Exp) error {
	ctx := context.Background()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	queryJournal := `
		INSERT INTO electronic_journal (
			creator_id, data_post, fab, №adm_material, №stati, vid_exp, organ, name_organ,
			name_naznch, second_name_naznch, patronymic_naznch, 
			question_count, object_count, srok_exp, stop_date, stop_reason, resuming_date,
			srok_resuming, end_date, day_count, exp_day_count, cat_vivod, possible_vivod,
			impossible_vivod, hour_count, expert_cost, material_cost, exploitation_cost, full_cost,
			full_cost_nds, descrip, is_closed, stat_id, category_id, region_id, iz_nix_id,
			diff_cat_id, exp_res_id
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20,
			$21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38
		) RETURNING id;`

	var journalID int
	err = tx.QueryRowContext(
		ctx,
		queryJournal,
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
		exp.Question_Count,     // $12
		exp.Object_Count,       // $13
		exp.Srok_Exp,           // $14
		exp.Stop_Date,          // $15
		exp.Stop_Reason,        // $16
		exp.Resuming_Date,      // $17
		exp.Srok_Resuming,      // $18
		exp.End_Date,           // $19
		exp.Day_Count,          // $20
		exp.Exp_Day_Count,      // $21
		exp.Cat_Vivod,          // $22
		exp.Possible_Vivod,     // $23
		exp.Impossible_Vivod,   // $24
		exp.Hour_Count,         // $25
		exp.Expert_Cost,        // $26
		exp.Material_Cost,      // $27
		exp.Exploitation_Cost,  // $28
		exp.Full_Cost,          // $29
		exp.Full_Cost_Nds,      // $30
		exp.Descrip,            // $31
		exp.Is_Closed,          // $32
		exp.Stat_Id,            // $33
		exp.Category_Id,        // $34
		exp.Region_Id,          // $35
		exp.Iz_Nix_Id,          // $36
		exp.Diff_Cat_Id,        // $37
		exp.Exp_Res_Id,         // $38
	).Scan(&journalID)

	if err != nil {
		return err
	}

	queryLinks := `
		INSERT INTO electronic_journal_experts (journal_id, expert_id) 
		VALUES ($1, $2);`

	for _, expert := range exp.Experts {
		_, err = tx.ExecContext(ctx, queryLinks, journalID, expert.Id)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
