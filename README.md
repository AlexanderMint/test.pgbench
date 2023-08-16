run:

```bash
docker compose run app init risk_address1
docker compose run app init risk_address2


docker compose run app insert risk_address1 10000000
docker compose run app select risk_address1

docker compose run app insert risk_address2 10000000
docker compose run app select risk_address2
```

init:

```sql
    DROP TABLE IF EXISTS risk_address1;
    DROP TABLE IF EXISTS risk_address2;

	CREATE TABLE risk_address1
	(
			address VARCHAR          NOT NULL,
			version BIGINT DEFAULT 1 NOT NULL,
			PRIMARY KEY (address, version)
	);
	
	CREATE TABLE risk_address2
	(
			id      BIGSERIAL CONSTRAINT risk_address_pk PRIMARY KEY,
			address VARCHAR          NOT NULL,
			version BIGINT DEFAULT 1 NOT NULL
	);
	CREATE UNIQUE INDEX risk_address2_address_version_uindex
			ON risk_address2 (address ASC, version DESC);
```