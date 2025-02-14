-- Login Account
CREATE TABLE accounts (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL, -- ROOT, WAREHOUSE_STAFF, WAREHOUSE_VIEWER
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Staff
CREATE TABLE staffs (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    phone VARCHAR(20),
    address TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    account_id INT NOT NULL REFERENCES accounts(id)
);

-- Owner
CREATE TABLE owners (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    phone VARCHAR(20),
    address TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Warehouse Management
CREATE TABLE warehouses (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    location_address TEXT NOT NULL,
    total_capacity DECIMAL(10,2),
    current_capacity DECIMAL(10,2),
    description TEXT,
    status VARCHAR(20), -- active, maintenance, closed
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Storage Locations within Warehouses
CREATE TABLE storage_locations (
    id SERIAL PRIMARY KEY,
    warehouse_id INTEGER REFERENCES warehouses(id),
    area_name VARCHAR(50) NOT NULL, -- Zone identification
    -- minimum takes at least 1 unit
    capacity DECIMAL(10,2),
    current_occupancy DECIMAL(10,2), -- Current space used
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Inventory (Combined with Products)
CREATE TABLE inventory (
    id SERIAL PRIMARY KEY,
    sku VARCHAR(50) NOT NULL,
    name VARCHAR(200) NOT NULL,
    description TEXT,
    category VARCHAR(100),
    subcategory VARCHAR(100),
    owner_id INT REFERENCES owners(id),
    storage_location_id INTEGER REFERENCES storage_locations(id),
    quantity INTEGER NOT NULL,
    weight DECIMAL(10,2),
    dimensions JSON, -- {length, width, height}
    status VARCHAR(20), -- available, reserved, damaged
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    removed_at TIMESTAMP,
    UNIQUE(owner_id, storage_location_id, sku) -- Updated unique constraint to include SKU
);

-- Transportation Service Providers
CREATE TABLE carriers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    carrier_type VARCHAR(50), -- internal, external
    contact_person VARCHAR(100),
    contact_phone VARCHAR(20),
    email VARCHAR(100),
    description TEXT,
    status VARCHAR(20), -- active, inactive
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Stock Transactions
CREATE TABLE stock_transactions (
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    id SERIAL PRIMARY KEY,
    inventory_id INTEGER NOT NULL REFERENCES inventory(id),
    transaction_type VARCHAR(20) NOT NULL, -- IN, OUT, TRANSFER
    carrier_id INTEGER REFERENCES carriers(id), 
    quantity INTEGER NOT NULL,
    previous_quantity INTEGER NOT NULL,
    current_quantity INTEGER NOT NULL,
    reference_id VARCHAR(50), -- Related shipment or order ID
    reason VARCHAR(100), -- Purpose of transaction
    staff_performed_id INT NOT NULL REFERENCES staffs(id),
    meta JSON
);

-- Shipments
CREATE TABLE shipments (
    id SERIAL PRIMARY KEY,
    origin_warehouse_id INTEGER REFERENCES warehouses(id),
    destination_address TEXT NOT NULL,
    carrier_id INTEGER REFERENCES carriers(id),
    scheduled_departure TIMESTAMP,
    scheduled_arrival TIMESTAMP,
    actual_departure TIMESTAMP,
    actual_arrival TIMESTAMP,
    status VARCHAR(20), -- IN_TRANSIT_ON_THE_WAY, IN_TRANSIT_IN_WAREHOUSE, DELIEVERED, CANCELLED
    total_weight DECIMAL(10,2),
    total_volume DECIMAL(10,2),
    special_instructions TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Alerts and Notifications
CREATE TABLE alerts (
    id SERIAL PRIMARY KEY,
    related_entity_type VARCHAR(50), -- shipment, inventory, warehouse
    related_entity_id INTEGER,
    alert_type VARCHAR(50), -- delay, low_stock, maintenance
    severity VARCHAR(20), -- low, medium, high
    description TEXT,
    status VARCHAR(20), -- new, acknowledged, resolved
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);