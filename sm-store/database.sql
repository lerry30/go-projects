---------- create unsigned integer type -------------
CREATE DOMAIN uint AS INTEGER CHECK (VALUE >= 0);

------ trigger function for updated_at column -------
CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
  IF NEW IS DISTINCT FROM OLD THEN
    NEW.updated_at = CURRENT_TIMESTAMP;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

------- assign trigger function to table column------

CREATE TRIGGER trigger_set_updated_at
BEFORE UPDATE ON table_name
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

--------------------- TABLES ------------------------

CREATE TABLE site_user(
    first_name VARCHAR(30) NOT NULL,
    last_name VARCHAR(30) NOT NULL,
    email_address VARCHAR(60) NOT NULL,
    phone_number VARCHAR(13) NOT NULL,
    password VARCHAR(256) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    id SERIAL PRIMARY KEY
);

-- address

CREATE TABLE country(
    country_name VARCHAR(60) NOT NULL,
    id SERIAL PRIMARY KEY
);

CREATE TABLE address(
    address_line1 VARCHAR(256) NOT NULL,
    address_line2 VARCHAR(100) NOT NULL DEFAULT '',
    city VARCHAR(100) NOT NULL,
    region VARCHAR(100) NOT NULL,
    zip_code SMALLINT NOT NULL,
    country_id INTEGER REFERENCES country(id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    id SERIAL PRIMARY KEY
);

CREATE TABLE user_address(
    user_id INTEGER REFERENCES site_user(id),
    address_id INTEGER REFERENCES address(id),
    is_default BOOLEAN NOT NULL DEFAULT FALSE,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- product

CREATE TABLE product_category(
    parent_category_id INTEGER REFERENCES product_category(id),
    category_name VARCHAR(60),
    id SERIAL PRIMARY KEY
);

CREATE TABLE product(
    category_id INTEGER REFERENCES product_category(id),
    name VARCHAR NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    product_image VARCHAR(256) NOT NULL,
    id SERIAL PRIMARY KEY
);

CREATE TABLE variation(
    category_id INTEGER REFERENCES product_category(id),
    name VARCHAR(100) NOT NULL,
    id SERIAL PRIMARY KEY
);

CREATE TABLE variation_option(
    variation_id INTEGER REFERENCES variation(id),
    value VARCHAR(100) NOT NULL,
    id SERIAL PRIMARY KEY
);

CREATE TABLE product_item(
    product_id INTEGER REFERENCES product(id),
    sku VARCHAR(60) NOT NULL,
    qty_in_stock uint NOT NULL, --uint custom datatype
    product_image VARCHAR(256) NOT NULL,
    price NUMERIC(6, 2) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    id SERIAL PRIMARY KEY
);

CREATE TABLE product_configuration(
    product_item_id INTEGER REFERENCES product_item(id),
    variation_option_id INTEGER REFERENCES variation_option(id)
);

-- payment

CREATE TABLE payment_type(
    value VARCHAR(40) NOT NULL,
    id SERIAL PRIMARY KEY
);

CREATE TABLE user_payment_method(
    user_id INTEGER REFERENCES site_user(id),
    payment_type_id INTEGER REFERENCES payment_type(id),
    provider VARCHAR(40) NOT NULL,
    credit_card_number VARCHAR(16) NOT NULL,
    credit_card_exp_month SMALLINT NOT NULL,
    credit_card_exp_day SMALLINT NOT NULL,
    credit_card_secret_code SMALLINT NOT NULL,
    name_on_card VARCHAR(100) NOT NULL,
    is_default BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    id SERIAL PRIMARY KEY
);

-- shopping cart

CREATE TABLE shopping_cart(
    user_id INTEGER REFERENCES site_user(id),
    id SERIAL PRIMARY KEY
);

CREATE TABLE shopping_cart_item(
    cart_id INTEGER REFERENCES shopping_cart(id),
    product_item_id INTEGER REFERENCES product_item(id),
    qty uint NOT NULL, -- uint custom datatype
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    id SERIAL PRIMARY KEY
);

-- order

CREATE TABLE shipping_method(
    name VARCHAR(16) NOT NULL,
    price NUMERIC(5, 2) NOT NULL,
    id SERIAL PRIMARY KEY
);

CREATE TABLE order_status(
    status VARCHAR(16) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    id SERIAL PRIMARY KEY
);

CREATE TABLE shop_order(
    user_id INTEGER REFERENCES site_user(id),
    order_date TIMESTAMP NOT NULL,
    purchase_order_number INTEGER NOT NULL,
    payment_method_id INTEGER REFERENCES user_payment_method(id),
    shipping_address INTEGER REFERENCES address(id),
    shipping_method INTEGER REFERENCES shipping_method(id),
    order_total NUMERIC(6, 2) NOT NULL,
    order_status INTEGER REFERENCES order_status(id),
    id SERIAL PRIMARY KEY
);

-- order line

CREATE TABLE order_line(
    product_item_id INTEGER REFERENCES product_item(id),
    order_id INTEGER REFERENCES shop_order(id),
    qty uint NOT NULL,
    price NUMERIC(6, 2) NOT NULL,
    id SERIAL PRIMARY KEY
);

-- reviews

CREATE TABLE user_review(
    user_id INTEGER REFERENCES site_user(id),
    ordered_product_id INTEGER REFERENCES order_line(id),
    rating_value SMALLINT NOT NULL,
    comment TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    id SERIAL PRIMARY KEY
);

-- promotion

CREATE TABLE promotion(
    name VARCHAR(40) NOT NULL,
    description TEXT NOT NULL,
    discount_rate VARCHAR(8) NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    id SERIAL PRIMARY KEY 
);

CREATE TABLE promotion_category(
    category_id INTEGER REFERENCES product_category(id),
    promotion_id INTEGER REFERENCES promotion(id)
);

---------------------------------------------------

CREATE TRIGGER trigger_set_updated_at
BEFORE UPDATE ON site_user
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER trigger_set_updated_at
BEFORE UPDATE ON address
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER trigger_set_updated_at
BEFORE UPDATE ON user_address
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER trigger_set_updated_at
BEFORE UPDATE ON product_item
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER trigger_set_updated_at
BEFORE UPDATE ON user_payment_method
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER trigger_set_updated_at
BEFORE UPDATE ON shopping_cart_item
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER trigger_set_updated_at
BEFORE UPDATE ON user_review
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

---------------------------------------------------

-- DROP TABLE promotion_category;
-- DROP TABLE promotion;
-- DROP TABLE user_review;
-- DROP TABLE order_line;
-- DROP TABLE shop_order;
-- DROP TABLE order_status;
-- DROP TABLE shipping_method;
-- DROP TABLE shopping_cart_item;
-- DROP TABLE shopping_cart;
-- DROP TABLE user_payment_method;
-- DROP TABLE payment_type;
-- DROP TABLE product_configuration;
-- DROP TABLE product_item;
-- DROP TABLE variation_option;
-- DROP TABLE variation;
-- DROP TABLE product;
-- DROP TABLE product_category;
-- DROP TABLE user_address;
-- DROP TABLE address;
-- DROP TABLE country;
-- DROP TABLE site_user;