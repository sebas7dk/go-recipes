# Go Recipes

[![Join the chat at https://gitter.im/sebas7dk/go-recipes](https://badges.gitter.im/sebas7dk/go-recipes.svg)](https://gitter.im/sebas7dk/go-recipes?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

An Api to store recipes in Elastic Search


Elastic Search
---

Run Elastic Search on port 9200.

```
docker run -d -p 9200:9200 -p 9300:9300 --name elasticsearch elasticsearch
```

Installation
---

Create a .env file in the root directory with the following values:

```
# App settings
APP_PORT=:7000
APP_VERSION=v1.0

# Elastic Search
ES_DOMAIN=localhost
ES_PORT=9200
ES_INDEX=recipe_index
ES_TEST_INDEX=test_index
```

## Running the Application

To start the application you can use the `run` command.

```
$ make && make run
```

Add a recipe to the index:

```
curl -i -X POST \
  --url http://127.0.0.1:7000/api/v1.0/recipes \
  --data '{"title":"miso soup", "category":"asian", "ingredients":"miso, rice noodles, ginger", "instructions":"boil it for 60 minutes", "time": 60, "people": 2}'
```

Expected response:

```
HTTP/1.1 201 Created
```

Fetch all the recipes:

```
curl -i http://127.0.0.1:7000/api/v1.0/recipes
```

Expected response:

```
[{"id":"AVP_q_Xc6hie8WoeSicN","title":"miso soup","category":"asian","ingredients":"miso, rice noodles,noodles","instructions":"boil it for 60 minutes","time":60,"people":2}]
```
