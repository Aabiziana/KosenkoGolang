-- Script for complete database reset
-- This script will:
-- 1. Drop all project tables (in dependency order)
-- 2. Create all tables
-- 3. Fill tables with test data
-- Execute: psql -U root -d proxy_go -f sql/reset_database.sql

-- ==========================================
-- PART 1: DROP ALL TABLES
-- ==========================================

DROP TABLE IF EXISTS payroll_bodies CASCADE;
DROP TABLE IF EXISTS payrolls CASCADE;
DROP TABLE IF EXISTS proxy_bodies CASCADE;
DROP TABLE IF EXISTS proxies CASCADE;
DROP TABLE IF EXISTS products CASCADE;
DROP TABLE IF EXISTS organizations CASCADE;
DROP TABLE IF EXISTS employees CASCADE;
DROP TABLE IF EXISTS customers CASCADE;
DROP TABLE IF EXISTS accounts CASCADE;
DROP TABLE IF EXISTS units CASCADE;

-- ==========================================
-- PART 2: CREATE ALL TABLES
-- ==========================================

-- 1. Accounts
CREATE TABLE accounts (
    id SERIAL PRIMARY KEY,
    account VARCHAR(20) NOT NULL,
    bank_name TEXT NOT NULL,
    bank_identity_number VARCHAR(9) NOT NULL
);

-- 2. Units
CREATE TABLE units (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

-- 3. Products
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    price INT NOT NULL,
    unit_id INT NOT NULL DEFAULT 0 REFERENCES units(id) ON DELETE SET DEFAULT ON UPDATE CASCADE
);

-- 4. Organizations
CREATE TABLE organizations (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    address TEXT NOT NULL,
    account_id INT NOT NULL REFERENCES accounts(id) ON DELETE RESTRICT ON UPDATE CASCADE,
    chief TEXT NOT NULL,
    financial_chief TEXT NOT NULL,
    okud VARCHAR(20),
    okpo VARCHAR(20)
);

-- 5. Customers
CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

-- 6. Employees
CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    middle_name TEXT,
    post TEXT NOT NULL,
    passport_series VARCHAR(4) NOT NULL,
    passport_number VARCHAR(6) NOT NULL,
    passport_issued_by TEXT NOT NULL,
    passport_date_of_issue DATE NOT NULL
);

-- 7. Proxies
CREATE TABLE proxies (
    id SERIAL NOT NULL PRIMARY KEY,
    organization_id INT NOT NULL REFERENCES organizations(id) ON DELETE RESTRICT ON UPDATE CASCADE,
    customer_id INT NOT NULL REFERENCES customers(id) ON DELETE RESTRICT ON UPDATE CASCADE,
    employee_id INT NOT NULL REFERENCES employees(id) ON DELETE RESTRICT ON UPDATE CASCADE,
    date_of_issue DATE NOT NULL,
    is_valid_until DATE NOT NULL
);

-- 8. Proxy Bodies
CREATE TABLE proxy_bodies (
    id SERIAL NOT NULL PRIMARY KEY,
    product_id INT NOT NULL REFERENCES products(id) ON DELETE RESTRICT ON UPDATE CASCADE,
    proxy_id INT NOT NULL REFERENCES proxies(id) ON DELETE CASCADE ON UPDATE CASCADE,
    product_amount INT NOT NULL
);

CREATE TABLE payrolls (
    id SERIAL NOT NULL PRIMARY KEY,
    organization_id INT NOT NULL REFERENCES organizations(id) ON DELETE RESTRICT ON UPDATE CASCADE,
    document_number VARCHAR(50) NOT NULL,
    date_of_issue DATE NOT NULL,
    period_start DATE NOT NULL,
    period_end DATE NOT NULL,
    subdivision TEXT NOT NULL,
    corresponding_account VARCHAR(20) NOT NULL,
    payment_deadline_start DATE NOT NULL,
    payment_deadline_end DATE NOT NULL,
    total_amount_words TEXT NOT NULL,
    total_amount DECIMAL(12, 2) NOT NULL,
    deposited_amount DECIMAL(12, 2) DEFAULT 0,
    deposited_amount_words TEXT,
    chief TEXT NOT NULL,
    financial_chief TEXT NOT NULL,
    cashier TEXT NOT NULL,
    cash_order_number VARCHAR(50) NOT NULL,
    cash_order_date DATE NOT NULL,
    accountant TEXT NOT NULL,
    sheets_count INT DEFAULT 1
);

