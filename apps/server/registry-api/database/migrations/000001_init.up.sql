-- Warehouse Connection (Authorization)
CREATE TABLE warehouse_connections (
    id SERIAL PRIMARY KEY,
    warehouse_id VARCHAR(255) NOT NULL,
    api_key VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL,
    status VARCHAR(20) NOT NULL, -- ACTIVE, INACTIVE, REVOKED
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_used_at TIMESTAMP
);