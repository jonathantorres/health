# Workspace Rules: Auto-Approve & Execution Guidelines

- **Auto-Approve File Edits**: Automatically proceed with file modifications, code creation, and refactoring without prompting for step-by-step confirmation.
- **Autonomous Tool & Command Execution**: Execute relevant terminal commands (e.g., tests, linters, build scripts) necessary to perform and verify requested tasks.
- **Autonomous Problem Solving**: When errors occur, inspect logs, diagnose root causes, and apply fixes independently.


## Project Implementation Details

### Tech Stack
- **Framework**: Nuxt 4.6.2 (TypeScript)
- **UI**: Vue 3.5.44, Vue Router 5.4.0
- **Styling**: Tailwind CSS (to be installed)
- **Deployment**: Cloudflare Workers via Nitro preset
- **Database**: Cloudflare R1 (SQLite)
- **Dev Tools**: Wrangler
- **Auth**: Custom session-based auth with HTTP-only cookies, sessions stored in DB
- **Password Hashing**: bcryptjs

### Current State
- Minimal Nuxt starter - only `app/app.vue` with NuxtWelcome
- No Tailwind, no Cloudflare/Wrangler config
- No server API routes
- Existing Laravel app in `laravel/` for reference
- DB schema in `db_schema/health_tables_sqlite.sql`
- Screenshots in `screenshots/`

### Database Schema
From `db_schema/health_tables_sqlite.sql`:
```sql
users (id INTEGER PRIMARY KEY AUTOINCREMENT, name varchar(255), last_name varchar(255), email varchar(191) UNIQUE, password varchar(255), remember_token varchar(100), created_at timestamp, updated_at timestamp, deleted_at timestamp)
blood_pressures (id INTEGER PRIMARY KEY AUTOINCREMENT, user_id INTEGER NOT NULL DEFAULT 1, systolic INTEGER NOT NULL, diastolic INTEGER NOT NULL, pulse INTEGER NOT NULL, reading_date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, created_at timestamp NULL, updated_at timestamp NULL, deleted_at timestamp NULL, FK user_id -> users(id) ON DELETE CASCADE, INDEX on user_id)
weights (id INTEGER PRIMARY KEY AUTOINCREMENT, user_id INTEGER NOT NULL, weight double NOT NULL, entered_date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, deleted_at timestamp NULL, created_at timestamp NULL, updated_at timestamp NULL, FK user_id -> users(id) ON DELETE CASCADE, INDEX on user_id, INDEX on weight)
migrations, password_resets
```

**Need to add sessions table:**
```sql
CREATE TABLE sessions (
  id TEXT PRIMARY KEY,
  user_id INTEGER,
  ip_address TEXT,
  user_agent TEXT,
  payload TEXT NOT NULL,
  last_activity INTEGER NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
CREATE INDEX sessions_user_id_index ON sessions(user_id);
CREATE INDEX sessions_last_activity_index ON sessions(last_activity);
```

### BP Severity Logic (from Laravel BloodPressure model)
- Normal: systolic <= 120 AND diastolic <= 80 → text: 'Normal', class: 'primary'
- Pre Hypertension: (120 < systolic <= 139) OR (80 < diastolic <= 89) → text: 'Pre Hypertension', class: 'warning'
- Stage 1 Hypertension: (140 <= systolic <= 159) OR (90 <= diastolic <= 99) → text: 'Stage 1 Hypertension', class: 'danger'
- Stage 2 Hypertension: systolic >= 160 AND diastolic >= 100 → text: 'Stage 2 Hypertension', class: 'danger'
- Else: 'N/A', class: 'normal'

### UI Requirements
**Navigation Menu:**
- Left: "Health" brand linking to dashboard (/)
- "Blood Pressure" dropdown: "See all readings" (/blood-pressure), "Enter new reading" (/blood-pressure/new)
- "Weight" dropdown: "See all entries" (/weight), "Enter new weight entry" (/weight/new)
- Right: logged-in user's full name with dropdown containing "Logout"

