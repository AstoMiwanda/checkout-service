CREATE
EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE products
(
    id         UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    sku        VARCHAR          NOT NULL,
    name       VARCHAR          NOT NULL,
    price      DOUBLE PRECISION NOT NULL,
    is_active  BOOLEAN          DEFAULT TRUE,
    created_at TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL
);

CREATE TABLE stocks
(
    id         UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    product_id UUID REFERENCES products (id),
    qty        INTEGER NOT NULL,
    updated_at TIMESTAMP NULL,
    created_at TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE TABLE discounts
(
    id             UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    type           VARCHAR NOT NULL,
    description    TEXT,
    discount_value DOUBLE PRECISION,
    is_active      BOOLEAN          DEFAULT TRUE,
    valid_from     TIMESTAMP,
    valid_to       TIMESTAMP,
    created_at     TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP NULL,
    deleted_at     TIMESTAMP NULL
);

CREATE TABLE discount_rules
(
    id                UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    discount_id       UUID REFERENCES discounts (id),
    product_id        UUID REFERENCES products (id),
    role              VARCHAR(20) NOT NULL,
    quantity          INTEGER     NOT NULL,
    quantity_operator VARCHAR(20)      DEFAULT 'equal',
    created_at        TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at        TIMESTAMP NULL,
    deleted_at        TIMESTAMP NULL
);

CREATE TABLE customers
(
    id         UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name       VARCHAR        NOT NULL,
    email      VARCHAR UNIQUE NOT NULL,
    created_at TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL
);

CREATE TABLE orders
(
    id             UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    customer_id    UUID REFERENCES customers (id),
    status         VARCHAR(20)      DEFAULT 'pending',
    total_amount   DOUBLE PRECISION,
    total_discount DOUBLE PRECISION,
    total_payment  DOUBLE PRECISION,
    created_at     TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP NULL,
    deleted_at     TIMESTAMP NULL
);

CREATE TABLE order_details
(
    id         UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    order_id   UUID REFERENCES orders (id),
    product_id UUID REFERENCES products (id),
    quantity   INTEGER          NOT NULL,
    price      DOUBLE PRECISION NOT NULL,
    subtotal   DOUBLE PRECISION NOT NULL,
    created_at TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL
);

CREATE TABLE order_discounts
(
    id         UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    order_id   UUID NOT NULL,
    product_id UUID NOT NULL,
    name       TEXT,
    qty        INTEGER,
    created_at TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
