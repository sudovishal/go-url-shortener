# URL Shortener

Tech Stack- Go, Air, Postgres(For Live Deploy, probably will use NeonDB), htmx, templ
This is my [Postman Collection](https://www.postman.com/vishshinde/workspace/vishal-workspace/collection/24020296-1183a52c-243e-468d-9269-5b02c426b554?action=share&creator=24020296)

Tentative Database Schema designed
```sql
CREATE TABLE guest_users (
    id SERIAL PRIMARY KEY,
    guest_id VARCHAR(36) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE urls (
    id SERIAL PRIMARY KEY,
    original_url TEXT NOT NULL,
    short_code VARCHAR(10) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP WITH TIME ZONE,
    click_count INTEGER DEFAULT 0,
    guest_user_id INTEGER REFERENCES guest_users(id) ON DELETE SET NULL
);

CREATE INDEX idx_urls_short_code ON urls(short_code);
CREATE INDEX idx_urls_guest_user_id ON urls(guest_user_id);
```

## Reference Links 
[Why _ is used while importing SQL Drivers](https://www.calhoun.io/why-we-import-sql-drivers-with-the-blank-identifier/)
[HTMX Docs](https://htmx.org/)
[Templ Official Guide](https://templ.guide/)