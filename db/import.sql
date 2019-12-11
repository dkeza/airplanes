--
-- PostgreSQL database dump
--

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: jets; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE public.jets (
    id integer NOT NULL,
    pilot_id integer NOT NULL,
    age integer NOT NULL,
    name text NOT NULL,
    color text NOT NULL
);


ALTER TABLE public.jets OWNER TO postgres;

--
-- Name: jets_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.jets_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.jets_id_seq OWNER TO postgres;

--
-- Name: jets_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.jets_id_seq OWNED BY public.jets.id;


--
-- Name: languages; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE public.languages (
    id integer NOT NULL,
    language text NOT NULL
);


ALTER TABLE public.languages OWNER TO postgres;

--
-- Name: languages_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.languages_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.languages_id_seq OWNER TO postgres;

--
-- Name: languages_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.languages_id_seq OWNED BY public.languages.id;


--
-- Name: pilot_languages; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE public.pilot_languages (
    pilot_id integer NOT NULL,
    language_id integer NOT NULL
);


ALTER TABLE public.pilot_languages OWNER TO postgres;

--
-- Name: pilots; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE public.pilots (
    id integer NOT NULL,
    name text NOT NULL
);


ALTER TABLE public.pilots OWNER TO postgres;

--
-- Name: pilots_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.pilots_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.pilots_id_seq OWNER TO postgres;

--
-- Name: pilots_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.pilots_id_seq OWNED BY public.pilots.id;


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.jets ALTER COLUMN id SET DEFAULT nextval('public.jets_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.languages ALTER COLUMN id SET DEFAULT nextval('public.languages_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pilots ALTER COLUMN id SET DEFAULT nextval('public.pilots_id_seq'::regclass);


--
-- Data for Name: jets; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.jets (id, pilot_id, age, name, color) FROM stdin;
2	1	4	Airbus A333-300Galeb	Blue
3	2	2	Boeing 777-200	Grey
5	1	4	Learjet 23	Yellow
4	3	1	Cirrus SR22	Red
\.


--
-- Name: jets_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.jets_id_seq', 5, true);


--
-- Data for Name: languages; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.languages (id, language) FROM stdin;
1	English
2	German
3	Esperanto
\.


--
-- Name: languages_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.languages_id_seq', 2, true);


--
-- Data for Name: pilot_languages; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.pilot_languages (pilot_id, language_id) FROM stdin;
1	1
2	1
3	1
3	2
1	2
1	3
\.


--
-- Data for Name: pilots; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.pilots (id, name) FROM stdin;
1	Moore
2	Williams
3	Doe
4	Jackson
\.


--
-- Name: pilots_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.pilots_id_seq', 25, true);


--
-- Name: jets_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY public.jets
    ADD CONSTRAINT jets_pkey PRIMARY KEY (id);


--
-- Name: languages_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY public.languages
    ADD CONSTRAINT languages_pkey PRIMARY KEY (id);


--
-- Name: pilot_language_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY public.pilot_languages
    ADD CONSTRAINT pilot_language_pkey PRIMARY KEY (pilot_id, language_id);


--
-- Name: pilots_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY public.pilots
    ADD CONSTRAINT pilots_pkey PRIMARY KEY (id);


--
-- Name: jet_pilots_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.jets
    ADD CONSTRAINT jet_pilots_fkey FOREIGN KEY (pilot_id) REFERENCES public.pilots(id);


--
-- Name: pilot_language_languages_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pilot_languages
    ADD CONSTRAINT pilot_language_languages_fkey FOREIGN KEY (language_id) REFERENCES public.languages(id);


--
-- Name: pilot_language_pilots_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pilot_languages
    ADD CONSTRAINT pilot_language_pilots_fkey FOREIGN KEY (pilot_id) REFERENCES public.pilots(id);


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

