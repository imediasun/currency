CREATE TABLE IF NOT EXISTS transactions (
                                            id INTEGER PRIMARY KEY AUTOINCREMENT,
                                            user_id INTEGER NOT NULL,
                                            base_currency TEXT NOT NULL,
                                            target_currency TEXT NOT NULL,
                                            amount REAL NOT NULL,
                                            converted_amount REAL NOT NULL,
                                            fee REAL NOT NULL,
                                            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
