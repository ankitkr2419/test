CREATE FUNCTION maintain_process_count_trg() RETURNS TRIGGER AS
$$
BEGIN
  IF TG_OP IN ('DELETE' , 'UPDATE') THEN
    UPDATE recipes SET process_count = process_count - 1 WHERE id = old.recipe_id;
  END IF;
  IF TG_OP IN ('INSERT', 'UPDATE') THEN
    UPDATE recipes SET process_count = process_count + 1 WHERE id = new.recipe_id;
  END IF;
  IF TG_OP IN ('TRUNCATE') THEN
    UPDATE recipes SET process_count = 0;
  END IF;
  RETURN NULL;
END
$$
LANGUAGE plpgsql;

-- NOTE that a process might have different recipe_id thus be careful in UPDATE as well
CREATE TRIGGER maintain_process_count
AFTER INSERT OR UPDATE OF recipe_id OR DELETE ON processes
FOR EACH ROW
EXECUTE PROCEDURE maintain_process_count_trg();

-- TRUNCATE will not be FOR EACH ROW
CREATE TRIGGER maintain_process_count_for_truncate
AFTER TRUNCATE ON processes
EXECUTE PROCEDURE maintain_process_count_trg();