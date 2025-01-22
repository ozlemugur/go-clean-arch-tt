CREATE TABLE IF NOT EXISTS messages(
    id serial PRIMARY KEY,
    content TEXT CHECK (length(content) <= 160),
    recipient_phone VARCHAR(20),
    status VARCHAR(20) DEFAULT 'unsent',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);