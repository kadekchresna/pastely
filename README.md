
## Calculate

A paste maybe around 1KB
``` sql
-- public.paste definition

-- Drop table

-- DROP TABLE public.paste;

CREATE TABLE public.paste (
	shortlink varchar(7) NOT NULL, -- ~7 bytes
	expiration_length_in_minutes int4 DEFAULT 0 NOT NULL, -- ~4 bytes
	paste_url varchar(255) DEFAULT ''::character varying NOT NULL, -- ~255 bytes
	created_at timestamp DEFAULT now() NOT NULL, -- ~64-bit integer 8 bytes

	CONSTRAINT paste_pk PRIMARY KEY (shortlink)
);
```

the total maybe arround ~1,27 KB