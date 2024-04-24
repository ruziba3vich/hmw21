CREATE TABLE IF NOT EXISTS MessagesOfChats (
    id SERIAL PRIMARY KEY,
    chat_id INTEGER REFERENCES Chats(id),
    message_id INTEGER REFERENCES Messages(id)
);