CREATE TABLE payroll_bodies (
    id SERIAL NOT NULL PRIMARY KEY,
    payroll_id INT NOT NULL REFERENCES payrolls(id) ON DELETE CASCADE ON UPDATE CASCADE,
    row_number INT NOT NULL,
    tabular_number VARCHAR(20) NOT NULL,
    employee_name TEXT NOT NULL,
    amount DECIMAL(12, 2) NOT NULL,
    signature TEXT,
    note TEXT
);

-- ==========================================
-- PART 3: INSERT TEST DATA
-- ==========================================

-- 1. Accounts
INSERT INTO accounts (account, bank_name, bank_identity_number) VALUES
('40817810000000000011', 'ПАО "Сбербанк России"', '044525225'),
('40817810000000000022', 'АО "Альфа-Банк"', '044525593'),
('40817810000000000033', 'ПАО "ВТБ"', '044525187');

-- 2. Units
INSERT INTO units (id, name) VALUES
(0, 'Не указано');
INSERT INTO units (name) VALUES
('шт'),
('кг'),
('т'),
('м'),
('м²'),
('м³'),
('л'),
('упак'),
('пачка'),
('рулон');

-- 3. Products
INSERT INTO products (name, price, unit_id) VALUES
('Бетон М300', 4200, (SELECT id FROM units WHERE name = 'м³' LIMIT 1)),
('Щебень фракция 5-20', 1200, (SELECT id FROM units WHERE name = 'м³' LIMIT 1)),
('Доска обрезная 50x150', 8500, (SELECT id FROM units WHERE name = 'м³' LIMIT 1)),
('Блоки газобетонные', 4200, (SELECT id FROM units WHERE name = 'м³' LIMIT 1)),
('Штукатурка цементная', 280, (SELECT id FROM units WHERE name = 'кг' LIMIT 1));

-- 4. Organizations
INSERT INTO organizations (name, address, account_id, chief, financial_chief, okud, okpo) VALUES
('ООО "ПромТехСервис"', 'г. Москва, ул. Промышленная, д. 15, офис 3', 
 (SELECT id FROM accounts WHERE account = '40817810000000000011' LIMIT 1),
 'Волков Сергей Иванович', 'Петрова Анна Владимировна', '0306004', '12345678'),
('ООО "ОфисТех"', 'г. Санкт-Петербург, Невский проспект, д. 28, офис 12',
 (SELECT id FROM accounts WHERE account = '40817810000000000022' LIMIT 1),
 'Смирнов Дмитрий Петрович', 'Козлова Елена Сергеевна', '0306004', '87654321'),
('ООО "СтройМаш"', 'г. Екатеринбург, ул. Машиностроителей, д. 45, склад 2',
 (SELECT id FROM accounts WHERE account = '40817810000000000033' LIMIT 1),
 'Иванов Андрей Николаевич', 'Соколова Мария Дмитриевна', '0306004', '11223344');

-- 5. Customers
INSERT INTO customers (name) VALUES
('ООО "СтройМонтаж"'),
('ООО "РемонтСтрой"'),
('ИП Новиков Дмитрий Сергеевич'),
('ООО "СтройТех"'),
('ООО "БыстроСтрой"');

