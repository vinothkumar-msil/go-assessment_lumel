-- Customers Table
CREATE TABLE customers (
    customer_id VARCHAR(255) PRIMARY KEY,
    customer_name VARCHAR(255),
    customer_email VARCHAR(255),
    customer_address TEXT
);

-- Products Table
CREATE TABLE products (
    product_id VARCHAR(255) PRIMARY KEY,
    product_name VARCHAR(255),
    category VARCHAR(255)
);

-- Regions Table
CREATE TABLE regions (
    region_id SERIAL PRIMARY KEY,
    region_name VARCHAR(255)
);

-- Shipping Table
CREATE TABLE shipping (
    shipping_id SERIAL PRIMARY KEY,
    region_id INT REFERENCES regions(region_id),
    shipping_cost NUMERIC
);

-- Payments Table
CREATE TABLE payments (
    payment_id SERIAL PRIMARY KEY,
    payment_method VARCHAR(255)
);

-- Orders Table
CREATE TABLE orders (
    order_id VARCHAR(255) PRIMARY KEY,
    customer_id VARCHAR(255) REFERENCES customers(customer_id),
    order_date DATE,
    shipping_id INT REFERENCES shipping(shipping_id),
    payment_id INT REFERENCES payments(payment_id)
);

-- Sales Table
CREATE TABLE sales (
    sale_id SERIAL PRIMARY KEY,
    order_id VARCHAR(255) REFERENCES orders(order_id),
    product_id VARCHAR(255) REFERENCES products(product_id),
    quantity_sold INT,
    unit_price NUMERIC,
    discount NUMERIC
);
