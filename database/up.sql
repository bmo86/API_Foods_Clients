DROP TABLE IF EXISTS food;

CREATE TABLE food(
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(32) NOT NULL,
    price FLOAT NOT NULL,
    created_at timestamp NOT NULL,
    update_at timestamp NOT NULL,
    status BOOLEAN NOT NULL   
);

DROP TABLE IF EXISTS ingredient;

CREATE TABLE ingredient(
    idIngredient SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(32) NOT NULL,
    status BOOLEAN DEFAULT TRUE NOT NULL,
    id INT NOT NULL,    
    CONSTRAINT FK_ingredient_food FOREIGN KEY (id)
        REFERENCES food(id) 
);