insert into users (email, first_name, last_name, role, kdf_algorithm, kdf_password_hash, kdf_salt, kdf_time_factor,
                   kdf_memory_factor, kdf_thread_factor, kdf_key_len)
        values ('me@naftuli.wtf', 'Naftuli', 'Kay', 'admin', 'argon2id-v1',
                x'3C4EC52BA9D6CA7EC545E6D33AE1D5E31AB4CE398C635CDFC56E1C252920E19E3F53AA5B4B6F0C1D55892C83BFCF07FA7198AAF6A65E86EA5D7BCCE0912E1BBB',
                x'2638B17862565857C5603CA2881946B66AB7C52134560ABB1C6B54BB838AB2C2F65A64A02D7B32C683AD2FC522271E03AC594AF38CB9EB5B4AB8BFBBB1FAE927',
                2, 65536, 4, 64);