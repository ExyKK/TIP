CREATE TABLE suppliers(
	id bigint PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
	name varchar(255) NOT NULL,
	details text
);

CREATE TABLE manufacturers(
	id bigint PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
	name varchar(255) NOT NULL,
	details text
);

CREATE TABLE categories(
	id bigint PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
	title varchar(255) NOT NULL
);

CREATE TABLE images(
	id bigint PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
	image1 varchar(255),
	image2 varchar(255),
	image3 varchar(255),
	image4 varchar(255),
	image5 varchar(255)
);

CREATE TABLE products(
	id bigint PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
	supplier_id bigint REFERENCES suppliers(id) NOT NULL,
	manufacturer_id bigint REFERENCES manufacturers(id) NOT NULL,
	category_id bigint REFERENCES categories(id) NOT NULL,
	images_id bigint REFERENCES images(id),
	internal_code bigint NOT NULL,
	part_number varchar(255) NOT NULL,
	title varchar(255) NOT NULL,
	description text,
	technical_info text,
	price decimal(12, 2) NOT NULL,
	status varchar(255) NOT NULL
);

CREATE TABLE users(
	id bigint PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
	login varchar(255) NOT NULL,
	password varchar(255) NOT NULL,
	name varchar(255) NOT NULL,
	surname varchar(255) NOT NULL,
	middle_name varchar(255),
	email varchar(255) NOT NULL,
	phone_number varchar(255)
);

CREATE TABLE reviews(
	id bigint PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
	user_id bigint REFERENCES users(id) NOT NULL,
	product_id bigint REFERENCES products(id) NOT NULL,
	title varchar(255) NOT NULL,
	body text NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP,
	rating smallint NOT NULL
);

CREATE TABLE carts(
	id bigint PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
	user_id bigint REFERENCES users(id) NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE carts_products(
	id bigint PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
	cart_id bigint REFERENCES carts(id) NOT NULL,
	product_id bigint REFERENCES products(id) NOT NULL,
	quantity bigint NOT NULL
);

CREATE TABLE shipping_addresses(
	id bigint PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
	user_id bigint REFERENCES users(id) NOT NULL,
	postal_code bigint NOT NULL,
	country varchar(255),
	province varchar(255),
	city varchar(255),
	address varchar(255) NOT NULL
);

CREATE TABLE orders(
	id bigint PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
	user_id bigint REFERENCES users(id) NOT NULL,
	shipping_address_id bigint REFERENCES shipping_addresses(id),
	code varchar(255) NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP,
	estimated_delivery_time timestamp,
	status varchar(255) NOT NULL,
	details text
);

CREATE TABLE orders_products(
	id bigint PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
	order_id bigint REFERENCES orders(id) NOT NULL,
	product_id bigint REFERENCES products(id) NOT NULL,
	quantity bigint NOT NULL
);

CREATE TABLE payments(
	id bigint PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
	order_id bigint REFERENCES orders(id) NOT NULL,
	payment_method varchar(255) NOT NULL,
	amount decimal(12, 2) NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP,
	status varchar(255) NOT NULL
);

CREATE TABLE warehouses(
	id bigint PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
	name varchar(255) NOT NULL,
	address varchar(255) NOT NULL,
	status varchar(255) NOT NULL,
	details text
);

CREATE TABLE warehouses_products(
	id bigint PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
	warehouse_id bigint REFERENCES warehouses(id) NOT NULL,
	product_id bigint REFERENCES products(id) NOT NULL,
	quantity_in_stock bigint NOT NULL
);
