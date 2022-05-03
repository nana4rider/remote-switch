DROP TABLE computers;

CREATE TABLE computers(
  id integer not null,
  name text not null,
  ssh_user text,
  ssh_key text,
  ssh_port integer,
  ip_address text not null,
  mac_address text not null,
  PRIMARY KEY (id)
);
