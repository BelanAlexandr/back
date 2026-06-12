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
			srok_resuming, stat_id, category_id, region_id, iz_nix_id, diff_cat_id
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20,
			$21, $22, $23, $24, $25, $26
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
		exp.Stat_Id,            // $22
		exp.Category_Id,        // $23
		exp.Region_Id,          // $24
		exp.Iz_Nix_Id,          // $25
		exp.Diff_Cat_Id,        // $26
	)

	return err
}
