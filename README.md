# WRMApp

* Frontend: react or angular
* Backend: Golang
* DB: Mongo

## Install/Setup

### Mongo
*Terminal 1: Powershell (For now)*

`cd WRMApp\backend\mongo`

Start Mongo: `.\runMongo.ps1`

*Terminal 2: Any Term*

`cd WRMApp/backend/mongo`

Connect to mongo: `mongo`

Run Setup: `load("setup.js");`

_Optional Populate with test-data:_: `load("populate.js");`

## TODO

### Backend 

- [ ] Model correct struct for Characters
- [ ] Move characters handler functions to package
- [ ] Make handler functions use MongoDb
