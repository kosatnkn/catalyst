# Catalyst
Go clean architecture RESTful API

[![Build Status](https://travis-ci.org/kosatnkn/catalyst.svg?branch=master)](https://travis-ci.org/kosatnkn/catalyst)

## Features
- A server to handle web requests
- Configuration parser
- Container for dependency injection
- Router
- Controllers
- Error handler
- Logger
- Request mapper
- Response mapper and a Transformer
- Application Metrics


## Application Initialization Process

- Parse configurations
- Resolve container
- Initialize router
- Run server


## Request Response Cycle
```text
     + -------- +                + ------- +
     | RESPONSE |                | REQUEST |
     + -------- +                + ------- +
          /\                         ||
          ||                         \/
          ||                  + ------------ +  =>  + ---------- +
          ||                  |    Router    |      | Middleware |
          ||                  + ------------ +  <=  + ---------- +
          ||                             ||
          ||                             ||
     + --------------------------- +     ||
     | Transformer | Error Handler |     ||
     + --------------------------- +     ||
                                /\       ||
                                ||       \/
                            + -------------- +  =>  + --------- +
                            |   Controller   |      | Validator |
                            + -------------- +  <=  + --------- +
                                /\       ||
                                ||       \/
                            + -------------- +
                            |    Use Case    |
                            + -------------- +
                                /\       ||
                                ||       \/
              ______________________________________________
               + ------- +    + ---------- +    + ------- +
               | Adapter |    | Repository |    | Service |
               + ------- +    + ---------- +    + ------- +
                  /\  ||         /\    ||          /\  ||
                  ||  \/         ||    \/          ||  \/
               + ------- +    + ---------- +    + ------- +
               | Library |    |  Database  |    |   APIs  |
               + ------- +    + ---------- +    + ------- +
```
