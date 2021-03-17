DROP TRIGGER IF EXISTS maintain_process_count ON processes;

DROP TRIGGER IF EXISTS maintain_process_count_for_truncate ON processes;

DROP FUNCTION IF EXISTS maintain_process_count_trg;