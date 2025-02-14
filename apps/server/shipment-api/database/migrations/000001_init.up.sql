-- Login Account
CREATE TABLE accounts (
    account_id SERIAL PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL, -- ADMIN, CARRIER_STAFF, CARRIER_VIEWER, OWNER
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Owner
CREATE TABLE owners (
    owner_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    phone VARCHAR(20),
    address TEXT,
    account_id INTEGER REFERENCES accounts(account_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Transportation Service Providers
CREATE TABLE carriers (
    carrier_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    carrier_type VARCHAR(50), -- internal, external
    contact_person VARCHAR(100),
    contact_phone VARCHAR(20),
    email VARCHAR(100),
    description TEXT,
    status VARCHAR(20), -- active, inactive
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Shipments
CREATE TABLE shipments (
    shipment_id SERIAL PRIMARY KEY,
    origin_warehouse_id INTEGER,
    destination_address TEXT NOT NULL,
    carrier_id INTEGER REFERENCES carriers(carrier_id),
    scheduled_departure TIMESTAMP,
    scheduled_arrival TIMESTAMP,
    actual_departure TIMESTAMP,
    actual_arrival TIMESTAMP,
    status VARCHAR(20), -- IN_TRANSIT_ON_THE_WAY, IN_TRANSIT_IN_WAREHOUSE, DELIVERED, CANCELLED
    total_weight DECIMAL(10,2),
    total_volume DECIMAL(10,2),
    special_instructions TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Transport Schedules
CREATE TABLE transport_schedules (
    schedule_id SERIAL PRIMARY KEY,
    shipment_id INTEGER REFERENCES shipments(shipment_id),
    planned_route int[],
    planned_departure TIMESTAMP,
    planned_arrival TIMESTAMP,
    estimated_duration INTEGER, -- in minutes
    route_status VARCHAR(20), -- scheduled, in-progress, completed
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Performance Reports
CREATE TABLE performance_reports (
    report_id SERIAL PRIMARY KEY,
    period_start DATE,
    period_end DATE,
    warehouse_id INTEGER,
    total_shipments INTEGER,
    on_time_delivery_rate DECIMAL(5,2),
    average_delivery_time INTEGER, -- in minutes
    inventory_turnover_rate DECIMAL(5,2),
    storage_utilization_rate DECIMAL(5,2),
    report_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Alerts and Notifications (Shipment-specific)
CREATE TABLE alerts (
    alert_id SERIAL PRIMARY KEY,
    related_entity_type VARCHAR(50), -- shipment, carrier
    related_entity_id INTEGER,
    alert_type VARCHAR(50), -- delay, route_change, carrier_issue
    severity VARCHAR(20), -- low, medium, high
    description TEXT,
    status VARCHAR(20), -- new, acknowledged, resolved
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Warehouse Connection (Authorization)
CREATE TABLE warehouse_connections (
    connection_id SERIAL PRIMARY KEY,
    warehouse_id INTEGER NOT NULL,
    api_key VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL,
    status VARCHAR(20) NOT NULL, -- active, inactive, revoked
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_used_at TIMESTAMP,
    created_by INTEGER REFERENCES accounts(account_id)
); 