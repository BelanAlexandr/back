package repository

import (
	"context"
	"fmt"

	"github.com/BelanAlexandr/back/internal/models"
	"github.com/lib/pq" // Важно для эффективного запроса через ANY($1)
)

func IndexGetRepo(offset int, limit int, sortField string, sortOrder string, statusFilter string, dateFrom string, dateTo string) ([]models.Exp, int, error) {
	ctx := context.Background()

	// --- 1. СНАЧАЛА СЧИТАЕМ ОБЩЕЕ КОЛИЧЕСТВО ЗАПИСЕЙ (COUNT) С УЧЕТОМ ФИЛЬТРОВ ---
	countQuery := `SELECT COUNT(*) FROM electronic_journal WHERE 1=1`
	var countArgs []interface{}
	countPlaceholderIdx := 1

	// Формируем условия фильтрации (одинаковые для COUNT и для SELECT)
	filterSQL := ""
	if statusFilter == "open" {
		filterSQL += fmt.Sprintf(" AND is_closed = $%d", countPlaceholderIdx)
		countArgs = append(countArgs, false)
		countPlaceholderIdx++
	} else if statusFilter == "closed" {
		filterSQL += fmt.Sprintf(" AND is_closed = $%d", countPlaceholderIdx)
		countArgs = append(countArgs, true)
		countPlaceholderIdx++
	}

	if dateFrom != "" && dateTo != "" {
		filterSQL += fmt.Sprintf(" AND data_post BETWEEN $%d AND $%d", countPlaceholderIdx, countPlaceholderIdx+1)
		countArgs = append(countArgs, dateFrom, dateTo)
		countPlaceholderIdx += 2
	} else if dateFrom != "" {
		filterSQL += fmt.Sprintf(" AND data_post >= $%d", countPlaceholderIdx)
		countArgs = append(countArgs, dateFrom)
		countPlaceholderIdx++
	} else if dateTo != "" {
		filterSQL += fmt.Sprintf(" AND data_post <= $%d", countPlaceholderIdx)
		countArgs = append(countArgs, dateTo)
		countPlaceholderIdx++
	}

	// Выполняем запрос COUNT
	var totalCount int
	err := db.QueryRowContext(ctx, countQuery+filterSQL, countArgs...).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	// Если в БД вообще нет записей под этот фильтр, нет смысла делать тяжелый выборку
	if totalCount == 0 {
		return make([]models.Exp, 0), 0, nil
	}

	// --- 2. ПОЛУЧАЕМ ДАННЫЕ С ТЕКУЩЕЙ СТРАНИЦЫ ---
	selectQuery := `
        SELECT 
            id, creator_id, data_post, fab, №adm_material, №stati, vid_exp, organ, name_organ,
            name_naznch, second_name_naznch, patronymic_naznch,
            question_count, object_count, srok_exp, stop_date, stop_reason, resuming_date,
            srok_resuming, end_date, day_count, exp_day_count, cat_vivod, possible_vivod,
            impossible_vivod, hour_count, expert_cost, material_cost, exploitation_cost, full_cost,
            full_cost_nds, descrip, is_closed, stat_id, category_id, region_id, iz_nix_id,
            diff_cat_id, exp_res_id
        FROM electronic_journal
        WHERE 1=1`

	// Копируем аргументы фильтрации для основного запроса
	selectArgs := make([]interface{}, len(countArgs))
	copy(selectArgs, countArgs)
	selectPlaceholderIdx := countPlaceholderIdx

	// Подставляем динамическую сортировку (sortField и sortOrder безопасны, проверены в Handler)
	// Важно: если на бэке поле называется №adm_material или №stati, подставляем корректные имена для SQL
	sqlSortField := sortField
	if sortField == "adm_material" {
		sqlSortField = "№adm_material"
	} else if sortField == "state" {
		sqlSortField = "№stati"
	}

	orderBySQL := fmt.Sprintf(" ORDER BY %s %s", sqlSortField, sortOrder)

	// Добавляем LIMIT и OFFSET под пагинацию MUI
	paginationSQL := fmt.Sprintf(" LIMIT $%d OFFSET $%d;", selectPlaceholderIdx, selectPlaceholderIdx+1)
	selectArgs = append(selectArgs, limit, offset)

	// Собираем всё воедино
	finalQuery := selectQuery + filterSQL + orderBySQL + paginationSQL

	rows, err := db.QueryContext(ctx, finalQuery, selectArgs...)
	if err != nil {
		return nil, 0, err
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
			return nil, 0, err
		}
		exp.Experts = make([]models.Expert, 0)
		exps = append(exps, exp)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	if len(exps) > 0 {
		// Подтягиваем экспертов одной пачкой через ANY
		if err := fillExpertsForExps(ctx, exps); err != nil {
			return nil, 0, err
		}
	}

	return exps, totalCount, nil
}

