
CREATE TABLE people (
	insee int NOT NULL,
	given varchar(40),
	last varchar(40)
	PRIMARY KEY (insee),
);

CREATE TABLE staff (
	insee int,
	salary int,
	company int,
	people int,
	PRIMARY KEY (insee),
	FOREIGN KEY (company) REFERENCES companies(siret)
    FOREIGN KEY (people) REFERENCES people(insee)
);

UPDATE staff SET name = (
	SELECT given, last
	FROM people
	WHERE staff.people = people.insee
);