-- 6. Employees
INSERT INTO employees (first_name, last_name, middle_name, post, passport_series, passport_number, passport_issued_by, passport_date_of_issue) VALUES
('Сергей', 'Волков', 'Иванович', 'Генеральный директор', '4501', '112233', 'Отделением УФМС России по г. Москве', '2015-03-10'),
('Анна', 'Петрова', 'Владимировна', 'Главный бухгалтер', '4502', '223344', 'Отделением УФМС России по г. Москве', '2016-04-15'),
('Дмитрий', 'Смирнов', 'Петрович', 'Директор', '7801', '334455', 'Отделением УФМС России по г. Санкт-Петербургу', '2017-05-20'),
('Елена', 'Козлова', 'Сергеевна', 'Бухгалтер', '7802', '445566', 'Отделением УФМС России по г. Санкт-Петербургу', '2018-06-25'),
('Андрей', 'Иванов', 'Николаевич', 'Директор', '6601', '556677', 'Отделением УФМС России по Свердловской области', '2019-07-30'),
('Мария', 'Соколова', 'Дмитриевна', 'Бухгалтер', '6602', '667788', 'Отделением УФМС России по Свердловской области', '2020-08-05'),
('Виктор', 'Николаенков', 'Леонидович', 'Менеджер по продажам', '4503', '778899', 'Отделением УФМС России по г. Москве', '2014-02-18'),
('Алексей', 'Макеев', 'Олегович', 'Старший менеджер', '4504', '889900', 'Отделением УФМС России по г. Москве', '2013-09-22'),
('Ольга', 'Старовойтова', 'Александровна', 'Кассир', '4505', '990011', 'Отделением УФМС России по г. Москве', '2016-11-05'),
('Анастасия', 'Фролкова', 'Викторовна', 'Бухгалтер', '4506', '001122', 'Отделением УФМС России по г. Москве', '2017-04-12'),
('Игорь', 'Сидоров', 'Павлович', 'Инженер', '4507', '112234', 'Отделением УФМС России по г. Москве', '2018-07-19'),
('Татьяна', 'Кузнецова', 'Михайловна', 'Экономист', '4508', '223345', 'Отделением УФМС России по г. Москве', '2019-01-28'),
('Павел', 'Морозов', 'Сергеевич', 'Водитель', '4509', '334456', 'Отделением УФМС России по г. Москве', '2015-08-14'),
('Наталья', 'Васильева', 'Ивановна', 'Секретарь', '4510', '445567', 'Отделением УФМС России по г. Москве', '2020-03-07'),
('Александр', 'Федоров', 'Дмитриевич', 'Снабженец', '7803', '556678', 'Отделением УФМС России по г. Санкт-Петербургу', '2016-06-21'),
('Екатерина', 'Новикова', 'Андреевна', 'Менеджер', '7804', '667789', 'Отделением УФМС России по г. Санкт-Петербургу', '2017-12-03');

-- 7. Proxies
INSERT INTO proxies (organization_id, customer_id, employee_id, date_of_issue, is_valid_until) VALUES
((SELECT id FROM organizations WHERE name = 'ООО "ПромТехСервис"' LIMIT 1),
 (SELECT id FROM customers WHERE name = 'ООО "СтройМонтаж"' LIMIT 1),
 (SELECT id FROM employees WHERE last_name = 'Волков' AND first_name = 'Сергей' LIMIT 1),
 '2024-10-15', '2025-01-15'),
((SELECT id FROM organizations WHERE name = 'ООО "ОфисТех"' LIMIT 1),
 (SELECT id FROM customers WHERE name = 'ООО "РемонтСтрой"' LIMIT 1),
 (SELECT id FROM employees WHERE last_name = 'Смирнов' AND first_name = 'Дмитрий' LIMIT 1),
 '2024-11-01', '2025-02-01'),
((SELECT id FROM organizations WHERE name = 'ООО "ПромТехСервис"' LIMIT 1),
 (SELECT id FROM customers WHERE name = 'ООО "СтройТех"' LIMIT 1),
 (SELECT id FROM employees WHERE last_name = 'Федоров' AND first_name = 'Александр' LIMIT 1),
 '2025-01-10', '2025-04-10'),
((SELECT id FROM organizations WHERE name = 'ООО "СтройМаш"' LIMIT 1),
 (SELECT id FROM customers WHERE name = 'ООО "БыстроСтрой"' LIMIT 1),
 (SELECT id FROM employees WHERE last_name = 'Иванов' AND first_name = 'Андрей' LIMIT 1),
 '2025-01-20', '2025-07-20');

-- 8. Proxy Bodies
INSERT INTO proxy_bodies (product_id, proxy_id, product_amount) VALUES
((SELECT id FROM products WHERE name = 'Бетон М300' LIMIT 1),
 (SELECT id FROM proxies WHERE date_of_issue = '2024-10-15' LIMIT 1), 8),
((SELECT id FROM products WHERE name = 'Щебень фракция 5-20' LIMIT 1),
 (SELECT id FROM proxies WHERE date_of_issue = '2024-10-15' LIMIT 1), 15),
((SELECT id FROM products WHERE name = 'Доска обрезная 50x150' LIMIT 1),
 (SELECT id FROM proxies WHERE date_of_issue = '2024-11-01' LIMIT 1), 3),
