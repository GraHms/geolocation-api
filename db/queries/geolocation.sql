-- name: CreateGeolocation :one
INSERT INTO geolocations (
    country_code, city_name, ip_address, latitude, longitude, mystery
) VALUES (
             $1, $2, $3, $4, $5, $6
         )
RETURNING *;

-- name: GetGeolocation :one
SELECT * FROM geolocations
WHERE ip_address = $1 LIMIT 1;