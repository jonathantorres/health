# Health
[![Build Status](https://travis-ci.org/jonathantorres/health.svg?branch=master)](https://travis-ci.org/jonathantorres/health)

Web based app for health tracking and monitoring. Work in progress.

#### Installation
Clone the repository:
```bash
git clone git@bitbucket.org:jonathantorres/health.git
```

Go is required to build form source, to build the binary, `cd` into the project root and run:
```bash
go build
```

#### Install front-end dependencies
Install these by running the `npm` command
```bash
npm install
```

You can compile the css yourself of use the compiled css that is included in the repository, if you wish to compile the css yourself, run:
```bash
npm run sass
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

#### Running the server
Once everything is installed in order, you can run your server.
```bash
./health
```