((SELECT id FROM products WHERE name = 'Блоки газобетонные' LIMIT 1),
 (SELECT id FROM proxies WHERE date_of_issue = '2024-11-01' LIMIT 1), 12),
((SELECT id FROM products WHERE name = 'Бетон М300' LIMIT 1),
 (SELECT id FROM proxies WHERE date_of_issue = '2025-01-10' LIMIT 1), 25),
((SELECT id FROM products WHERE name = 'Щебень фракция 5-20' LIMIT 1),
 (SELECT id FROM proxies WHERE date_of_issue = '2025-01-10' LIMIT 1), 40),
((SELECT id FROM products WHERE name = 'Штукатурка цементная' LIMIT 1),
 (SELECT id FROM proxies WHERE date_of_issue = '2025-01-10' LIMIT 1), 500),
((SELECT id FROM products WHERE name = 'Блоки газобетонные' LIMIT 1),
 (SELECT id FROM proxies WHERE date_of_issue = '2025-01-10' LIMIT 1), 50),
((SELECT id FROM products WHERE name = 'Доска обрезная 50x150' LIMIT 1),
 (SELECT id FROM proxies WHERE date_of_issue = '2025-01-20' LIMIT 1), 10),
((SELECT id FROM products WHERE name = 'Бетон М300' LIMIT 1),
 (SELECT id FROM proxies WHERE date_of_issue = '2025-01-20' LIMIT 1), 30),
((SELECT id FROM products WHERE name = 'Блоки газобетонные' LIMIT 1),
 (SELECT id FROM proxies WHERE date_of_issue = '2025-01-20' LIMIT 1), 100);

INSERT INTO payrolls (organization_id, document_number, date_of_issue, period_start, period_end, subdivision, corresponding_account, payment_deadline_start, payment_deadline_end, total_amount_words, total_amount, deposited_amount, deposited_amount_words, chief, financial_chief, cashier, cash_order_number, cash_order_date, accountant, sheets_count) VALUES
(
    (SELECT id FROM organizations WHERE name = 'ООО "ПромТехСервис"' LIMIT 1),
    '133',
    '2025-10-31',
    '2025-10-01',
    '2025-10-31',
    'Отдел продаж',
    '70',
    '2025-10-31',
    '2025-11-02',
    'Сто двадцать три тысячи рублей 00 копеек',
    123000.00,
    0.00,
    '',
    'Волков Сергей Иванович',
    'Петрова Анна Владимировна',
    'Старовойтова О. А.',
    '155',
    '2025-10-31',
    'Фролкова А. В.',
    1
),
(
    (SELECT id FROM organizations WHERE name = 'ООО "ПромТехСервис"' LIMIT 1),
    '134',
    '2025-11-30',
    '2025-11-01',
    '2025-11-30',
    'Бухгалтерия',
    '70',
    '2025-11-30',
    '2025-12-03',
    'Триста пятьдесят восемь тысяч пятьсот рублей 00 копеек',
    358500.00,
    15000.00,
    'Пятнадцать тысяч рублей 00 копеек',
    'Волков Сергей Иванович',
    'Петрова Анна Владимировна',
    'Старовойтова О. А.',
    '187',
    '2025-11-30',
    'Фролкова А. В.',
    1
),
(
    (SELECT id FROM organizations WHERE name = 'ООО "ОфисТех"' LIMIT 1),
    '45',
    '2025-12-25',
    '2025-12-01',
    '2025-12-31',
    'Производственный отдел',
    '70',
    '2025-12-25',
    '2025-12-28',
    'Пятьсот двадцать одна тысяча двести рублей 00 копеек',
    521200.00,
    0.00,
    '',
    'Смирнов Дмитрий Петрович',
    'Козлова Елена Сергеевна',
    'Новикова Е. А.',
    '98',
    '2025-12-25',
    'Козлова Е. С.',
    1
),
(
    (SELECT id FROM organizations WHERE name = 'ООО "СтройМаш"' LIMIT 1),
    '12',
    '2026-01-15',
    '2026-01-01',
    '2026-01-15',
    'Административный отдел',
    '70',
    '2026-01-15',
    '2026-01-18',
    'Двести восемьдесят девять тысяч рублей 00 копеек',
    289000.00,
    45000.00,
    'Сорок пять тысяч рублей 00 копеек',
    'Иванов Андрей Николаевич',
    'Соколова Мария Дмитриевна',
    'Васильева Н. И.',
    '5',
    '2026-01-15',
    'Соколова М. Д.',
    1
);

