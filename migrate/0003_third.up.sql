
CREATE TABLE dict_status (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);
CREATE TABLE dict_iz_nix(
id SERIAL PRIMARY KEY,
name VARCHAR(255) NOT NULL UNIQUE
);
CREATE TABLE dict_category (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE dict_region (
    te INTEGER NOT NULL UNIQUE,
    ab INTEGER NOT NULL,
    cd INTEGER NOT NULL,
    ef INTEGER NOT NULL,
    hij INTEGER NOT NULL,
    k INTEGER NOT NULL,
    kaz_name VARCHAR(255) NOT NULL,
    rus_name VARCHAR(255) NOT NULL,
    nn INTEGER
);
CREATE TABLE dict_diff_cat(
     id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);
CREATE TABLE dict_exp_res(
     id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);

INSERT INTO dict_status (name) VALUES 
('Первичная'), ('Повторная'), ('Дополнительная');
INSERT INTO dict_iz_nix (name) VALUES 
('Комиссионная'), ('Комплексная');
INSERT INTO dict_category (name) VALUES 
('Уголовное'), ('Гражданское'), ('Административное');
INSERT INTO dict_diff_cat (name) VALUES 
('Простая'), ('Средней степени Сложности'), ('Сложная'), ('Особо Сложная');
INSERT INTO dict_exp_res (name) VALUES 
('Заключение'), ('СНДЗ'), ('Возврат без исполнения');

ALTER TABLE electronic_journal 
    DROP COLUMN IF EXISTS stat,
    DROP COLUMN IF EXISTS category,
    DROP COLUMN IF EXISTS region,
    DROP COLUMN IF EXISTS iz_nix,
    DROP COLUMN IF EXISTS category_diff,
    DROP COLUMN IF EXISTS exp_res;

ALTER TABLE electronic_journal 
    ADD COLUMN stat_id INTEGER REFERENCES dict_status(id),
    ADD COLUMN category_id INTEGER REFERENCES dict_category(id),
    ADD COLUMN region_id INTEGER REFERENCES dict_region(te),
    ADD COLUMN iz_nix_id INTEGER REFERENCES dict_iz_nix(id),
    ADD COLUMN diff_cat_id INTEGER REFERENCES dict_diff_cat(id),
    ADD COLUMN exp_res_id INTEGER REFERENCES dict_exp_res(id);
