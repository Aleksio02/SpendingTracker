-- TODO: alexeyi: correct for own project!

CREATE SEQUENCE seq_ref_user_role START WITH 1 INCREMENT BY 1;
CREATE SEQUENCE seq_ref_spending_category START WITH 1 INCREMENT BY 1;
CREATE SEQUENCE seq_users START WITH 1 INCREMENT BY 1;
CREATE SEQUENCE seq_spendings START WITH 1 INCREMENT BY 1;

-- create references table
CREATE TABLE ref_user_role (
                    id INTEGER DEFAULT nextval('seq_ref_user_role') CONSTRAINT pk_ref_user_role PRIMARY KEY,
                    name VARCHAR CONSTRAINT un_ref_user_role_name UNIQUE
);

CREATE TABLE ref_spending_category (
                    id INTEGER DEFAULT nextval('seq_ref_spending_category') CONSTRAINT pk_ref_category PRIMARY KEY,
                    name VARCHAR CONSTRAINT un_ref_category_name UNIQUE
);


--create data table
CREATE TABLE users (
                    id INTEGER DEFAULT nextval('seq_users') CONSTRAINT pk_users PRIMARY KEY,
                    username VARCHAR CONSTRAINT un_users_username UNIQUE,
                    password VARCHAR CONSTRAINT nn_users_password NOT NULL,
                    role_id INTEGER CONSTRAINT nn_users_role NOT NULL,
                    CONSTRAINT fk_users_user_type FOREIGN KEY (role_id) references ref_user_role(id)
);

CREATE TABLE spending (
                    id INTEGER DEFAULT nextval('seq_spendings') CONSTRAINT pk_spending PRIMARY KEY,
                    description VARCHAR CONSTRAINT nn_spending_description NOT NULL,
                    create_date DATE CONSTRAINT nn_spending_creation_date NOT NULL,
                    spent FLOAT CONSTRAINT nn_spent NOT NULL,
                    category_id INTEGER CONSTRAINT nn_spending_category_id NOT NULL,
                    user_id INTEGER CONSTRAINT nn_spending_user_id NOT NULL,
                    CONSTRAINT fk_spending_category FOREIGN KEY (category_id) references ref_spending_category(id),
                    CONSTRAINT fk_spending_user FOREIGN KEY (user_id) references users(id)
);