INSERT INTO payroll_bodies (payroll_id, row_number, tabular_number, employee_name, amount, signature, note) VALUES
(
    (SELECT id FROM payrolls WHERE document_number = '133' LIMIT 1),
    1,
    '10',
    'Николаенков В. Л.',
    56000.00,
    'Николаенков',
    ''
),
(
    (SELECT id FROM payrolls WHERE document_number = '133' LIMIT 1),
    2,
    '44',
    'Макеев А. О.',
    67000.00,
    'Макеев',
    ''
),
(
    (SELECT id FROM payrolls WHERE document_number = '134' LIMIT 1),
    1,
    '01',
    'Волков С. И.',
    95000.00,
    'Волков',
    ''
),
(
    (SELECT id FROM payrolls WHERE document_number = '134' LIMIT 1),
    2,
    '02',
    'Петрова А. В.',
    78500.00,
    'Петрова',
    ''
),
(
    (SELECT id FROM payrolls WHERE document_number = '134' LIMIT 1),
    3,
    '05',
    'Старовойтова О. А.',
    52000.00,
    'Старовойтова',
    ''
),
(
    (SELECT id FROM payrolls WHERE document_number = '134' LIMIT 1),
    4,
    '06',
    'Фролкова А. В.',
    58000.00,
    'Фролкова',
    ''
),
(
    (SELECT id FROM payrolls WHERE document_number = '134' LIMIT 1),
    5,
    '08',
    'Сидоров И. П.',
    60000.00,
    '',
    'депонировано'
),
(
    (SELECT id FROM payrolls WHERE document_number = '134' LIMIT 1),
    6,
    '12',
    'Кузнецова Т. М.',
    15000.00,
    'Кузнецова',
    'аванс'
),
(
    (SELECT id FROM payrolls WHERE document_number = '45' LIMIT 1),
    1,
    '101',
    'Смирнов Д. П.',
    120000.00,
    'Смирнов',
    ''
),
(
    (SELECT id FROM payrolls WHERE document_number = '45' LIMIT 1),
    2,
    '102',
    'Козлова Е. С.',
    89000.00,
    'Козлова',
    ''
),
(
    (SELECT id FROM payrolls WHERE document_number = '45' LIMIT 1),
    3,
    '103',
    'Федоров А. Д.',
    67200.00,
    'Федоров',
    ''
),
(
    (SELECT id FROM payrolls WHERE document_number = '45' LIMIT 1),
    4,
    '104',
    'Новикова Е. А.',
    72000.00,
    'Новикова',
    ''
),
(
    (SELECT id FROM payrolls WHERE document_number = '45' LIMIT 1),
    5,
    '105',
    'Морозов П. С.',
    55000.00,
    'Морозов',
    ''
),
(
    (SELECT id FROM payrolls WHERE document_number = '45' LIMIT 1),
    6,
    '106',
    'Васильева Н. И.',
    48000.00,
    'Васильева',
    ''
),
(
    (SELECT id FROM payrolls WHERE document_number = '45' LIMIT 1),
    7,
    '107',
    'Иванов А. Н.',
    70000.00,
    'Иванов',
    ''
),
(
    (SELECT id FROM payrolls WHERE document_number = '12' LIMIT 1),
    1,
    '201',
    'Иванов А. Н.',
    98000.00,
    'Иванов',
    ''
),
(
    (SELECT id FROM payrolls WHERE document_number = '12' LIMIT 1),
    2,
    '202',
    'Соколова М. Д.',
    75000.00,
    'Соколова',
    ''
),
(
    (SELECT id FROM payrolls WHERE document_number = '12' LIMIT 1),
    3,
    '203',
    'Морозов П. С.',
    56000.00,
    '',
    'депонировано'
),
(
    (SELECT id FROM payrolls WHERE document_number = '12' LIMIT 1),
    4,
    '204',
    'Кузнецова Т. М.',
    60000.00,
    'Кузнецова',
    ''
);

-- ==========================================
-- СООБЩЕНИЕ О ЗАВЕРШЕНИИ
-- ==========================================

DO $$
BEGIN
    RAISE NOTICE 'Сброс базы данных выполнен успешно!';
    RAISE NOTICE 'Все таблицы удалены, созданы заново и заполнены тестовыми данными.';
END $$;
