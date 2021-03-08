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
// In the context of `Catalyst` we use a concept called `Communication Channels` (simply channels)
// to define ways in which you can communicate with the microservice
// (do not confuse these with `channels` in `Go`, which is an entirely different thing).
//
// A `channel` in Catalyst is a package inside the `channels` directory. This package consists of
// all the logic needed to handle communication with the client side.
//
// Out of the box Catalyst contain two such channels.
// - http (to handle REST web requests)
// - metrics (to publish application metrics)
//
// What makes Catalyst a REST API is this `http` package which handles the complete lifecycle of REST web requests.
//
//                            + ------- +           + -------- +
//                            | REQUEST |           | RESPONSE |
//                            + ------- +           + -------- +
//                                ||                     /\
//                                \/                     ||
//                         + ------------ +              ||
//                         |  Middleware  |              ||
//                         + ------------ +              ||
//                                ||                     ||
//                                \/                     ||
//                         + ------------ +              ||
//                         |    Router    |              ||
//                         + ------------ +              ||
//                                    ||                 ||
//                                    ||                 ||
//                                    ||   + --------------------------- +
//                                    ||   | Transformer | Error Handler |
//                                    ||   + --------------------------- +
//                                    ||    /\
//                                    \/    ||
// + -------------------- +  =>  + -------------- +
// | Unpacker | Validator |      |   Controller   |
// + -------------------- +  <=  + -------------- +
//                                   ||       /\
//                                   \/       ||
//                               + -------------- +
//                               |    Use Case    |
//                               + -------------- +
//                                   ||       /\
//                                   \/       ||
//                       _____________________________________
//                           + ---------- +    + ------- +
//                           | Repository |    | Service |
//                           + ---------- +    + ------- +
//                             ||    /\          ||  /\
//                             \/    ||          \/  ||
//                           + ---------- +    + ------- +
//                           |  Database  |    |   APIs  |
//                           + ---------- +    + ------- +
//
// Likewise the `metrics` channel exposes an endpoint to let `Prometheus` scrape application metrics.
//
// You can add other `communication channels` to leverage a project based on Catalyst.
// For an example a `stream` package can be added to communicate with a streaming platform like `Kafka`.
// Or an `mqtt` package can be added to communicate with `IoT` devices.
package main
