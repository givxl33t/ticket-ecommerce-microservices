CREATE TABLE tickets(
    id INT NOT NULL AUTO_INCREMENT,
    title VARCHAR(100),
    price VARCHAR(100),
    user_id VARCHAR(255),
    order_id VARCHAR(255) DEFAULT NULL,
    created_at BIGINT,
    updated_at BIGINT,
    PRIMARY KEY(id)
);