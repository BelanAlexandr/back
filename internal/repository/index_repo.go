package repository

import (
	"context"
	"fmt"

	"github.com/BelanAlexandr/back/internal/models"
	"github.com/lib/pq"
)

func IndexGetRepo(offset int, limit int, sortField string, sortOrder string, statusFilter string, dateFrom string, dateTo string) ([]models.ExpListItem, int, error) {
	ctx := context.Background()

	// 1. СЧИТАЕМ ОБЩЕЕ КОЛИЧЕСТВО ЗАПИСЕЙ (COUNT)
	countQuery := `SELECT COUNT(*) FROM electronic_journal WHERE 1=1`
	var countArgs []interface{}
	countPlaceholderIdx := 1

	// Формируем условия фильтрации
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
		return make([]models.ExpListItem, 0), 0, nil
	}

	// 2. ВЫБОРКА ТОЛЬКО ТЕХ ПОЛЕЙ, КОТОРЫЕ ОТОБРАЖАЮТСЯ В ТАБЛИЦЕ
	// Оставляем строго: id, data_post, is_closed, fab
	selectQuery := `
        SELECT 
            id, data_post, is_closed, fab
        FROM electronic_journal
        WHERE 1=1`

	selectArgs := make([]interface{}, len(countArgs))
	copy(selectArgs, countArgs)
	selectPlaceholderIdx := countPlaceholderIdx

	orderBySQL := fmt.Sprintf(" ORDER BY %s %s", sortField, sortOrder)
	paginationSQL := fmt.Sprintf(" LIMIT $%d OFFSET $%d;", selectPlaceholderIdx, selectPlaceholderIdx+1)
	selectArgs = append(selectArgs, limit, offset)

	finalQuery := selectQuery + filterSQL + orderBySQL + paginationSQL

	rows, err := db.QueryContext(ctx, finalQuery, selectArgs...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var exps []models.ExpListItem
	for rows.Next() {
		var exp models.ExpListItem
		// Сканируем только выбранные 4 поля. Остальные поля структуры models.Exp останутся дефолтными (нули/пустые строки)
		err := rows.Scan(
			&exp.Id,
			&exp.Data_Post,
			&exp.Is_Closed,
			&exp.Fab,
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

	// 3. ДОЗАГРУЖАЕМ ЭКСПЕРТОВ ДЛЯ ЭТИХ СТРОК
	// Функция fillExpertsForExps отработает отлично, так как ей нужен только exp.Id
	if len(exps) > 0 {
		if err := fillExpertsForExps(ctx, exps); err != nil {
			return nil, 0, err
		}
	}

	return exps, totalCount, nil
}

func IndexGetEmployeeRepo(creator_id int, offset int, limit int, sortField string, sortOrder string, statusFilter string, dateFrom string, dateTo string) ([]models.ExpListItem, int, error) {
	ctx := context.Background()

	// 1. СЧИТАЕМ ОБЩЕЕ КОЛИЧЕСТВО ЗАПИСЕЙ СОТРУДНИКА
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
		return make([]models.ExpListItem, 0), 0, nil
	}

	// 2. ВЫБОРКА ТОЛЬКО ОТОБРАЖАЕМЫХ ПОЛЕЙ ДЛЯ СОТРУДНИКА
	selectQuery := `
        SELECT 
            id, data_post, is_closed, fab
        FROM electronic_journal
        WHERE creator_id = $1`

	selectArgs := make([]interface{}, len(countArgs))
	copy(selectArgs, countArgs)
	selectPlaceholderIdx := countPlaceholderIdx

	orderBySQL := fmt.Sprintf(" ORDER BY %s %s", sortField, sortOrder)
	paginationSQL := fmt.Sprintf(" LIMIT $%d OFFSET $%d;", selectPlaceholderIdx, selectPlaceholderIdx+1)
	selectArgs = append(selectArgs, limit, offset)

	finalQuery := selectQuery + filterSQL + orderBySQL + paginationSQL

	rows, err := db.QueryContext(ctx, finalQuery, selectArgs...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var exps []models.ExpListItem
	for rows.Next() {
		var exp models.ExpListItem
		err := rows.Scan(
			&exp.Id,
			&exp.Data_Post,
			&exp.Is_Closed,
			&exp.Fab,
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

	// 3. ДОЗАГРУЖАЕМ ЭКСПЕРТОВ ДЛЯ ЭТИХ СТРОК
	if len(exps) > 0 {
		if err := fillExpertsForExps(ctx, exps); err != nil {
			return nil, 0, err
		}
	}

	return exps, totalCount, nil
}

// Функция fillExpertsForExps остается БЕЗ ИЗМЕНЕНИЙ (она вытаскивает только ФИО экспертов)
func fillExpertsForExps(ctx context.Context, exps []models.ExpListItem) error {
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
