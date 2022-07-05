CREATE TABLE "geolocations" (
                                "id" bigserial PRIMARY KEY,
                                "country_code" varchar NOT NULL ,
                                "city_name" varchar NOT NULL UNIQUE ,
                                "ip_address" varchar NOT NULL UNIQUE ,
                                "latitude" varchar NOT NULL ,
                                "longitude" varchar NOT NULL ,
                                "mystery" varchar NOT NULL ,
                                "created_at" timestamptz DEFAULT (now())
);

CREATE INDEX ON "geolocations" ("ip_address");
CREATE INDEX ON "geolocations" ("city_name")

