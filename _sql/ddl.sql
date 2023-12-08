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
                          id INTEGER DEFAULT nextval('seq_spendings') CONSTRAINT pk_articles PRIMARY KEY,
                          title VARCHAR CONSTRAINT nn_article_title NOT NULL,
                          content VARCHAR CONSTRAINT nn_article_content NOT NULL,
                          create_date DATE CONSTRAINT nn_article_creation_date NOT NULL,
                          author_id INTEGER CONSTRAINT nn_spending_author NOT NULL,
                          category_id INTEGER CONSTRAINT nn_article_category NOT NULL,
                          status_id INTEGER CONSTRAINT nn_article_status NOT NULL,
                          CONSTRAINT fk_article_users FOREIGN KEY (author_id) references users(id),
                          CONSTRAINT fk_spending_category FOREIGN KEY (category_id) references ref_spending_category(id),
);
