-- 1. Создаем справочник экспертов (как у вас)
CREATE TABLE dict_expert (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    second_name VARCHAR(255) NOT NULL,
    patronymic VARCHAR(255) NOT NULL -- Исправил опечатку в patronymic
);

-- 2. Создаем таблицу связей
CREATE TABLE electronic_journal_experts (
    journal_id INTEGER REFERENCES electronic_journal(id) ON DELETE CASCADE,
    expert_id INTEGER REFERENCES dict_expert(id) ON DELETE CASCADE,
    PRIMARY KEY (journal_id, expert_id) -- Гарантирует, что одного и того же эксперта не привяжут к одной экспертизе дважды
);
-- 2. Удаляем старые текстовые колонки из журнала
ALTER TABLE electronic_journal
    DROP COLUMN IF EXISTS name_exp,
    DROP COLUMN IF EXISTS second_name_exp,
    DROP COLUMN IF EXISTS patronymic_exp,
    DROP COLUMN IF EXISTS exp_id; -- На всякий случай, если уже создали

-- 3. Создаем таблицу связи «многие ко многим»
CREATE TABLE electronic_journal_experts (
    journal_id INTEGER REFERENCES electronic_journal(id) ON DELETE CASCADE,
    expert_id INTEGER REFERENCES dict_expert(id) ON DELETE CASCADE,
    PRIMARY KEY (journal_id, expert_id) -- Защищает от дублирования связи
);