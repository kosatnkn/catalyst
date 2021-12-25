// Catalyst started out as a microservice base that can be used to create REST APIs.
// It contains many essential parts that you would need for a microservice such as,
// - Configurability
// - A basic dependency injection mechanism
// - Request response cycle handling
// - Structure and field validations
// - Error handling
// - Logging
// - Database resource management
// - Application metrics
//
// Written using the Clean Architecture paradigm it offers clean separation between
// business (domain) logic and facilitation logic.
//
// In the context of `Catalyst` we use a concept called `Transport mediums` to define ways in which you can communicate
// with the microservice.
//
// A package inside the `transport` directory consists of all the logic needed to handle communication with the
// outside world using one type of transport medium.
//
// Out of the box, Catalyst contain two such transport mediums.
// - http (to handle REST web requests)
// - metrics (to expose application metrics)
//
// What makes Catalyst a REST API is this `http` package which handles the complete lifecycle of REST web requests.
//
//                                + ------- +           + -------- +
//                                | REQUEST |           | RESPONSE |
//                                + ------- +           + -------- +
//                                    ||                     /\
//                                    \/                     ||
//                             + ------------ +              ||
//                             |  Middleware  |              ||
//                             + ------------ +              ||
//                                    ||                     ||
//                                    \/                     ||
//                             + ------------ +              ||
//                             |    Router    |              ||
//                             + ------------ +              ||
//                                        ||                 ||
//                                        ||                 ||
//                                        ||   + --------------------------- +
//                                        ||   | Transformer | Error Handler |
//                                        ||   + --------------------------- +
//                                        ||    /\
//                                        \/    ||
//     + -------------------- +  =>  + -------------- +
//     | Unpacker | Validator |      |   Controller   |
//     + -------------------- +  <=  + -------------- +
//                                       ||       /\
//                                       \/       ||
//                                   + -------------- +
//                                   |    Use Case    |
//                                   + -------------- +
//                                       ||       /\
//                                       \/       ||
//                           _____________________________________
//                               + ---------- +    + ------- +
//                               | Repository |    | Service |
//                               + ---------- +    + ------- +
//                                 ||    /\          ||  /\
//                                 \/    ||          \/  ||
//                               + ---------- +    + ------- +
//                               |  Database  |    |   APIs  |
//                               + ---------- +    + ------- +
//
// Likewise the `metrics` transport medium exposes an endpoint to let `Prometheus` scrape application metrics.
//
// You can add other transport mediums to leverage a project based on Catalyst.
//
// For an example a `stream` package can be added to communicate with a streaming platform like `Kafka`.
// Or an `mqtt` package can be added to communicate with `IoT` devices.
package main
