CREATE TABLE IF NOT EXISTS coffees (
  id serial PRIMARY KEY,
  name text NOT NULL,
  origin text,
  roast text
);

CREATE TABLE IF NOT EXISTS orders (
  id serial PRIMARY KEY,
  coffee_id integer NOT NULL REFERENCES coffees(id),
  quantity integer NOT NULL,
  created_at timestamptz DEFAULT now()
);
