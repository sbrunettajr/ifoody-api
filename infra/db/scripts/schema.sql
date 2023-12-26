CREATE TABLE tb_store(
    id INT NOT NULL AUTO_INCREMENT,
    created_at DATETIME NOT NULL DEFAULT NOW(),
    updated_at DATETIME NOT NULL DEFAULT NOW() ON UPDATE NOW(),
    deleted_at DATETIME,
    uuid VARCHAR(36) NOT NULL,
    fantasy_name VARCHAR(100) NOT NULL,
    corporate_name VARCHAR(100) NOT NULL,
    cnpj VARCHAR(14) NOT NULL, -- UNIQUE,
    PRIMARY KEY(id)
); 

CREATE TABLE tb_category(
    id INT NOT NULL AUTO_INCREMENT,
    created_at DATETIME NOT NULL DEFAULT NOW(),
    updated_at DATETIME NOT NULL DEFAULT NOW() ON UPDATE NOW(),
    deleted_at DATETIME,
    uuid VARCHAR(36) NOT NULL,
    name VARCHAR(50) NOT NULL,
    store_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (store_id) REFERENCES tb_store(id)
);

CREATE TABLE tb_item(
    id INT NOT NULL AUTO_INCREMENT,
    created_at DATETIME NOT NULL DEFAULT NOW(),
    updated_at DATETIME NOT NULL DEFAULT NOW() ON UPDATE NOW(),
    deleted_at DATETIME,
    uuid VARCHAR(36) NOT NULL,
    code VARCHAR(20) NOT NULL,
    name VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    price DOUBLE NOT NULL,
    category_id INT NOT NULL,
    store_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (category_id) REFERENCES tb_category(id),
    FOREIGN KEY (store_id) REFERENCES tb_store(id)
);