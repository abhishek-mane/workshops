from random import randint
import pprint

# import required client package
from pymongo import MongoClient

# create connection with mongodb database
client = MongoClient('localhost', 27017)

# get/create database & collection
db = client.pymongo_db
collection = db.students

print "DB - " + db._Database__name
print "Collection - " + collection._Collection__name

# count documents
print "Number of documents : " + collection.count()

# insert some documents
for i in range(0, 10):
	document = {
		'_id' : i+1,
		'name' : 'Student'+str(i+1),
		'marks' : randint(0,101)
	}
	collection.insert(document)

# count documents
print "Number of documents : " + collection.count()

# querying documents
print collection.find_one()

print collection.find_one({
	'name' : 'Student5'
})

for document in collection.find():
	print document



# close connection with database
client.close()