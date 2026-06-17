package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/BelanAlexandr/back/internal/models"
)

// Передаем ctx аргументом, чтобы контролировать таймауты и отмену запросов со стороны клиента
func UpdateExpRepo(exp models.Exp, closed bool) error {
	exp.Is_Closed = closed
	ctx := context.Background()
	// 1. Начинаем транзакцию
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	// Гарантированный откат в случае panic или раннего return err
	defer tx.Rollback()

	// 2. Обновляем основную запись экспертизы
	queryJournal := `
		UPDATE electronic_journal 
		SET 
			creator_id = $1, data_post = $2, fab = $3, "№adm_material" = $4, "№stati" = $5, 
			vid_exp = $6, organ = $7, name_organ = $8, name_naznch = $9, second_name_naznch = $10, 
			patronymic_naznch = $11, question_count = $12, object_count = $13, srok_exp = $14, 
			stop_date = $15, stop_reason = $16, resuming_date = $17, srok_resuming = $18, 
			end_date = $19, day_count = $20, exp_day_count = $21, cat_vivod = $22, 
			possible_vivod = $23, impossible_vivod = $24, hour_count = $25, expert_cost = $26, 
			material_cost = $27, exploitation_cost = $28, full_cost = $29, full_cost_nds = $30, 
			descrip = $31, is_closed = $32, stat_id = $33, category_id = $34, region_id = $35, 
			iz_nix_id = $36, diff_cat_id = $37, exp_res_id = $38
		WHERE id = $39;`

	result, err := tx.ExecContext(
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
		exp.Id,                 // $39
	)
	if err != nil {
		return fmt.Errorf("failed to update electronic_journal: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("экспертиза с id %d не найдена для обновления", exp.Id)
	}

	// 3. Удаляем СТАРЫЕ связи этой экспертизы с экспертами (подготовка к перезаписи)
	queryDeleteLinks := `DELETE FROM electronic_journal_experts WHERE journal_id = $1;`
	_, err = tx.ExecContext(ctx, queryDeleteLinks, exp.Id)
	if err != nil {
		return fmt.Errorf("failed to delete old links: %w", err)
	}

	// Подготовка SQL-запросов для цикла экспертов
	queryCheckExpert := `
		SELECT id FROM dict_expert 
		WHERE name = $1 AND second_name = $2 AND patronymic = $3;`

	queryCreateExpert := `
		INSERT INTO dict_expert (name, second_name, patronymic) 
		VALUES ($1, $2, $3) RETURNING id;`

	queryLinks := `
		INSERT INTO electronic_journal_experts (journal_id, expert_id) 
		VALUES ($1, $2);`

	// 4. ОБРАБОТКА МАССИВА ЭКСПЕРТОВ (Поиск / Создание -> Привязка)
	for _, expert := range exp.Experts {
		var expertID int

		// Проверяем по ФИО, есть ли уже такой эксперт в справочнике
		err = tx.QueryRowContext(ctx, queryCheckExpert, expert.Name, expert.Second_Name, expert.Patronymic).Scan(&expertID)

		if err == sql.ErrNoRows {
			// Эксперт не найден — создаем новую запись в справочнике
			err = tx.QueryRowContext(ctx, queryCreateExpert, expert.Name, expert.Second_Name, expert.Patronymic).Scan(&expertID)
			if err != nil {
				return fmt.Errorf("failed to create new expert %s: %w", expert.Name, err)
			}
		} else if err != nil {
			// Любая другая ошибка при чтении из базы данных
			return fmt.Errorf("failed to check expert: %w", err)
		}

		// Теперь точно связываем экспертизу с валидным ID эксперта (новым или старым)
		_, err = tx.ExecContext(ctx, queryLinks, exp.Id, expertID)
		if err != nil {
			return fmt.Errorf("failed to link expert %d to journal %d: %w", expertID, exp.Id, err)
		}
	}

	// 5. Коммитим транзакцию, если все этапы прошли успешно
	return tx.Commit()
}
