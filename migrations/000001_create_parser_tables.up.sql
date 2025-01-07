CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    article VARCHAR(50) NOT NULL UNIQUE, 
    name VARCHAR(50) NOT NULL
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    article VARCHAR(50) NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL,
    image_path VARCHAR(50),
    category_id INT NOT NULL,
    price NUMERIC(10,2),
    stock INT,

    CONSTRAINT fk_product_category
        FOREIGN KEY (category_id)
        REFERENCES categories (id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE INDEX category_article ON categories (article);
CREATE INDEX product_article ON products (article);