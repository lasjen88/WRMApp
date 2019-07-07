db = connect("localhost:27017/wrm");

db.createCollection("masters");
print(db.masters)
db.createCollection("characters");
print(db.characters)
