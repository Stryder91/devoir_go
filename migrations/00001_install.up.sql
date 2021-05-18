CREATE TABLE companies (
  siret int NOT NULL,
  date date,
  total_salary int,
  PRIMARY KEY (siret)
);
CREATE TABLE staff (
  insee int(13),
  name varchar(40),
  salary int,
  staff int,
  PRIMARY KEY (insee),
  FOREIGN KEY (staff) REFERENCES companies(siret)
);