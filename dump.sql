--
-- PostgreSQL database dump
--

-- Dumped from database version 12.2
-- Dumped by pg_dump version 12.2

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: public; Type: SCHEMA; Schema: -; Owner: ferrozmatic
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO ferrozmatic;

--
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: ferrozmatic
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: products; Type: TABLE; Schema: public; Owner: ferrozmatic
--

CREATE TABLE public.products (
    id integer NOT NULL,
    model character varying(255),
    company character varying(255),
    price real
);


ALTER TABLE public.products OWNER TO ferrozmatic;

--
-- Name: products_id_seq; Type: SEQUENCE; Schema: public; Owner: ferrozmatic
--

CREATE SEQUENCE public.products_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.products_id_seq OWNER TO ferrozmatic;

--
-- Name: products_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ferrozmatic
--

ALTER SEQUENCE public.products_id_seq OWNED BY public.products.id;


--
-- Name: subscriptions; Type: TABLE; Schema: public; Owner: ferrozmatic
--

CREATE TABLE public.subscriptions (
    product_id integer,
    user_id integer
);


ALTER TABLE public.subscriptions OWNER TO ferrozmatic;

--
-- Name: users; Type: TABLE; Schema: public; Owner: ferrozmatic
--

CREATE TABLE public.users (
    id integer NOT NULL,
    username character varying(255),
    email character varying(255),
    password character varying(255),
    role character varying(255) DEFAULT 'User'::character varying,
    balance real DEFAULT 0
);


ALTER TABLE public.users OWNER TO ferrozmatic;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: ferrozmatic
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO ferrozmatic;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ferrozmatic
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: products id; Type: DEFAULT; Schema: public; Owner: ferrozmatic
--

ALTER TABLE ONLY public.products ALTER COLUMN id SET DEFAULT nextval('public.products_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: ferrozmatic
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: ferrozmatic
--

INSERT INTO public.products (id, model, company, price) VALUES (1, 'iPhone 12 mini', 'Apple', 699.99);
INSERT INTO public.products (id, model, company, price) VALUES (2, 'iPhone XR', 'Apple', 549.99);
INSERT INTO public.products (id, model, company, price) VALUES (3, 'iPhone SE 2020', 'Apple', 549.99);


--
-- Data for Name: subscriptions; Type: TABLE DATA; Schema: public; Owner: ferrozmatic
--

INSERT INTO public.subscriptions (product_id, user_id) VALUES (1, 1);
INSERT INTO public.subscriptions (product_id, user_id) VALUES (1, 2);
INSERT INTO public.subscriptions (product_id, user_id) VALUES (1, 3);
INSERT INTO public.subscriptions (product_id, user_id) VALUES (2, 2);
INSERT INTO public.subscriptions (product_id, user_id) VALUES (2, 3);
INSERT INTO public.subscriptions (product_id, user_id) VALUES (3, 2);
INSERT INTO public.subscriptions (product_id, user_id) VALUES (3, 3);
INSERT INTO public.subscriptions (product_id, user_id) VALUES (3, 1);
INSERT INTO public.subscriptions (product_id, user_id) VALUES (2, 1);


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: ferrozmatic
--

INSERT INTO public.users (id, username, email, password, role, balance) VALUES (2, 'Aza', 'azatkali.16@gmail.com', '456', 'User', 0);
INSERT INTO public.users (id, username, email, password, role, balance) VALUES (3, 'Madik', 'madiyar.uss@gmail.com', '789', 'User', 0);
INSERT INTO public.users (id, username, email, password, role, balance) VALUES (4, 'admin', 'chingiz.azimbayev@gmail.com', 'admin', 'Admin', 0);
INSERT INTO public.users (id, username, email, password, role, balance) VALUES (5, 'test', 'test@gmail.com', '123456', 'User', 0);
INSERT INTO public.users (id, username, email, password, role, balance) VALUES (1, 'Chinga', 'ferrousxp@gmail.com', '123', 'User', 610.35376);
INSERT INTO public.users (id, username, email, password, role, balance) VALUES (6, 'test1', 'test1@gmail.com', '123456', 'User', 0);


--
-- Name: products_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ferrozmatic
--

SELECT pg_catalog.setval('public.products_id_seq', 6, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ferrozmatic
--

SELECT pg_catalog.setval('public.users_id_seq', 6, true);


--
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: ferrozmatic
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: ferrozmatic
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: ferrozmatic
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: users users_username_key; Type: CONSTRAINT; Schema: public; Owner: ferrozmatic
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- Name: subscriptions subscriptions_product_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ferrozmatic
--

ALTER TABLE ONLY public.subscriptions
    ADD CONSTRAINT subscriptions_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(id) ON DELETE CASCADE;


--
-- Name: subscriptions subscriptions_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ferrozmatic
--

ALTER TABLE ONLY public.subscriptions
    ADD CONSTRAINT subscriptions_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

