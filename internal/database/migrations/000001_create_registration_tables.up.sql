CREATE TABLE registrations_temp (
                                    user_id UUID PRIMARY KEY,
                                    step_completed INT DEFAULT 0,

                                    first_name TEXT,
                                    last_name TEXT,
                                    email TEXT,

                                    address_street TEXT,
                                    address_house_number TEXT,
                                    address_zip_code TEXT,
                                    address_city TEXT,

                                    payment_owner TEXT,
                                    payment_iban TEXT,

                                    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE registrations_final (
                                     id UUID PRIMARY KEY,

                                     first_name TEXT,
                                     last_name TEXT,
                                     email TEXT,

                                     address_street TEXT,
                                     address_house_number TEXT,
                                     address_zip_code TEXT,
                                     address_city TEXT,

                                     payment_owner TEXT,
                                     payment_iban TEXT,
                                     payment_api_id TEXT,

                                     submitted_at TIMESTAMP DEFAULT now()
);

