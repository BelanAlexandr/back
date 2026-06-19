-- 1. Возвращаем структуру таблицы electronic_journal обратно
ALTER TABLE electronic_journal 
    DROP COLUMN IF EXISTS vid_exp_id;

ALTER TABLE electronic_journal 
    ADD COLUMN IF NOT EXISTS vid_exp VARCHAR(255); -- или какой тип данных у вас был изначально (например, TEXT)

-- 2. Удаляем созданную таблицу справочника видов экспертиз
DROP TABLE IF EXISTS dict_vid CASCADE;

-- 3. Удаляем добавленное значение из таблицы dict_iz_nix
DELETE FROM dict_iz_nix 
WHERE name = 'Единоличная';