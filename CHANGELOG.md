# Changelog

#### vx.x.x `xxxx-xx-xx`
- New Feature
  - Update Blood Pressure readings
  - Delete Blood Pressure readings

- Improvements
  - Removed composer.lock and package-lock.json files, they are being ignored now
  - Add TravisCI build

#### v0.3.0 `2017-12-18`
- New Feature
  - Displaying and entering weight entries

- Improvements
  - Blood Pressure records are now soft deletable
  - Added an MIT license and contribution guidelines
  - A dump of the database is saved on every deployment

#### v0.2.0 `2017-09-26`
- Bug fix
  - Fixed url to redirect to if the user is already authenticated

- New Feature
  - Implemented password resets for users

- Improvements
  - Browser tests can be run on local environment
  - Add acceptance tests for blood pressure readings
  - Add initial docs for installation and running the application locally
  - Add initial TravisCI file
  - Added color code for the severity on the blood pressure readings

#### v0.1.2 `2017-09-07`
- Improvements:
  - Ignoring compiled css and js files
  - Including footer on registration and login screens
  - Version string appears in the footer
  - Added a deployment script

#### v0.1.1 `2017-08-10`
- Improvements:
  - Added a simple footer
  - Users and blood pressures are now related models
  - Showing and creating blood pressures by the authenticated user
  - Adding flash messages for errors

#### v0.1.0 `2017-07-09`
- First release.
