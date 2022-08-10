# UrlShortener API

This is pretty simple API for url shortening. Written for portfolio, so using Clean Architecture here

# Endpoints
- POST /shorten:
  - Request body must be a valid url I can redirect to
  - Responds with id
- GET /\<id\>:
  - Redirects you to url was specified in body of request to the /shorten endpoint
  
# How to run

First, install and setup PostgreSQL

Then, run this command:

```bash
git clone https://github.com/fakefloordiv/urlshortener
```

After that, go to deployment/ and fill example.env, but with your own values

After done, just run 
```bash
source your-file.env && go run src/cmd/api
```
