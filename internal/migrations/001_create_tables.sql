-- +goose Up
PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS company (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    gst TEXT NOT NULL,
    address TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS payment_detail (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    acc_holder TEXT NOT NULL,
    acc_number INTEGER NOT NULL,
    ifsc TEXT NOT NULL,
    branch TEXT NOT NULL,
    bank_name TEXT NOT NULL,
    virtual_payment_addr TEXT NOT NULL,
    fk_company_id INTEGER NOT NULL,
    FOREIGN KEY (fk_company_id)
        REFERENCES company (id)
            ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS shipping_address (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    address TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS item (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    hsn INTEGER NOT NULL,
    price DECIMAL(19, 2) NOT NULL,
    gst INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS invoice_item (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    fk_invoice INTEGER,
    fk_item INTEGER,
    qty INTEGER,

    FOREIGN KEY (fk_invoice)
        REFERENCES invoice (id)
            ON DELETE CASCADE,

    FOREIGN KEY (fk_item)
        REFERENCES item (id)
            ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS invoice (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    invoice_number INTEGER NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    fk_from_company INTEGER NOT NULL,
    fk_to_company INTEGER NOT NULL,
    fk_payment_detail INTEGER NOT NULL,
    igst DECIMAL(19, 2) NOT NULL,
    TOTAL DECIMAL(19, 2) NOT NULL,
    fk_shipping_address integer NOT NULL,

    FOREIGN KEY (fk_from_company)
        REFERENCES company (id)
            ON DELETE CASCADE,

    FOREIGN KEY (fk_to_company)
        REFERENCES company (id)
            ON DELETE CASCADE,

    FOREIGN KEY (fk_payment_detail)
        REFERENCES payment_detail (id)
            ON DELETE CASCADE,

    FOREIGN KEY (fk_shipping_address)
        REFERENCES shipping_address (id)
            ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS invoice;
DROP TABLE IF EXISTS invoice_item;
DROP TABLE IF EXISTS item;
DROP TABLE IF EXISTS shipping_address;
DROP TABLE IF EXISTS payment_detail;
DROP TABLE IF EXISTS company;
PARGMA foreign_keys = OFF;
