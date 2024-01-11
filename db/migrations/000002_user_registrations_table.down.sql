-- Drop the index
DROP INDEX IF EXISTS idx_ref_code_expiration_time;

-- Drop the tables in reverse order
DROP TABLE IF EXISTS otp_users;
DROP TABLE IF EXISTS user_registrations;
