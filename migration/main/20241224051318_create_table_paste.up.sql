BEGIN;

CREATE TABLE IF NOT EXISTS public.paste (
	id bigserial NOT NULL,
	shortlink varchar(7) NOT NULL,
	paste_url varchar(255) DEFAULT ''::character varying NOT NULL,
	created_at timestamp DEFAULT now() NOT NULL,
	status varchar DEFAULT 'active'::character varying NULL,
	expired_at timestamp NULL,
	CONSTRAINT paste_pk PRIMARY KEY (id)
);
CREATE UNIQUE INDEX IF NOT EXISTS paste_shortlink_idx ON public.paste USING btree (shortlink);

COMMIT;