
CREATE TABLE messages (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    uuid char(32) UNIQUE NOT NULL,
    text text  NOT NULL,
    created_on TIMESTAMP default CURRENT_TIMESTAMP
);

