-- Default total_time is 900 seconds
ALTER TABLE IF EXISTS recipes ADD COLUMN total_time int DEFAULT 900;