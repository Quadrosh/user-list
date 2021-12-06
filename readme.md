# User manage service


Simple Golang service for creating and listing users on database.
As database used PostgreSQL

The server builds via `docker compose` and runs on port 8000.

API interacts by HTML via POST requests in JSON format


## Methods 
1. Create user

URL: localhost:8000/create

Request format:
```
{
    "id": "",
    "first_name": "Mary",
    "last_name": "Lenon",
    "age": 24,
    "recording_date": 234234
} 
```
Responce format: 
Status 200 if the user successfully added to the database

2. Find user by property

URL: localhost:8000/find

Request format:
```
{
    "property": "first_name",
    "value": "Mary"
} 
```
Valid properties : id, first_name, last_name, age. Format : string.
Values format : 

id, first_name, last_name - String

age - Number

Response format: array of users
```
[
    {
        "id": "BnDwedh7R",
        "first_name": "Mary",
        "last_name": "Lenon",
        "age": 24,
        "recording_date": 1638806356
    }
]
```
3. Filter users by `recording_date` and `age`

URL: localhost:8000/filter

Request format:

```
{
    "recording_date_from": 45,
    "recording_date_to": 0,
    "age_from": 0,
    "age_to": 29
} 
```

Response format: {"users": [array of users],"sum":[Number]}
```
{
    "users": [
        {
            "id": "BnDwedh7R",
            "first_name": "Mary",
            "last_name": "Lenon",
            "age": 24,
            "recording_date": 1638806356
        }
    ],
    "sum": 1
}
```

Bilding in Docker environment runs by `docker compose up` command.

Maintenance in local environment runs by makefile. Essensial variables passes by flags.





