package repository

import (
	"context"
	"fmt"

	"github.com/BelanAlexandr/back/internal/models"
)

func CloseExpRepo(exp models.CloseExp) error {
	query := `
		UPDATE electronic_journal 
		SET 
			end_date = $1, 
			day_count = $2, 
			exp_day_count = $3, 
			cat_vivod = $4, 
			possible_vivod = $5, 
			impossible_vivod = $6, 
			hour_count = $7, 
			expert_cost = $8, 
			material_cost = $9, 
			exploitation_cost = $10, 
			full_cost = $11, 
			full_cost_nds = $12, 
			descrip = $13, 
			is_closed = $14, 
			exp_res_id = $15
		WHERE id = $16;`

	result, err := db.ExecContext(
		context.Background(),
		query,
		exp.End_Date,          // $1
		exp.Day_Count,         // $2
		exp.Exp_Day_Count,     // $3
		exp.Cat_Vivod,         // $4
		exp.Possible_Vivod,    // $5
		exp.Impossible_Vivod,  // $6
		exp.Hour_Count,        // $7
		exp.Expert_Cost,       // $8  (указатель, запишет NULL, если пустой)
		exp.Material_Cost,     // $9  (указатель)
		exp.Exploitation_Cost, // $10 (указатель)
		exp.Full_Cost,         // $11 (указатель)
		exp.Full_Cost_Nds,     // $12 (указатель)
		exp.Descrip,           // $13 (указатель)
		exp.Is_Closed,         // $14 (обычный bool, улетит true/false)
		exp.Exp_Res_Id,        // $15
		exp.Id,                // $16 (ID из параметров)
	)
	if err != nil {
		return err
	}

	// Опциональная проверка: обновилась ли вообще хоть одна строка
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("экспертиза с id %d не найдена для обновления", exp.Id)
	}

	return nil
}
