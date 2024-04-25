CREATE TABLE IF NOT EXISTS Chat_Users (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES Users(id),
    chat_id INTEGER REFERENCES Chats(id)
);