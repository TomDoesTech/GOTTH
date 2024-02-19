CREATE TABLE users (
    id SERIAL PRIMARY KEY, 
    firstname TEXT NOT NULL, 
    lastname TEXT NOT NULL
);

CREATE TABLE todos (
    id SERIAL PRIMARY KEY, 
    user_id SERIAL NOT NULL REFERENCES users(id),
    task TEXT NOT NULL,
    done BOOLEAN NOT NULL
);