package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/BelanAlexandr/back/internal/models"
)

func AddExpRepo(exp models.Exp) error {
	ctx := context.Background()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	// Исправлено: добавлены кавычки для спецсимволов и vid_exp изменен на vid_exp_id
	queryJournal := `
		INSERT INTO electronic_journal (
			creator_id, data_post, fab, "№adm_material", "№stati", vid_exp_id, organ, name_organ,
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
		exp.Creator_id, exp.Data_Post, exp.Fab, exp.Adm_Material, exp.Nom_Statyi, exp.Vid_Exp, exp.Organ, exp.Name_Organ,
		exp.Name_Naznch, exp.Second_Name_Naznch, exp.Patronymic_Naznch, exp.Question_Count, exp.Object_Count, exp.Srok_Exp,
		exp.Stop_Date, exp.Stop_Reason, exp.Resuming_Date, exp.Srok_Resuming, exp.End_Date, exp.Day_Count, exp.Exp_Day_Count,
		exp.Cat_Vivod, exp.Possible_Vivod, exp.Impossible_Vivod, exp.Hour_Count, exp.Expert_Cost, exp.Material_Cost,
		exp.Exploitation_Cost, exp.Full_Cost, exp.Full_Cost_Nds, exp.Descrip, exp.Is_Closed, exp.Stat_Id, exp.Category_Id,
		exp.Region_Id, exp.Iz_Nix_Id, exp.Diff_Cat_Id, exp.Exp_Res_Id,
	).Scan(&journalID)

	if err != nil {
		return err
	}

	queryCheckExpert := `SELECT id FROM dict_expert WHERE name = $1 AND second_name = $2 AND patronymic = $3;`
	queryCreateExpert := `INSERT INTO dict_expert (name, second_name, patronymic) VALUES ($1, $2, $3) RETURNING id;`
	queryLinks := `INSERT INTO electronic_journal_experts (journal_id, expert_id) VALUES ($1, $2);`

	for _, expert := range exp.Experts {
		var expertID int

		err = tx.QueryRowContext(ctx, queryCheckExpert, expert.Name, expert.Second_Name, expert.Patronymic).Scan(&expertID)
		if err == sql.ErrNoRows {
			err = tx.QueryRowContext(ctx, queryCreateExpert, expert.Name, expert.Second_Name, expert.Patronymic).Scan(&expertID)
			if err != nil {
				return err
			}
		} else if err != nil {
			return err
		}

		_, err = tx.ExecContext(ctx, queryLinks, journalID, expertID)
		if err != nil {
			return err
		}
	}

	message := fmt.Sprintf("Создана новая экспертиза №%d в журнале", journalID)
	_, err = AddNotification(ctx, tx, exp.Creator_id, message)
	if err != nil {
		return err
	}

	return tx.Commit()
}
