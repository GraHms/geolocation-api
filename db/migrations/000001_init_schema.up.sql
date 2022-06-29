CREATE TABLE "geolocations" (
                                "id" bigserial PRIMARY KEY,
                                "country_code" varchar,
                                "city_name" varchar,
                                "ip_address" varchar,
                                "latitude" varchar,
                                "longitude" varchar,
                                "mystery" varchar,
                                "created_at" timestamptz DEFAULT (now())
);

CREATE INDEX ON "geolocations" ("ip_address");
CREATE INDEX ON "geolocations" ("city_name")

