function getIDOfMaster(name){
    var master = db.masters.findOne({"name": name}); 
    return master._id; 
}

print("Masters: ")
db.masters.insert({"name": "Søren"});
db.masters.insert({"name": "Lasse"});
db.masters.find().forEach(printjson)


print("Characters: ")
db.characters.insert({"name": "Gargamel", "age": 200, "master": getIDOfMaster("Søren")});
db.characters.insert({"name": "Sir Smurfs-alot", "age": 37, "master": getIDOfMaster("Lasse")});
db.characters.find().forEach(printjson)
