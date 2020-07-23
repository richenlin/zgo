package sqlxc

var schema = `
CREATE TABLE demo (
	id     int AUTO_INCREMENT PRIMARY KEY,
	parent int,
	code   varchar(255),
	name   varchar(255),
	demo   varchar(255),
	status int,

    creator varchar(255),
    updator varchar(255),
	created_at timestamp,
	updated_at timestamp
);
`
