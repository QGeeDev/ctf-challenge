CREATE TABLE "shortlinks" (
  "id" SERIAL PRIMARY KEY,
  "slug" varchar(255) NOT NULL,
  "full_link" varchar(1024) NOT NULL,
  "qr_images_id_fk" int
);

CREATE TABLE "qr_images" (
  "id" SERIAL PRIMARY KEY,
  "image_path" varchar(2083) NOT NULL
);

CREATE TABLE "flags" (
  "id" SERIAL PRIMARY KEY,
  "flag" varchar(255) NOT NULL
);

ALTER TABLE "shortlinks" ADD FOREIGN KEY ("qr_images_id_fk") REFERENCES "qr_images" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

CREATE UNIQUE INDEX "idx_shortlink_slug" ON "shortlinks" ("slug");
