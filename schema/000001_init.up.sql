CREATE TABLE
    users (
        id serial primary key,
        name varchar(255) not null,
        username varchar(255) not null unique,
        password_hash varchar(255) not null
    );

CREATE TABLE
    todo_lists (
        id serial primary key,
        title varchar(255) not null,
        description varchar(255)
    );

CREATE TABLE
    users_lists (
        id serial primary key,
        user_id int not null,
        list_id int not null,
        foreign key (user_id) references users on delete cascade,
        foreign key (list_id) references todo_lists on delete cascade
    );

CREATE TABLE
    todo_items (
        id serial primary key,
        title varchar(255) not null,
        description varchar(255),
        done boolean not null default false
    );

CREATE TABLE
    lists_items (
        id serial primary key,
        item_id int not null,
        list_id int not null,
        foreign key (item_id) references todo_items on delete cascade,
        foreign key (list_id) references todo_lists on delete cascade
    );