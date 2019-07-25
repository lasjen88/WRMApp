# WRMApp

* Frontend: react or angular
* Backend: Golang
* DB: Mongo

## Install/Setup

### Mongo
Install a local mongo service from `https://www.mongodb.com/download-center/community`

## API
We currently have three backend services

### Character Service
|Get Characters          |                           |
|---------------|---------------------------|
| Method        | GET                       |
| URL           | ``<host>:8000/v1/characters`` |
| Description   | Get all characters from the database collection |

|Create Character          |                           |
|---------------|---------------------------|
| Method        | POST                       |
| URL           | ``<host>:8000/v1/characters`` |
| Description   | Create a charcter in the database collection |

|Get Character          |                           |
|---------------|---------------------------|
| Method        | GET                       |
| URL           | ``<host>:8000/v1/characters/<id>`` |
| Description   | Get a specific character with the id from the database collection |

|Update Character          |                           |
|---------------|---------------------------|
| Method        | PUT                       |
| URL           | ``<host>:8000/v1/characters/<id>`` |
| Description   | Update an existing character with the id in the database collection |

|Delete Character          |                           |
|---------------|---------------------------|
| Method        | DELETE                       |
| URL           | ``<host>:8000/v1/characters/<id>`` |
| Description   | Delete a specific character with the id from the database collection |


### Initiative Service

|Get Initiative          |                           |
|---------------|---------------------------|
| Method        | GET                       |
| URL           | ``<host>:8000/v1/initiative`` |
| Description   | Provides an initiative order for the provided characters. |


### Item Service
|Get Items          |                           |
|---------------|---------------------------|
| Method        | GET                       |
| URL           | ``<host>:8000/v1/items`` |
| Description   | Get all items from the database |

|Create Item          |                           |
|---------------|---------------------------|
| Method        | POST                       |
| URL           | ``<host>:8000/v1/items`` |
| Description   | Create an item in the database collection |

|Get Spells          |                           |
|---------------|---------------------------|
| Method        | GET                       |
| URL           | ``<host>:8000/v1/spells`` |
| Description   | Get all spells from the database |

|Create Spell          |                           |
|---------------|---------------------------|
| Method        | POST                       |
| URL           | ``<host>:8000/v1/spells`` |
| Description   | Create a apell in the database collection |

|Get Spells     |                           |
|---------------|---------------------------|
| Method        | GET                       |
| URL           | ``<host>:8000/v1/spells/<circle>`` |
| Description   | Get all spells within a specific circle from the database |

## TODO






### Backend 

- [ ] Model correct struct for Characters
- [ ] Move characters handler functions to package
- [ ] Make handler functions use MongoDb
