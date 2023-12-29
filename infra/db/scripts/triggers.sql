DELIMITER // -- Study
CREATE TRIGGER tr_item_history AFTER UPDATE ON tb_item
FOR EACH ROW
BEGIN
    INSERT INTO tb_item_history(item_id, updated_at, deleted_at, code, name, description, price, category_id)
    VALUES (OLD.id, OLD.updated_at, OLD.deleted_at, OLD.code, OLD.name, OLD.description, OLD.price, OLD.category_id);
END;
//
DELIMITER //