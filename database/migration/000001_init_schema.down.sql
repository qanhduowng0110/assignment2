CREATE TABLE "Earthquakes" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "feature_id" VARCHAR(255) NOT NULL,
  "magnitude" DECIMAL(4,2) NOT NULL,
  "occur_time" TIMESTAMP NOT NULL,
  "update_time" TIMESTAMP NOT NULL,
  "url" VARCHAR(255) NOT NULL,
  "detail_url" VARCHAR(255) NOT NULL,
  "status" VARCHAR(50) NOT NULL,
  "tsunami" INT NOT NULL,
  "sig" INT NOT NULL,
  "net" VARCHAR(20) NOT NULL,
  "code" VARCHAR(20) NOT NULL,
  "nst" INT NOT NULL,
  "dmin" DECIMAL(7,4) NOT NULL,
  "rms" DECIMAL(3,2) NOT NULL,
  "gap" DECIMAL(5,2) NOT NULL,
  "mag_type" VARCHAR(10) NOT NULL,
  "earthquake_type" VARCHAR(50) NOT NULL
);

CREATE TABLE "Geometry" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "earthquake_id" INT NOT NULL,
  "longitude" FLOAT NOT NULL,
  "latitude" FLOAT NOT NULL,
  "depth" DECIMAL(6,2) NOT NULL,
  "place" VARCHAR(255) NOT NULL
);

CREATE TABLE "Associated_events" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "earthquake_id" INT NOT NULL,
  "associate_id" INT NOT NULL
);

CREATE TABLE "Feature_types" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "earthquake_id" INT NOT NULL,
  "feature_product_type" VARCHAR(255) NOT NULL
);

CREATE TABLE "Felt_reports" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "earthquake_id" INT NOT NULL,
  "felt" INT,
  "cdi" DECIMAL(2,1),
  "mmi" DECIMAL(2,1),
  "alert" VARCHAR(50)
);

CREATE TABLE "Event_types" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "earthquake_id" INT,
  "types" VARCHAR(255)
);

CREATE TABLE "API_request_logs" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "request_datetime" TIMESTAMP NOT NULL,
  "request_parameters" JSON NOT NULL,
  "request_body" JSON,
  "request_headers" JSON,
  "request_metadata" JSON,
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
);

ALTER TABLE "Geometry" ADD FOREIGN KEY ("earthquake_id") REFERENCES "Earthquakes" ("id");

ALTER TABLE "Associated_events" ADD FOREIGN KEY ("earthquake_id") REFERENCES "Earthquakes" ("id");

ALTER TABLE "Associated_events" ADD FOREIGN KEY ("associate_id") REFERENCES "Earthquakes" ("id");

ALTER TABLE "Feature_types" ADD FOREIGN KEY ("earthquake_id") REFERENCES "Earthquakes" ("id");

ALTER TABLE "Felt_reports" ADD FOREIGN KEY ("earthquake_id") REFERENCES "Earthquakes" ("id");

ALTER TABLE "Event_types" ADD FOREIGN KEY ("earthquake_id") REFERENCES "Earthquakes" ("id");