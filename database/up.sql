DROP TABLE IF EXISTS foods;

CREATE TABLE foods(
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(32) NOT NULL,
    price FLOAT NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp,
    status BOOLEAN NOT NULL   
);

DROP TABLE IF EXISTS ingredients;

CREATE TABLE ingredients(
    idIngredients SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(32) NOT NULL,
    status BOOLEAN DEFAULT TRUE NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp,
    id INT NOT NULL,    
    CONSTRAINT FK_ingredients_foods FOREIGN KEY (id)
        REFERENCES foods(id) 
);