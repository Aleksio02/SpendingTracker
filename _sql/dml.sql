-- Filling reference of user roles
insert into ref_user_role(name)
    values ('admin');

insert into ref_user_role(name)
    values ('user');

-- Filling reference of spending categories
insert into ref_spending_category(name)
    values ('Продукты');

insert into ref_spending_category(name)
    values ('Бытовая техника');

insert into ref_spending_category(name)
    values ('Авто');

insert into ref_spending_category(name)
    values ('Коммунальные услуги');

insert into ref_spending_category(name)
    values ('Другое');


-- Filling users
insert into users(username, password, role_id) values ('supervisor', 'supervisor', (select id from ref_user_role where name = 'admin'));

insert into users(username, password, role_id) values ('programmer2000', 'qwerty', (select id from ref_user_role where name = 'user'));


-- Filling some spendings
insert into spending(description, create_date, spent, category_id, user_id)
    values ('', TO_DATE('2024-01-11', 'YYYY-MM-DD'), 2487,
            (SELECT id FROM ref_spending_category WHERE name = 'Продукты'),
            (SELECT id FROM users WHERE username = 'supervisor'));

insert into spending(description, create_date, spent, category_id, user_id)
    values ('', TO_DATE('2024-01-09', 'YYYY-MM-DD'), 29999,
            (SELECT id FROM ref_spending_category WHERE name = 'Авто'),
            (SELECT id FROM users WHERE username = 'supervisor'));

insert into spending(description, create_date, spent, category_id, user_id)
    values ('', TO_DATE('2024-01-12', 'YYYY-MM-DD'), 10000,
            (SELECT id FROM ref_spending_category WHERE name = 'Другое'),
            (SELECT id FROM users WHERE username = 'supervisor'));

insert into spending(description, create_date, spent, category_id, user_id)
    values ('', TO_DATE('2024-01-03', 'YYYY-MM-DD'), 1250,
            (SELECT id FROM ref_spending_category WHERE name = 'Продукты'),
            (SELECT id FROM users WHERE username = 'programmer2000'));

insert into spending(description, create_date, spent, category_id, user_id)
    values ('', TO_DATE('2023-12-29', 'YYYY-MM-DD'), 12312,
            (SELECT id FROM ref_spending_category WHERE name = 'Продукты'),
            (SELECT id FROM users WHERE username = 'programmer2000'));

insert into spending(description, create_date, spent, category_id, user_id)
    values ('', TO_DATE('2024-01-08', 'YYYY-MM-DD'), 5700,
            (SELECT id FROM ref_spending_category WHERE name = 'Другое'),
            (SELECT id FROM users WHERE username = 'programmer2000'));