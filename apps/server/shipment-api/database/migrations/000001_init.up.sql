-- Login Account
CREATE TABLE accounts (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL, -- ADMIN, CARRIER_STAFF, CARRIER_VIEWER, OWNER
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Owner
CREATE TABLE owners (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    phone VARCHAR(20),
    address TEXT,
    account_id INTEGER REFERENCES accounts(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Transportation Service Providers
CREATE TABLE carriers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    contact_person VARCHAR(100),
    contact_phone VARCHAR(20),
    email VARCHAR(100),
    description TEXT,
    status VARCHAR(20) NOT NULL, -- active, inactive
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- Warehouse Connection (Authorization)
CREATE TABLE warehouse_connections (
    id SERIAL PRIMARY KEY,
    warehouse_id INT NOT NULL,
    api_key VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL,
    status VARCHAR(20) NOT NULL, -- ACTIVE, INACTIVE, REVOKED
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_used_at TIMESTAMP,
    created_by INTEGER REFERENCES accounts(id)
);

-- Shipments
CREATE TABLE shipments (
    id SERIAL PRIMARY KEY,
    route TEXT[] DEFAULT ARRAY[]::text[],
    last_warehouse_id TEXT,
    departure_warehouse_id TEXT NOT NULL,
    departure_address TEXT,
    destination_warehouse_id TEXT NOT NULL,
    destination_address TEXT NOT NULL,
    carrier_id INTEGER REFERENCES carriers(id),
    status VARCHAR(20) NOT NULL, -- WAITING_FOR_PICKUP, IN_TRANSIT_ON_THE_WAY, DELIVERED, CANCELLED
    total_weight DECIMAL(10,2) NOT NULL,
    total_volume DECIMAL(10,2) NOT NULL,
    special_instructions TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Alerts and Notifications (Shipment-specific)
CREATE TABLE alerts (
    id SERIAL PRIMARY KEY,
    related_entity_type VARCHAR(50) NOT NULL, -- shipment, carrier
    related_entity_id INTEGER NOT NULL,
    alert_type VARCHAR(50) NOT NULL, -- delay, route_change, carrier_issue
    severity VARCHAR(20) NOT NULL, -- low, medium, high
    description TEXT,
    status VARCHAR(20) NOT NULL, -- new, acknowledged, resolved
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
); 