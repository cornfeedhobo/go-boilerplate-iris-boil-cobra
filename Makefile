postgres:
	docker run --name postgres -dit -p '5432:5432' postgres
	sleep 15

	docker exec postgres psql -U postgres -c "CREATE USER boilerplate WITH PASSWORD 'boilerplate';"
	docker exec postgres psql -U postgres -c "CREATE DATABASE boilerplate;"

	docker exec postgres psql -U postgres --dbname boilerplate -c "CREATE SCHEMA logging;"
	docker exec postgres psql -U postgres --dbname boilerplate -c "ALTER SCHEMA logging OWNER TO postgres;"
	docker exec postgres psql -U postgres --dbname boilerplate -c "GRANT CREATE ON SCHEMA logging TO boilerplate;"
	docker exec postgres psql -U postgres --dbname boilerplate -c "GRANT INSERT ON ALL TABLES IN SCHEMA logging TO boilerplate;"
	docker exec postgres psql -U postgres --dbname boilerplate -c "CREATE TABLE logging.table_history (id SERIAL, timestamp TIMESTAMP WITH TIME ZONE DEFAULT (now() AT TIME ZONE 'utc'), schema_name TEXT NOT NULL, table_name TEXT NOT NULL, operation TEXT NOT NULL, user_name TEXT DEFAULT current_user, new_value JSON, old_value JSON);"
	docker exec postgres psql -U postgres --dbname boilerplate -c "CREATE FUNCTION table_history_trigger() RETURNS trigger AS \$$\$$ BEGIN IF TG_OP = 'INSERT' THEN INSERT INTO logging.table_history (table_name, schema_name, operation, new_value) VALUES (TG_RELNAME, TG_TABLE_SCHEMA, TG_OP, row_to_json(NEW)); RETURN NEW; ELSIF TG_OP = 'UPDATE' THEN INSERT INTO logging.table_history (table_name, schema_name, operation, new_value, old_value) VALUES (TG_RELNAME, TG_TABLE_SCHEMA, TG_OP, row_to_json(NEW), row_to_json(OLD)); RETURN NEW; ELSIF   TG_OP = 'DELETE' THEN INSERT INTO logging.table_history (table_name, schema_name, operation, old_value) VALUES (TG_RELNAME, TG_TABLE_SCHEMA, TG_OP, row_to_json(OLD)); RETURN OLD; END IF; END; \$$\$$ LANGUAGE 'plpgsql' SECURITY DEFINER;"

	docker exec postgres psql -U postgres --dbname boilerplate -c "CREATE SCHEMA boilerplate;"
	docker exec postgres psql -U postgres --dbname boilerplate -c "ALTER SCHEMA boilerplate OWNER TO postgres;"
	docker exec postgres psql -U postgres --dbname boilerplate -c "ALTER DEFAULT PRIVILEGES IN SCHEMA boilerplate GRANT ALL PRIVILEGES ON TABLES TO boilerplate;"
	docker exec postgres psql -U postgres --dbname boilerplate -c "GRANT ALL PRIVILEGES ON SCHEMA boilerplate TO boilerplate;"

	docker exec postgres psql -U postgres --dbname boilerplate -c "ALTER USER boilerplate SET search_path TO boilerplate, public;"


dev: postgres
	go build -o main .
	./main migrate init
	./main migrate up
	./main migrate current
	./main create role
	./main create user admin --password admin
	rm main
	sqlboiler --basedir ./models --output ./models --schema boilerplate --blacklist boilerplate_migration,ladon_policy_permission,ladon_policy_subject,ladon_policy_resource postgres


up: dev
	go run main.go serve


destroy:
	docker rm -f postgres
	docker volume rm $$(docker volume ls -qf dangling=true)
