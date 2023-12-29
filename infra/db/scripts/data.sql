INSERT
  INTO tb_store(uuid, fantasy_name, corporate_name, CNPJ)
VALUES (uuid(), 'Buratti Lanches', 'Buratti Lanches', '20475876000100');

INSERT
  INTO tb_category(uuid, name, store_id)
VALUES (uuid(), 'Xis', 1)
     , (uuid(), 'Torradas e cachorro-quente', 1)
     , (uuid(), 'Porções', 1)
     , (uuid(), 'Doces e sorvetes', 1)
     , (uuid(), 'Bebida', 1);