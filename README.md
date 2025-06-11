# Hello World

A docker image that provides an http API with the following endpoints:

## hello world endpoint

Request
```
GET /?message=somemessage
```

Response
```
{
    "hello": "world"
    "message": "somemessage"
}
```

## echo endpoint

An echo endpoint whose message and status code can be configured. Defaults to ``200``

Request
```
GET /echo?message=somemessage&statusCode=401
```

Response
```
{
    "message": "somemessage"
}
```


## healthz endpoint

A health check endpoint that provides a json response and returns a ``200`` status code

Request
```
GET /healthz
```

Response
```
{
    "healthy": true
}
```

