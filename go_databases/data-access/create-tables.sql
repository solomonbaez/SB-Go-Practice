DROP TABLE IF EXISTS album;
CREATE TABLE album (
  id         INT AUTO_INCREMENT NOT NULL,
  title      VARCHAR(128) NOT NULL,
  artist     VARCHAR(255) NOT NULL,
  price      DECIMALS(5,2) NOT NULL,
  PRIMARY KEY (`id`)
)

INSERT INTO album 
  (title, artist, price)
VALUES
  ("The Beach", "Amiture", 9.00),
  ("Careful", "Boy Harsher", 8.00),
  ("Visions", "Grimes", 14.99);
