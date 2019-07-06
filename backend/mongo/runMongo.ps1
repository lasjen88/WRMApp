New-Item -ItemType Directory -Force -Path ./data
New-Item -ItemType Directory -Force -Path ./data/db
New-Item -ItemType Directory -Force -Path ./log
mongod -f ./mongod.conf