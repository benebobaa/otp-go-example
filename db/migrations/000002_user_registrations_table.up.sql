CREATE TABLE IF NOT EXISTS "user_registrations" (
                                                  "id" SERIAL PRIMARY KEY,
                                                  "name" VARCHAR(255) NOT NULL,
                                                  "email" VARCHAR(255) NOT NULL,
                                                  "password" VARCHAR(255) NOT NULL,
                                                  "is_verified" BOOLEAN DEFAULT FALSE,
                                                  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "otp_users" (
                           "id" SERIAL PRIMARY KEY,
                           "ref_code" VARCHAR(16) NOT NULL,
                           "otp_value" VARCHAR(6) NOT NULL,
                           "expiration_time" TIMESTAMP,
                           "creation_time" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                            "user_regis_id" INTEGER REFERENCES "user_registrations" ("id"),
                            "user_uuid" UUID REFERENCES "users" ("id")
);

-- CREATE INDEX idx_user_regis_id ON otp_users(user_regis_id);
CREATE INDEX idx_ref_code_expiration_time ON otp_users(ref_code, expiration_time);

