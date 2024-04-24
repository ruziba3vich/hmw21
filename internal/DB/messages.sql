CREATE TABLE IF NOT EXISTS Messages (
    id SERIAL PRIMARY KEY,
    sender INTEGER REFERENCES Users(id),
    msg TEXT,
    sent_on TIMESTAMP
);