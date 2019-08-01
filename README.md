## Messaging Service

This is an extermely basic messaging service that does not store messages when the server is shut down.
Messages are currently stored in a hash map


### Running

First, ensure that go is properly installed on your computer by running

```bash
go version
```

The output should show you have at least go 1.12.x installed

To start the server, enter the command `make run` in your terminal.

To build the server, you may run `make build-dev` or `make build-prod` for dev and prod environments respectively.


### Configs
All configs are stored in the `config/` folder. The config will specifically look for a file called `shared.toml` which
should contain any configs that are shared between all environments (like the app name)

The configs will then look for a file based on `ENV`. You can check your `ENV` by running the bash command `echo $ENV`. If empty,
the server will default to `dev`. If you run something like this:
```bash
ENV=staging
make run
```
The server would expect to find the file `config/staging.toml`.


### API

GET `/:channel/messages?last_id=<uuid>`
The query param `last_id` is optional

Success Response Body:
```json
{
  "messages": [
    {
      "username": "<string>",
      "message": "<string>",
      "id": "<uuid>"
    },
    // other messages
  ]
}
```


POST `/:channel/messages`

Request Body:
```json
{
  "username": "<string>",
  "message": "<string>"
}
```

Success Response Body:
```json
{
  "id": "<uuid>"
}
```

Error Responses:
```json
{
  "message": "<string>"
}
```
The error code can be found in the response headers
