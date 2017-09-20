import json

with open('data.json') as f:
	json_string = json.load(f)

print json_string

data =  json.dumps({
	'name' : 'The Secret of Nagas',
	'author' : 'Amish',
	'publication' : 'XYZ'
})

print type(data)

print data