CREATE TABLE IF NOT EXISTS Messages_Of_Chats (
    id SERIAL PRIMARY KEY,
    chat_id INTEGER REFERENCES Chats(id),
    message_id INTEGER REFERENCES Messages(id)
);