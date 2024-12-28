# paste.ly
Backend service immitate bit.ly for generating short URL and files uploader


## Calculate

A paste maybe around 1KB
``` sql
-- public.paste definition

-- Drop table

-- DROP TABLE public.paste;
CREATE TABLE public.paste (
	id bigserial NOT NULL,
	shortlink varchar(7) NOT NULL,
	paste_url varchar(255) DEFAULT ''::character varying NOT NULL,
	created_at timestamp DEFAULT now() NOT NULL,
	status varchar DEFAULT 'active'::character varying NULL,
	expired_at timestamp NULL,
	CONSTRAINT paste_pk PRIMARY KEY (id)
);
CREATE INDEX paste_shortlink_idx ON public.paste USING btree (shortlink);
```

the total maybe around ~0,4 KB (excluding the paste_url content)


``` sql
CREATE TABLE paste_log (
  time TIMESTAMPTZ NOT NULL,
  shortlink text not null
);


SELECT create_hypertable('paste_log', by_range('time'));
```

the total maybe around ~0,04 KB (including timescaledb metadata)

## Reference

https://stackoverflow.com/questions/742013/how-do-i-create-a-url-shortener
https://aws.github.io/aws-sdk-go-v2/docs/configuring-sdk/endpoints/