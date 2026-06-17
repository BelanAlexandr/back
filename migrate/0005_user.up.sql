-- Шаг 1: Добавляем колонки БЕЗ ограничения NOT NULL
ALTER TABLE users 
ADD COLUMN first_name VARCHAR(255),
ADD COLUMN last_name VARCHAR(255),
ADD COLUMN middle_name VARCHAR(255),
ADD COLUMN email VARCHAR(255) UNIQUE,
ADD COLUMN phone VARCHAR(50);

-- Шаг 2: Заполняем пустые поля у старых записей (например, временными заглушками)
UPDATE users SET first_name = 'Имя', last_name = 'Фамилия', middle_name = 'Отчество' 
WHERE first_name IS NULL;

-- Шаг 3: Теперь, когда пустых значений нет, принудительно включаем NOT NULL
ALTER TABLE users ALTER COLUMN first_name SET NOT NULL;
ALTER TABLE users ALTER COLUMN last_name SET NOT NULL;
CREATE INDEX idx_filenames ON electronic_journal(id,data_post,is_closed);
CREATE INDEX idx_users ON users(id)