create table users
(
    email             longtext               null,
    first_name        longtext               null,
    last_name         longtext               null,
    role              enum ('user', 'admin') null,
    kdf_algorithm     enum ('argon2id-v1')   null,
    kdf_password_hash varbinary(64)          null,
    kdf_salt          varbinary(64)          null,
    kdf_time_factor   int unsigned           null,
    kdf_memory_factor int unsigned           null,
    kdf_thread_factor tinyint unsigned       null,
    kdf_key_len       int unsigned           null,
    id                bigint unsigned auto_increment
        primary key,
    created_at        datetime(3)            null,
    updated_at        datetime(3)            null,
    deleted_at        datetime(3)            null
);

create index idx_users_deleted_at
    on users (deleted_at);

