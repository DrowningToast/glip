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

-- Warehouses
CREATE TABLE warehouses (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    location_address TEXT NOT NULL,
    total_capacity DECIMAL(10,2) NOT NULL,
    current_capacity DECIMAL(10,2) NOT NULL,
    description TEXT,
    status VARCHAR(20), -- active, maintenance, closed
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Warehouse Connection (Authorization)
CREATE TABLE warehouse_connections (
    id SERIAL PRIMARY KEY,
    warehouse_id INTEGER NOT NULL REFERENCES warehouses(id),
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
    route int[] DEFAULT ARRAY[]::int[],
    last_warehouse_id INTEGER REFERENCES warehouses(id),
    destination_address TEXT NOT NULL,
    carrier_id INTEGER REFERENCES carriers(id),
    scheduled_departure TIMESTAMP NOT NULL,
    scheduled_arrival TIMESTAMP NOT NULL,
    actual_departure TIMESTAMP,
    actual_arrival TIMESTAMP,
    status VARCHAR(20), -- IN_TRANSIT_ON_THE_WAY, IN_TRANSIT_IN_WAREHOUSE, DELIVERED, CANCELLED
    total_weight DECIMAL(10,2) NOT NULL,
    total_volume DECIMAL(10,2) NOT NULL,
    special_instructions TEXT,
    estimated_duration INTEGER, -- in minutes
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Alerts and Notifications (Shipment-specific)
CREATE TABLE alerts (
    id SERIAL PRIMARY KEY,
    related_entity_type VARCHAR(50), -- shipment, carrier
    related_entity_id INTEGER,
    alert_type VARCHAR(50), -- delay, route_change, carrier_issue
    severity VARCHAR(20), -- low, medium, high
    description TEXT,
    status VARCHAR(20), -- new, acknowledged, resolved
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
); 