# Health
Web based app for health tracking and monitoring. Work in progress.

#### Installation
Just clone the repository.
```bash
git clone git@bitbucket.org:jonathantorres/health.git
```

#### Install dependencies
Install `php` dependencies with `composer` and `javascript` dependencies with `npm`.
```bash
composer install && npm install
```

#### Environment and database setup
Copy the included `.env.example` file to `.env`.
```bash
cp .env.example .env
```

Login to your local mysql installation and create your local and testing database.
```bash
mysql -u root -p ''
create database health;
create database heath_test;
```

Generate your application key
```bash
php artisan key:generate
```

And run your migrations and seeds, which will create your tables and a test user.
```bash
php artisan migrate && php artisan db:seed
```

#### Compile javascript and sass assets
```bash
npm run dev
```

#### Running the development server
Once everything is installed in order, you can run your server.
```
php artisan serve
```

Alternatively, if you have Laravel Valet installed (like I do) and configured you should be able to see your app running locally on `http://health.dev`.

#### Running unit tests
Unit tests are done with PHPUnit. Run the unit test with the following command.
```bash
vendor/bin/phpunit
```

#### Running integration tests
Integration tests are done with PHPUnit and Laravel Dusk. Run the following command to run the integration tests.
```bash
php artisan dusk
```
