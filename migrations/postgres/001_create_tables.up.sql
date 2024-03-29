CREATE TABLE IF NOT EXISTS clinic (
    id UUID PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
CREATE TABLE IF NOT EXISTS clinic_branch (
    id UUID PRIMARY KEY,
    clinic_id UUID REFERENCES clinic(id),
    address VARCHAR(50) NOT NULL,
    phone VARCHAR(13) NOT NULL,
    working_time VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
CREATE TABLE IF NOT EXISTS doctor_type (
    id UUID PRIMARY KEY,
    name   VARCHAR(50) NOT NULL,
    description TEXT NOT NULL,
    clinic_branch_id UUID REFERENCES clinic_branch(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
CREATE TABLE IF NOT EXISTS clinic_admin (
    id UUID PRIMARY KEY,
    clinic_branch_id UUID REFERENCES clinic_branch(id),
    doctor_type_id UUID REFERENCES doctor_type(id),
    first_name VARCHAR(25) NOT NULL,
    last_name VARCHAR(25) NOT NULL,
    email VARCHAR(50) NOT NULL,
    password VARCHAR(128) NOT NULL,
    phone VARCHAR(13) NOT NULL,
    gender VARCHAR(15) CHECK (gender IN('male', 'female')),
    birth_date DATE NOT NULL,
    age INT,
    address VARCHAR(120) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);-- great
CREATE TABLE IF NOT EXISTS doctor (
    id UUID PRIMARY KEY,
    doctor_type_id UUID REFERENCES doctor_type(id),
    first_name VARCHAR(25) NOT NULL,
    last_name VARCHAR(25) NOT NULL,
    email VARCHAR(50) NOT NULL,
    password VARCHAR(128) NOT NULL,
    phone VARCHAR(13) NOT NULL,
    gender VARCHAR(15) CHECK (gender IN('male', 'female')),
    birth_date DATE NOT NULL,
    age INT,
    address VARCHAR(120) NOT NULL,
    working_time VARCHAR(100) NOT NULL,
    status VARCHAR(15) CHECK (status IN('busy', 'empty')),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);-- great
CREATE TABLE IF NOT EXISTS customer (
    id UUID PRIMARY KEY,
    first_name VARCHAR(25) NOT NULL,
    last_name VARCHAR(25) NOT NULL,
    email VARCHAR(50) NOT NULL,
    password VARCHAR(128) NOT NULL,
    phone VARCHAR(13) NOT NULL,
    gender VARCHAR(15) CHECK (gender IN('male', 'female')),
    birth_date DATE NOT NULL,
    age INT,
    address VARCHAR(120) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
CREATE TABLE IF NOT EXISTS queue (
    id UUID PRIMARY KEY,
    customer_id UUID REFERENCES customer(id),
    doctor_id UUID REFERENCES doctor(id),
    queue_number VARCHAR(15) NOT NULL,
    queue_time VARCHAR(10) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
CREATE TABLE IF NOT EXISTS drug_store (
    id UUID PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description TEXT NOT NULL, 
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
CREATE TABLE IF NOT EXISTS drug_store_branch (
    id UUID PRIMARY KEY,
    drug_store_id UUID REFERENCES drug_store(id),
    address VARCHAR(120) NOT NULL,
    phone VARCHAR(13) NOT NULL,
    working_time VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
CREATE TABLE IF NOT EXISTS drug (
    id UUID PRIMARY KEY,
    drug_store_branch_id UUID REFERENCES drug_store_branch(id),
    name VARCHAR(50) NOT NULL,
    description TEXT NOT NULL,
    count INT NOT NULL,
    price NUMERIC(100,2) NOT NULL,
    date_of_manufacture VARCHAR(50) NOT NULL,
    best_before VARCHAR(15) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
CREATE TABLE IF NOT EXISTS pharmacist (
    id UUID PRIMARY KEY,
    drug_store_branch_id UUID REFERENCES drug_store_branch(id),
    first_name VARCHAR(25) NOT NULL,
    last_name VARCHAR(25) NOT NULL,
    email VARCHAR(50) NOT NULL,
    password VARCHAR(128) NOT NULL,
    phone VARCHAR(13) NOT NULL,
    gender VARCHAR(15) CHECK (gender IN('male', 'female')),
    birth_date DATE NOT NULL,
    age INT,
    address VARCHAR(120) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
CREATE TABLE IF NOT EXISTS orders (
    id UUID PRIMARY KEY,
    pharmacist_id UUID REFERENCES pharmacist(id),
    customer_id UUID REFERENCES customer(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
CREATE TABLE IF NOT EXISTS order_drug (
    id UUID PRIMARY KEY,
    drug_id UUID REFERENCES drug(id),
    orders_id UUID REFERENCES orders(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
CREATE TABLE IF NOT EXISTS author (
    id UUID PRIMARY KEY,
    first_name VARCHAR(25) NOT NULL,
    last_name VARCHAR(25) NOT NULL,
    email VARCHAR(50) NOT NULL,
    password VARCHAR(128) NOT NULL,
    phone VARCHAR(13) NOT NULL,
    gender VARCHAR(15) CHECK (gender IN('male', 'female')),
    birth_date DATE NOT NULL,
    age INT,
    address VARCHAR(120) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
CREATE TABLE IF NOT EXISTS journal (
    id UUID PRIMARY KEY,
    author_id UUID REFERENCES author(id),
    theme VARCHAR(150) NOT NULL,
    article TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
CREATE TABLE IF NOT EXISTS super_admin (
    id UUID PRIMARY KEY,
    clinic_id UUID REFERENCES clinic(id),
    drug_store_id UUID REFERENCES drug_store(id),
    author_id UUID REFERENCES author(id),
    first_name VARCHAR(25) NOT NULL,
    last_name VARCHAR(25) NOT NULL,
    email VARCHAR(50) NOT NULL,
    password VARCHAR(128) NOT NULL,
    phone VARCHAR(13) NOT NULL,
    gender VARCHAR(15) CHECK (gender IN('male', 'female')),
    birth_date DATE NOT NULL,
    age INT,
    address VARCHAR(120) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);