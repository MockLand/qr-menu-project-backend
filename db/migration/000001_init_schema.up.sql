CREATE TABLE Users (
    user_id SERIAL PRIMARY KEY NOT NULL,
    username VARCHAR(100) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE qr_codes (
    id SERIAL PRIMARY KEY,
    data VARCHAR(255),
    image BYTEA
);

CREATE TABLE Categories (
    category_id SERIAL PRIMARY KEY NOT NULL,
    user_id INT REFERENCES Users(user_id) NOT NULL,
    name VARCHAR(100) NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE Dishes (
    dish_id SERIAL PRIMARY KEY NOT NULL,
    user_id INT REFERENCES Users(user_id) NOT NULL,
    name VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    category_id INT REFERENCES Categories(category_id) NOT NULL,
    is_available BOOLEAN DEFAULT TRUE NOT NULL
);

CREATE TABLE Ingredients (
    ingredient_id SERIAL PRIMARY KEY NOT NULL,
    user_id INT REFERENCES Users(user_id) NOT NULL,
    name VARCHAR(100) NOT NULL,
    allergen_info TEXT NOT NULL
);

CREATE TABLE Dish_Ingredients (
    dish_id INT REFERENCES Dishes(dish_id) NOT NULL,
    ingredient_id INT REFERENCES Ingredients(ingredient_id) NOT NULL,
    user_id INT REFERENCES Users(user_id) NOT NULL,
    quantity VARCHAR(50) NOT NULL,
    PRIMARY KEY (dish_id, ingredient_id, user_id)
);

CREATE TABLE Menus (
    menu_id SERIAL PRIMARY KEY NOT NULL,
    user_id INT REFERENCES Users(user_id) NOT NULL,
    name VARCHAR(100) NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL
);

CREATE TABLE Menu_Dishes (
    menu_id INT REFERENCES Menus(menu_id) NOT NULL,
    dish_id INT REFERENCES Dishes(dish_id) NOT NULL,
    user_id INT REFERENCES Users(user_id) NOT NULL,
    PRIMARY KEY (menu_id, dish_id, user_id)
);