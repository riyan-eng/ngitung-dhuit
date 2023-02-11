CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE SCHEMA IF NOT EXISTS management;
CREATE SCHEMA IF NOT EXISTS finance;
CREATE SCHEMA IF NOT EXISTS master;

CREATE TABLE finance.coas(
  id VARCHAR DEFAULT uuid_generate_v4() PRIMARY KEY NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  code VARCHAR NOT NULL UNIQUE,
  parent VARCHAR,
  name VARCHAR,
  name_bahasa VARCHAR NOT NULL UNIQUE
);

INSERT INTO finance.coas(
  code, parent, name, name_bahasa
) VALUES
  ('1-1000', NULL, 'Current Assets', 'Aset lancar'),
  ('1-1100', '1-1000', 'Cash on Hand', 'Kas di tangan'),
  ('1-1200', '1-1000', 'Cash in Bank', 'Kas di bank'),
  ('1-1300', '1-1000', 'Accounts Receivalbe', 'Piutang dagang'),
  ('1-1400', '1-1000', 'Allowance for Doubtful Debt', 'Cadangan kerugian piutang'),
  ('1-1500', '1-1000', 'Merchandise Inventory', 'Persediaan barang dagangan'),
  ('1-1600', '1-1000', 'Store Supplies', 'Perlengkapan toko'),
  ('1-1700', '1-1000', 'Prepaid Insurance', 'Asuransi dibayar dimuka'),
  ('1-1800', '1-1000', 'Prepaid Rent', 'Sewa dibayar dimuka'),
  ('1-1900', '1-1000', 'Prepaid Tax', 'Uang muka PPh Ps 25'),
  ('1-2000', NULL, 'Fixed Assets', 'Aset tetap'),
  ('1-2100', '1-2000', 'Equipment at Cost', 'Peralatan'),
  ('1-2110', '1-2000', 'Equipment Accum Dep', 'Akumulasi penyusutan peralatan'),
  ('2-1000', NULL, 'Current Liabilities', 'Utang jangka pendek'),
  ('2-1100', '2-1000', 'Accounts Payable', 'Utang dagang'),
  ('2-1200', '2-1000', 'Expense Payable', 'Utang biaya'),
  ('2-1300', '2-1000', 'Income Tax Payable', 'Utang pajak penghasilan'),
  ('2-1400', '2-1000', 'PPN Payable', 'Utang PPN'),
  ('2-1500', '2-1000', 'PPN Outcome', 'PPN keluaran'),
  ('2-1600', '2-1000', 'PPN Income', 'PPN masukan'),
  ('2-2000', NULL, 'Long Term Liabilities', 'Utang jangka panjang'),
  ('2-2100', '2-2000', 'Bank Permata Loan', 'Utang Bank Permata'),
  ('3-1000', NULL, 'Equity', 'Ekuitas'),
  ('3-1100', '3-1000', 'Capital', 'Modal'),
  ('3-1200', '3-1000', 'Drawing', 'Prive'),
  ('3-1300', '3-1000', 'Income Summary', 'Ikhtisar laba rugi'),
  ('4-1000', NULL, 'Revenues', 'Pendapatan'),
  ('4-1100', '4-1000', 'Sales', 'Penjualan barang dagangan'),
  ('4-1200', '4-1000', 'Sales Retur', 'Retur penjualan'),
  ('5-1000', NULL, 'Cost of Goods Sold', 'HPP'),
  ('5-1100', '5-1000', 'Cost of Goods Sold', 'Harga pokok penjualan'),
  ('5-1200', '5-1000', 'Freight Paid', 'Beban transportasi pembelian'),
  ('6-0000', NULL, 'Operating Expenses', 'Beban operasi'),
  ('6-1000', '6-0000', 'Advertising Expenses', 'Beban iklan'),
  ('6-1100', '6-0000', 'Telephone & Electricity Expenses', 'Beban telepon dan listrik'),
  ('6-1200', '6-0000', 'Store Supplies Expenses', 'Beban perlengkapan toko'),
  ('6-1300', '6-0000', 'Bad Debts Expenses', 'Beban penghapusan kerugian piutang'),
  ('6-1400', '6-0000', 'Equipment Depreciation Expenses', 'Beban depresiasi peralatan'),
  ('6-1500', '6-0000', 'Insurance Expenses', 'Beban asuransi'),
  ('6-1600', '6-0000', 'Rent Expenses', 'Beban sewa toko'),
  ('6-1700', '6-0000', 'Wages  & Salaries', 'Beban upah & gaji'),
  ('6-1800', '6-0000', 'Other Operating Expenses', 'Beban-beban operasional lain'),
  ('8-1000', NULL, 'Other Revenues & Gains', 'Pendapatan dan keuntungan lain-lain'),
  ('8-1100', '8-1000', 'Interest Revenue', 'Pendapatan bunga'),
  ('9-1000', NULL, 'Other Expenses & Losses', 'Biaya dan kerugian lain-lain'),
  ('9-1100', '9-1000', 'Interest Expense', 'Beban bunga'),
  ('9-1200', '9-1000', 'Bank Service Charge', 'Beban administrasi bank'),
  ('9-1300', '9-1000', 'Income Tax Expense', 'Beban pajak penghasilan');


CREATE TABLE finance.goods(
  id VARCHAR DEFAULT uuid_generate_v4() PRIMARY KEY NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  code VARCHAR NOT NULL UNIQUE,
  name VARCHAR NOT NULL,
  description VARCHAR
);

INSERT INTO finance.goods(
  code, name, description
) VALUES
  ('HPP-14', 'HP Paviion 14', NULL),
  ('DI-15', 'Dell Inspiron 15', NULL);

CREATE TABLE finance.transactions(
  id VARCHAR DEFAULT uuid_generate_v4() PRIMARY KEY NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  description VARCHAR,
  amount NUMERIC NOT NULL
);

CREATE TABLE finance.purchase_journals(
  id VARCHAR DEFAULT uuid_generate_v4() PRIMARY KEY NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  transaction_id VARCHAR NOT NULL,
  coa_code VARCHAR NOT NULL,
  dc VARCHAR NOT NULL,
  amount NUMERIC NOT NULL,
  FOREIGN KEY (transaction_id) REFERENCES finance.transactions (id),
  FOREIGN KEY (coa_code) REFERENCES finance.coas (code)
);

CREATE TABLE finance.linked_accounts(
  id VARCHAR DEFAULT uuid_generate_v4() PRIMARY KEY NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  code VARCHAR NOT NULL UNIQUE,
  coa_code VARCHAR NOT NULL,
  FOREIGN KEY (coa_code) REFERENCES finance.coas (code)
);

INSERT INTO finance.linked_accounts(
  code, coa_code
) VALUES
  ('ppn_income', '2-1500'),
  ('ppn_outcome', '2-1600');

CREATE TABLE finance.taxes(
  id VARCHAR DEFAULT uuid_generate_v4() PRIMARY KEY NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  coa_code VARCHAR NOT NULL,
  description VARCHAR,
  rates_percent INTEGER NOT NULL,
  FOREIGN KEY (coa_code) REFERENCES finance.coas (code)
);

INSERT INTO finance.taxes(
  coa_code, description, rates_percent
) VALUES 
  ('2-1500', 'Pajak pertambahan nilai masuk', 11),
  ('2-1600', 'Pajak pertambahan nilai keluar', 11);