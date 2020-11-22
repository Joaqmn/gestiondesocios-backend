USE fullstackAsociacion;

CREATE TABLE assoc_partners (
  id INT PRIMARY KEY,
  partnerName VARCHAR(30) NOT NULL
);

CREATE TABLE games (
  id INT PRIMARY KEY,
  id_owner INT,
  entry_date VARCHAR(200),
  disponibility BOOL NOT NULL,
  comments VARCHAR(200),
  CONSTRAINT fk_idOwner FOREIGN KEY (id_owner) REFERENCES assoc_partners (id)
);

CREATE TABLE borrowed_games (
  id_game INT PRIMARY KEY,
  id_borrower INT NOT NULL,
  borrow_date VARCHAR(200),
  FOREIGN KEY (id_borrower) REFERENCES assoc_partners (id),
  CONSTRAINT fk_idGame FOREIGN KEY (id_game) REFERENCES games (id)

);

CREATE TABLE assoc_users (
  id INT PRIMARY KEY,
  username VARCHAR(30) NOT NULL,
  user_password VARCHAR(30) NOT NULL
);

INSERT INTO assoc_partners VALUES (1, "Pepe");


