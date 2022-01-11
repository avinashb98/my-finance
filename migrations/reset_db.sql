SET sql_safe_updates = FALSE;

USE defaultdb;
DROP DATABASE IF EXISTS myfin CASCADE;
CREATE DATABASE IF NOT EXISTS myfin;

USE myfin;

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    handle STRING UNIQUE NOT NULL,
    name STRING,
    email STRING UNIQUE,
    is_active BOOL DEFAULT true,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now() ON UPDATE now()
);

CREATE TABLE net_worth (
    user_id UUID PRIMARY KEY REFERENCES users(id),
    value INT,
    last_updated TIMESTAMP DEFAULT now() ON UPDATE now()
);

CREATE TYPE category_type AS ENUM ('expense', 'asset', 'liability', 'income');

CREATE TABLE category (
    id STRING PRIMARY KEY,
    type category_type NOT NULL,
    name STRING,
    symbol STRING,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now() ON UPDATE now()
);

CREATE TABLE asset_category (
    id STRING PRIMARY KEY,
    name STRING,
    symbol STRING,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now() ON UPDATE now()
);

CREATE TABLE income (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id),
    amount DECIMAL(20, 2),
    dest_id UUID NOT NULL,
    source_name STRING NOT NULL,
    description STRING,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now() ON UPDATE now()
);

CREATE TABLE expense (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id),
    amount DECIMAL(20, 2),
    source_id UUID NOT NULL,
    dest_name STRING NOT NULL,
    description STRING,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now() ON UPDATE now()
);

CREATE TABLE user_asset (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name STRING NOT NULL,
    user_id UUID REFERENCES users(id),
    category_id STRING REFERENCES category(id),
    amount DECIMAL(20, 2)
);

CREATE TABLE user_liability (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name STRING NOT NULL,
    user_id UUID REFERENCES users(id),
    category_id STRING REFERENCES category(id),
    amount DECIMAL(20, 2)
);