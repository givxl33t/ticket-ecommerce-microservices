CREATE TABLE orders(
    id INT NOT NULL,
    status VARCHAR(100),
    user_id VARCHAR(255),
    price VARCHAR(100),
    PRIMARY KEY(id)
);
CREATE TABLE payments(
    id INT NOT NULL AUTO_INCREMENT,
    order_id INT,
    stripe_id VARCHAR(255),
    expires_at BIGINT,
    created_at BIGINT,
    updated_at BIGINT,
    PRIMARY KEY(id),
    CONSTRAINT fk_order FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE
);