CREATE TABLE tickets(
    id INT NOT NULL,
    title VARCHAR(100),
    price VARCHAR(100),
    PRIMARY KEY(id)
);
CREATE TABLE orders(
    id INT NOT NULL AUTO_INCREMENT,
    status VARCHAR(100),
    user_id VARCHAR(255),
    ticket_id INT,
    expires_at BIGINT,
    created_at BIGINT,
    updated_at BIGINT,
    PRIMARY KEY(id),
    CONSTRAINT fk_ticket FOREIGN KEY (ticket_id) REFERENCES tickets(id) ON DELETE CASCADE
);