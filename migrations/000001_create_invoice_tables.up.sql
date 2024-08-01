-- 企業
CREATE TABLE companies (
    company_id BIGINT NOT NULL PRIMARY KEY,
    corporate_name VARCHAR(255) NOT NULL,
    representative_name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20),
    postal_code VARCHAR(10),
    address TEXT
);

-- ユーザー
CREATE TABLE users (
    user_id BIGINT NOT NULL PRIMARY KEY,
    company_id BIGINT NOT NULL,
    user_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    CONSTRAINT fk_users_company_id FOREIGN KEY (company_id) REFERENCES companies(company_id)
);

CREATE INDEX idx_users_company_id ON users(company_id);

-- 取引先
CREATE TABLE clients (
    client_id BIGINT NOT NULL PRIMARY KEY,
    company_id BIGINT NOT NULL,
    corporate_name VARCHAR(255) NOT NULL,
    representative_name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20),
    postal_code VARCHAR(10),
    address TEXT,
    CONSTRAINT fk_clients_company_id FOREIGN KEY (company_id) REFERENCES companies(company_id)
);

CREATE INDEX idx_clients_company_id ON clients(company_id);

-- 取引先銀行口座
CREATE TABLE client_bank_accounts (
    bank_account_id BIGINT NOT NULL PRIMARY KEY,
    client_id BIGINT NOT NULL,
    bank_name VARCHAR(255) NOT NULL,
    branch_name VARCHAR(255) NOT NULL,
    account_number VARCHAR(50) NOT NULL,
    account_name VARCHAR(255) NOT NULL,
    CONSTRAINT fk_client_bank_accounts_client_id FOREIGN KEY (client_id) REFERENCES clients(client_id)
);

CREATE INDEX idx_client_bank_accounts_client_id ON client_bank_accounts(client_id);

-- 請求書データ
CREATE TABLE invoices (
    invoice_id BIGINT NOT NULL PRIMARY KEY,
    company_id BIGINT NOT NULL,
    client_id BIGINT NOT NULL,
    issue_date DATE NOT NULL,
    payment_amount DECIMAL(10, 2) NOT NULL,
    fee DECIMAL(10, 2),
    fee_rate DECIMAL(5, 2),
    consumption_tax DECIMAL(10, 2),
    consumption_tax_rate DECIMAL(5, 2),
    billing_amount DECIMAL(10, 2) NOT NULL,
    payment_due_date DATE,
    status VARCHAR(20),
    CONSTRAINT fk_invoices_company_id FOREIGN KEY (company_id) REFERENCES companies(company_id),
    CONSTRAINT fk_invoices_client_id FOREIGN KEY (client_id) REFERENCES clients(client_id)
);

CREATE INDEX idx_invoices_company_id ON invoices(company_id);
CREATE INDEX idx_invoices_client_id ON invoices(client_id);
CREATE INDEX idx_invoices_payment_due_date ON invoices(payment_due_date);
