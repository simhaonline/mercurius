CREATE TABLE `user` (
  id   binary(16) PRIMARY KEY,
  firstname varchar(255) NOT NULL,
  lastname varchar(255),
  login  varchar(255) NOT NULL
);
