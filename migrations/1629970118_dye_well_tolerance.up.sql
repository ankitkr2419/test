CREATE TABLE IF NOT EXISTS dye_well_tolerance(
    id uuid default uuid_generate_v4(),
    dye_id uuid,
    well_no INTEGER,
    kit_id text,
    optical_result DECIMAL DEFAULT 0.0000,
    valid boolean DEFAULT true,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT dwell UNIQUE (dye_id,well_no),
    FOREIGN KEY (dye_id) REFERENCES dyes(id) ON UPDATE CASCADE ON DELETE CASCADE
);