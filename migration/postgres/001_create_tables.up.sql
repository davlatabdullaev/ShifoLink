CREATE TABLE clinic (
    id UUID PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
CREATE TABLE clinic_branch (
    id UUID PRIMARY KEY,
    clinic_id UUID REFERENCES clinic(id),
    address VARCHAR(50) NOT NULL,
    phone VARCHAR(13) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
CREATE TABLE doctor_type (
    id UUID PRIMARY KEY,
    name   VARCHAR(50) NOT NULL,
    description TEXT NOT NULL,
    clinic_branch_id UUID REFERENCES clinic_branch(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
CREATE TABLE clinic_admin (
    id UUID PRIMARY KEY,
    clinic_branch_id UUID REFERENCES clinic_branch(id),
    first_name VARCHAR(25) NOT NULL,
    last_name VARCHAR(25) NOT NULL,
    email VARCHAR(50) NOT NULL,
    password VARCHAR(128) NOT NULL,
    phone VARCHAR(13),
    gender VARCHAR(15) CHECK (gender IN('male', 'female')),
);


