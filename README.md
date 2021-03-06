# App Invite Service

## Deployed Solution

The API is deployed at [https://pulseid-fur355ca3q-uc.a.run.app](https://pulseid-fur355ca3q-uc.a.run.app)  
The API documentation is deployed at [https://pulseid.web.app](https://pulseid.web.app)

## Requirements

Create a service that will facilitate the invite token generation and validation for the Catalyst Experience App.

An invite token is a 6 to 12 digit alphanumeric string that app admins can share with potential customers.

The user workflow is as follow

1. The App Admin generates an invitation token using a web app
2. The invite token is then used to login into the Catalyst Experience App
   The outcome of the action above can be either a successful login or the user is
   asked to retry.

### Basic functional requirements

1. The APIs should be RESTfull
2. The admin endpoints should be authenticated. Propose an easy auth mechanism
3. Invite tokens to expire after 7 days
4. Invite tokens can be recalled (disabled)
5. A public endpoint for validating the invite token

### Nice to have functional requirements

1. The invite token validation logic needs to be throttled (limit the requests coming from a specific client)
2. An admin can get an overview of active and inactive tokens

### Basic nonfunctional requirements

1. Design and document the APIs that will facilitate the workflow outlined above
2. Develop the API in GO
3. Use any framework or library that will help you develop the solution faster
4. Make sure your code is well-formatted, clean, and follows best practices
5. Separate concerns
6. Write testable code
7. Use in-memory storage for the tokens

### Nice to have nonfunctional requirements

1. Document the APIs in Swagger or a similar tool
2. Write functional code
3. Tests, all levels of them
4. Use an actual DB (MySQL is preferred)
5. Provide deployment instructions

## Implementation Server

The app follows the [Twelve-Factor-App-Principles](https://12factor.net/)

### Environment variables

```sh
# database configurations
export PULSE_ENV="dev"
export PULSE_DB_USER="root"
export PULSE_DB_PWD=""
export PULSE_DB_NAME="pulse"
export PULSE_DB_PORT="3306"
export PULSE_DB_CLOUD=false
export PULSE_DB_HOST="127.0.0.1"
export PULSE_DB_TIMEZONE="Africa/Nairobi"
export PULSE_DB_INSTANCE_CONNECTION_NAME="theta-outrider-342406:us-central1:wallet"
```

### Authentication

The service uses the basic Authorization scheme.

username: myusername  
password: pass@123

Use base64 tool to encode the username and password joined by a colon

```sh
# base64 encode the username and password
echo "myusername:pass@123" | base64 -
bWFsdWtpbXV0aHVzaTpwYXNzMTIzCg==
```

For the endpoints that require authentication, use the generated string. Send the header

```txt
Authorization: Basic bWFsdWtpbXV0aHVzaTpwYXNzMTIzCg==
```

### /generate

```sh
# Send request
curl -H "Authorization: Basic bWFsdWtpbXV0aHVzaTpwYXNzMTIzCg==" http://localhost:8080/generate
```

## API Documentation

To generate the documentation for the swagger specification

```sh
# generate an html documentation of your api
java \
    -jar $HOME/programs/swagger-codegen-cli.jar generate \
    -i swagger.yml \
    -l html2 \
    -o docs
```

### Host the documentation

The folder api-docs can be deployed. as a documentation for the API.  
In this example I will host the documentation on firebase, follow the documentation at [firebase-hosting][1]

```sh
# deploy
firebase deploy --only hosting:pulseid
```

The api documentation is deployed at [https://pulseid.web.app](https://pulseid.web.app)

## References

[https://firebase.google.com/docs/hosting](https://firebase.google.com/docs/hosting)

[1]: https://firebase.google.com/docs/hosting
