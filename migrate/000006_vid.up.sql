INSERT INTO dict_iz_nix (name) VALUES 
('Единоличная');
CREATE TABLE IF NOT EXISTS dict_vid (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    shifr FLOAT NOT NULL
);
ALTER TABLE electronic_journal 
    DROP COLUMN IF EXISTS vid_exp;
 

ALTER TABLE electronic_journal 
    ADD COLUMN vid_exp_id INTEGER REFERENCES dict_vid(id);