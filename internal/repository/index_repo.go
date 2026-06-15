package repository

import (
	"context"
	"fmt"

	"github.com/BelanAlexandr/back/internal/models"
	"github.com/lib/pq" // Важно: нужен для pq.Array, чтобы сделать эффективный запрос через ANY($1)
)

func IndexGetRepo(lastID int, limit int, statusFilter string, dateFrom string, dateTo string) ([]models.Exp, error) {
	ctx := context.Background()

	query := `
        SELECT 
            id, creator_id, data_post, fab, adm_material, nom_statyi, vid_exp, organ, name_organ,
            name_naznch, second_name_naznch, patronymic_naznch,
            question_count, object_count, srok_exp, stop_date, stop_reason, resuming_date,
            srok_resuming, end_date, day_count, exp_day_count, cat_vivod, possible_vivod,
            impossible_vivod, hour_count, expert_cost, material_cost, exploitation_cost, full_cost,
            full_cost_nds, descrip, is_closed, stat_id, category_id, region_id, iz_nix_id,
            diff_cat_id, exp_res_id
        FROM electronic_journal
        WHERE id < $1`

	args := []interface{}{lastID}
	placeholderIdx := 2

	if statusFilter == "open" {
		query += fmt.Sprintf(" AND is_closed = $%d", placeholderIdx)
		args = append(args, false)
		placeholderIdx++
	} else if statusFilter == "closed" {
		query += fmt.Sprintf(" AND is_closed = $%d", placeholderIdx)
		args = append(args, true)
		placeholderIdx++
	}

	if dateFrom != "" && dateTo != "" {
		query += fmt.Sprintf(" AND data_post BETWEEN $%d AND $%d", placeholderIdx, placeholderIdx+1)
		args = append(args, dateFrom, dateTo)
		placeholderIdx += 2
	} else if dateFrom != "" {
		query += fmt.Sprintf(" AND data_post >= $%d", placeholderIdx)
		args = append(args, dateFrom)
		placeholderIdx++
	} else if dateTo != "" {
		query += fmt.Sprintf(" AND data_post <= $%d", placeholderIdx)
		args = append(args, dateTo)
		placeholderIdx++
	}

	query += fmt.Sprintf(" ORDER BY id DESC LIMIT $%d;", placeholderIdx)
	args = append(args, limit)

	rows, err := db.QueryContext(ctx, query, args...)
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
			&exp.Patronymic_Naznch,
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
		// Инициализируем слайс экспертов сразу, чтобы избежать null в JSON
		exp.Experts = make([]models.Expert, 0)
		exps = append(exps, exp)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	// Если экспертиз нет, возвращаем пустой слайс
	if len(exps) == 0 {
		return exps, nil
	}

	// Привязываем экспертов ко всему списку за 1 запрос к БД
	if err := fillExpertsForExps(ctx, exps); err != nil {
		return nil, err
	}

	return exps, nil
}

func IndexGetEmployeeRepo(creator_id int, lastID int, limit int, statusFilter string, dateFrom string, dateTo string) ([]models.Exp, error) {
	ctx := context.Background()

	query := `
        SELECT 
            id, creator_id, data_post, fab, adm_material, nom_statyi, vid_exp, organ, name_organ,
            name_naznch, second_name_naznch, patronymic_naznch,
            question_count, object_count, srok_exp, stop_date, stop_reason, resuming_date,
            srok_resuming, end_date, day_count, exp_day_count, cat_vivod, possible_vivod,
            impossible_vivod, hour_count, expert_cost, material_cost, exploitation_cost, full_cost,
            full_cost_nds, descrip, is_closed, stat_id, category_id, region_id, iz_nix_id,
            diff_cat_id, exp_res_id
        FROM electronic_journal
        WHERE id < $1 AND creator_id = $2`

	args := []interface{}{lastID, creator_id}
	placeholderIdx := 3

	if statusFilter == "open" {
		query += fmt.Sprintf(" AND is_closed = $%d", placeholderIdx)
		args = append(args, false)
		placeholderIdx++
	} else if statusFilter == "closed" {
		query += fmt.Sprintf(" AND is_closed = $%d", placeholderIdx)
		args = append(args, true)
		placeholderIdx++
	}

	if dateFrom != "" && dateTo != "" {
		query += fmt.Sprintf(" AND data_post BETWEEN $%d AND $%d", placeholderIdx, placeholderIdx+1)
		args = append(args, dateFrom, dateTo)
		placeholderIdx += 2
	} else if dateFrom != "" {
		query += fmt.Sprintf(" AND data_post >= $%d", placeholderIdx)
		args = append(args, dateFrom)
		placeholderIdx++
	} else if dateTo != "" {
		query += fmt.Sprintf(" AND data_post <= $%d", placeholderIdx)
		args = append(args, dateTo)
		placeholderIdx++
	}

	query += fmt.Sprintf(" ORDER BY id DESC LIMIT $%d;", placeholderIdx)
	args = append(args, limit)

	rows, err := db.QueryContext(ctx, query, args...)
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
			&exp.Patronymic_Naznch,
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
		exp.Experts = make([]models.Expert, 0)
		exps = append(exps, exp)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(exps) == 0 {
		return exps, nil
	}

	// Привязываем экспертов ко всему списку за 1 запрос к БД
	if err := fillExpertsForExps(ctx, exps); err != nil {
		return nil, err
	}

	return exps, nil
}

// Вспомогательная функция для эффективной загрузки экспертов пакетом
func fillExpertsForExps(ctx context.Context, exps []models.Exp) error {
	// Сначала собираем все ID полученных экспертиз в один срез
	journalIDs := make([]int, len(exps))
	for i, exp := range exps {
		journalIDs[i] = exp.Id
	}

	// Запрашиваем из БД экспертов сразу для ВСЕХ этих журналов.
	// Конструкция eje.journal_id = ANY($1) позволяет передать массив ID.
	queryExperts := `
        SELECT eje.journal_id, e.id, e.name, e.second_name, e.patronymic
        FROM dict_expert e
        JOIN electronic_journal_experts eje ON e.id = eje.expert_id
        WHERE eje.journal_id = ANY($1);`

	rows, err := db.QueryContext(ctx, queryExperts, pq.Array(journalIDs))
	if err != nil {
		return err
	}
	defer rows.Close()

	// Используем мапу списков для временной группировки экспертов по ID журнала
	expertsMap := make(map[int][]models.Expert)

	for rows.Next() {
		var journalID int
		var expert models.Expert
		err := rows.Scan(
			&journalID,
			&expert.Id,
			&expert.Name,
			&expert.LastName,
			&expert.Patronymic,
		)
		if err != nil {
			return err
		}
		expertsMap[journalID] = append(expertsMap[journalID], expert)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	// Раскладываем экспертов из мапы по исходным структурам экспертиз.
	// Так как в Go структуры в слайсе передаются по значению, обходим слайс по индексам.
	for i := range exps {
		if experts, ok := expertsMap[exps[i].Id]; ok {
			exps[i].Experts = experts
		}
	}

	return nil
}
