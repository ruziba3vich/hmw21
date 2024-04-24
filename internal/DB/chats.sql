CREATE TABLE IF NOT EXISTS Chats (
    id SERIAL PRIMARY KEY,
    chat_name VARCHAR(64),
    created_by INTEGER REFERENCES Users(id)
);