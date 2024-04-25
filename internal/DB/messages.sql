CREATE TABLE IF NOT EXISTS Messages (
    id SERIAL PRIMARY KEY,
    sender_id INTEGER REFERENCES Users(id),
    msg TEXT,
    sent_on TIMESTAMP
);