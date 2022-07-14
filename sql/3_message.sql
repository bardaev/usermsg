CREATE TABLE "go".message (
    id_msg INTEGER NOT NULL PRIMARY KEY,
    name VARCHAR(50),
    message TEXT,
    id_user INTEGER NOT NULL,
    FOREIGN KEY (id_user) REFERENCES "go".users (id_user)
)