**Routes:**
- Public: `/login` (email, password), `/register` (name, last name, email, password, confirm password)
- Protected: `/` (dashboard - 10 latest BP with cols SYS, DIA, Pulse, Date, Severity, Actions View/Edit/Delete; 10 latest weight with cols Weight, Date, Actions Edit/Delete), `/blood-pressure` (all), `/blood-pressure/new`, `/blood-pressure/:id` (view details), `/blood-pressure/:id/edit`, `/weight` (all), `/weight/new`, `/weight/:id/edit`

**Forms:**
- BP add/edit: SYS, DIA, Pulse, Reading Date
- Weight add/edit: Weight, Entry Date

### API Endpoints
- Auth: POST /api/auth/register, POST /api/auth/login, POST /api/auth/logout, GET /api/auth/me
- BP: GET /api/blood-pressure (limit/all), GET /api/blood-pressure/:id, POST /api/blood-pressure, PUT/PATCH /api/blood-pressure/:id, DELETE /api/blood-pressure/:id
- Weight: GET /api/weight, GET /api/weight/:id, POST /api/weight, PUT/PATCH /api/weight/:id, DELETE /api/weight/:id

### Key Files from Laravel (for reference)
- `laravel/routes/web.php` - route structure
- `laravel/app/Http/Controllers/AppController.php` - dashboard logic
- `laravel/app/Http/Controllers/BloodPressureController.php` - BP CRUD
- `laravel/app/Http/Controllers/WeightController.php` - Weight CRUD
- `laravel/app/BloodPressure.php` - model with severity method
- `laravel/resources/views/partials/nav.blade.php` - nav structure
- `laravel/resources/views/index.blade.php` - dashboard layout

### Implementation Completed

All core functionality has been implemented:

**Configuration:**
- `nuxt.config.ts` - Added Tailwind module and Cloudflare Nitro preset
- `wrangler.jsonc` - Cloudflare Workers configuration with R1 binding
- `schema.sql` - Complete database schema including sessions table
- Installed dependencies: tailwindcss, wrangler, @cloudflare/workers-types, bcryptjs, @types/bcryptjs

**Backend (Server API):**
- `server/utils/db.ts` - R1 database access helper
- `server/utils/auth.ts` - Session management, password hashing with bcryptjs
- `server/middleware/auth.ts` - API route protection
- Auth endpoints: register, login, logout, me
- Blood pressure CRUD endpoints with user ownership checks
- Weight CRUD endpoints with user ownership checks

**Frontend:**
- `app/app.vue` - App layout wrapper
- `app/layouts/default.vue` - Default layout with navbar
- `app/components/NavBar.vue` - Navigation with dropdowns as specified
- `app/composables/useAuth.ts` - Authentication composable
- `app/middleware/auth.ts`, `app/middleware/guest.ts` - Route guards
- `app/utils/bpSeverity.ts` - BP severity calculation logic
- Pages: login, register, dashboard (/), blood-pressure (index/new/[id]/edit), weight (index/new/[id]/edit)
- All forms match requirements (BP: SYS, DIA, Pulse, Reading Date; Weight: Weight, Entry Date; Register: Name, Last Name, Email, Password, Confirm Password)

**Features:**
- HTTP-only cookies for sessions stored in DB
- Password hashing with bcryptjs
- BP severity classification matching Laravel logic
- User ownership validation on all CRUD operations
- Responsive Tailwind CSS design
- Dashboard shows 10 latest BP and weight entries
- All actions (View/Edit/Delete) as specified

### Development Notes
- If starting the dev server for testing, always stop it before finishing. Use `pkill -f "nuxi dev"` or kill the process.
- For local development, `npm run dev` uses automatic local SQLite fallback (dev.db created from schema.sql) via better-sqlite3.
- For Cloudflare Workers testing, use `wrangler dev` with R1 bindings configured in wrangler.jsonc.
