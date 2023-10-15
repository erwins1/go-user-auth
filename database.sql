\c sawitpro;
CREATE TABLE public.users (
    user_id SERIAL PRIMARY KEY,
    full_name VARCHAR NOT NULL,
    phone_number VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    salt VARCHAR NOT NULL,
    login_count INT DEFAULT 0
);