func IndexGetEmployeeRepo(creator_id int, offset int, limit int, sortField string, sortOrder string, statusFilter string, dateFrom string, dateTo string) ([]models.Exp, int, error) {
	ctx := context.Background()

	// --- 1. СЧИТАЕМ ОБЩЕЕ КОЛИЧЕСТВО ЗАПИСЕЙ СОТРУДНИКА (COUNT) ---
	countQuery := `SELECT COUNT(*) FROM electronic_journal WHERE creator_id = $1`
	countArgs := []interface{}{creator_id}
	countPlaceholderIdx := 2

	filterSQL := ""
	if statusFilter == "open" {
		filterSQL += fmt.Sprintf(" AND is_closed = $%d", countPlaceholderIdx)
		countArgs = append(countArgs, false)
		countPlaceholderIdx++
	} else if statusFilter == "closed" {
		filterSQL += fmt.Sprintf(" AND is_closed = $%d", countPlaceholderIdx)
		countArgs = append(countArgs, true)
		countPlaceholderIdx++
	}

	if dateFrom != "" && dateTo != "" {
		filterSQL += fmt.Sprintf(" AND data_post BETWEEN $%d AND $%d", countPlaceholderIdx, countPlaceholderIdx+1)
		countArgs = append(countArgs, dateFrom, dateTo)
		countPlaceholderIdx += 2
	} else if dateFrom != "" {
		filterSQL += fmt.Sprintf(" AND data_post >= $%d", countPlaceholderIdx)
		countArgs = append(countArgs, dateFrom)
		countPlaceholderIdx++
	} else if dateTo != "" {
		filterSQL += fmt.Sprintf(" AND data_post <= $%d", countPlaceholderIdx)
		countArgs = append(countArgs, dateTo)
		countPlaceholderIdx++
	}

	var totalCount int
	err := db.QueryRowContext(ctx, countQuery+filterSQL, countArgs...).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	if totalCount == 0 {
		return make([]models.Exp, 0), 0, nil
	}

	// --- 2. ВЫБОРКА СТРОК ДЛЯ СОТРУДНИКА ---
	selectQuery := `
        SELECT 
            id, creator_id, data_post, fab, №adm_material, №stati, vid_exp, organ, name_organ,
            name_naznch, second_name_naznch, patronymic_naznch,
            question_count, object_count, srok_exp, stop_date, stop_reason, resuming_date,
            srok_resuming, end_date, day_count, exp_day_count, cat_vivod, possible_vivod,
            impossible_vivod, hour_count, expert_cost, material_cost, exploitation_cost, full_cost,
            full_cost_nds, descrip, is_closed, stat_id, category_id, region_id, iz_nix_id,
            diff_cat_id, exp_res_id
        FROM electronic_journal
        WHERE creator_id = $1`

	selectArgs := make([]interface{}, len(countArgs))
	copy(selectArgs, countArgs)
	selectPlaceholderIdx := countPlaceholderIdx

	sqlSortField := sortField
	if sortField == "adm_material" {
		sqlSortField = "№adm_material"
	} else if sortField == "state" {
		sqlSortField = "№stati"
	}

	orderBySQL := fmt.Sprintf(" ORDER BY %s %s", sqlSortField, sortOrder)
	paginationSQL := fmt.Sprintf(" LIMIT $%d OFFSET $%d;", selectPlaceholderIdx, selectPlaceholderIdx+1)
	selectArgs = append(selectArgs, limit, offset)

	finalQuery := selectQuery + filterSQL + orderBySQL + paginationSQL

	rows, err := db.QueryContext(ctx, finalQuery, selectArgs...)
	if err != nil {
		return nil, 0, err
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
			return nil, 0, err
		}
		exp.Experts = make([]models.Expert, 0)
		exps = append(exps, exp)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	if len(exps) > 0 {
		if err := fillExpertsForExps(ctx, exps); err != nil {
			return nil, 0, err
		}
	}

	return exps, totalCount, nil
}

// Вспомогательная функция (fillExpertsForExps) остаётся БЕЗ ИЗМЕНЕНИЙ, так как она уже написана идеально
func fillExpertsForExps(ctx context.Context, exps []models.Exp) error {
	journalIDs := make([]int, len(exps))
	for i, exp := range exps {
		journalIDs[i] = exp.Id
	}

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

	expertsMap := make(map[int][]models.Expert)
	for rows.Next() {
		var journalID int
		var expert models.Expert
		err := rows.Scan(&journalID, &expert.Id, &expert.Name, &expert.Second_Name, &expert.Patronymic)
		if err != nil {
			return err
		}
		expertsMap[journalID] = append(expertsMap[journalID], expert)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	for i := range exps {
		if experts, ok := expertsMap[exps[i].Id]; ok {
			exps[i].Experts = experts
		}
	}
	return nil
}
