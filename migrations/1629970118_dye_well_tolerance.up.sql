CREATE TABLE IF NOT EXISTS dye_well_tolerance(
    id uuid default uuid_generate_v4(),
    dye_id uuid,
    well_no INTEGER,
    Tolernace DECIMAL,
    CONSTRAINT dwell UNIQUE (dye_id,well_no),
    FOREIGN KEY (dye_id) REFERENCES dyes(id) ON UPDATE CASCADE ON DELETE CASCADE
);