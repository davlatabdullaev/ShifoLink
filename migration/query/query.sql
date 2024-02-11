--queue table ga malumot kiritganda queue number uchun quyidagi trigger ishlaydi

CREATE OR REPLACE FUNCTION generate_queue_number() RETURNS TRIGGER AS $$
DECLARE
    doc_count INTEGER;
    new_queue_number VARCHAR(15);
BEGIN

    SELECT COUNT(*) INTO doc_count FROM queue WHERE doctor_id = NEW.doctor_id;

    new_queue_number := (SELECT CONCAT(d.first_name, '-', LPAD((doc_count + 1)::TEXT, 4, '0')) FROM doctor d WHERE d.id = NEW.doctor_id);

    NEW.queue_number := new_queue_number;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER before_insert_queue
BEFORE INSERT ON queue
FOR EACH ROW
EXECUTE FUNCTION generate_queue_number();
