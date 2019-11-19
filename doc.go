// Catalyst is a REST API base that is written in Go using the Clean Architecture paradigm.
// Out of the box it contain following features.
//
// 	- A server to handle web requests
// 	- Configuration parser
// 	- Container for dependency injection
// 	- Router
// 	- Controllers
// 	- Error handler
// 	- Logger
// 	- Request mapper
// 	- Response mapper and a Transformer
// 	- Application Metrics
//
// The application initialization process is straightforward.
//
// 	- Parse configurations
// 	- Resolve container
// 	- Initialize router
// 	- Run server
//
// The request response cycle.
//
//  + -------- +                + ------- +
//  | RESPONSE |                | REQUEST |
//  + -------- +                + ------- +
//       /\                         ||
//       ||                         \/
//       ||                  + ------------ +  =>  + ---------- +
//       ||                  |    Router    |      | Middleware |
//       ||                  + ------------ +  <=  + ---------- +
//       ||                             ||
//       ||                             ||
//  + --------------------------- +     ||
//  | Transformer | Error Handler |     ||
//  + --------------------------- +     ||
//                             /\       ||
//                             ||       \/
//                         + -------------- +  =>  + -------------------- +
//                         |   Controller   |      | Unpacker | Validator |
//                         + -------------- +  <=  + -------------------- +
//                             /\       ||
//                             ||       \/
//                         + -------------- +
//                         |    Use Case    |
//                         + -------------- +
//                             /\       ||
//                             ||       \/
//           ______________________________________________
//            + ------- +    + ---------- +    + ------- +
//            | Adapter |    | Repository |    | Service |
//            + ------- +    + ---------- +    + ------- +
//               /\  ||         /\    ||          /\  ||
//               ||  \/         ||    \/          ||  \/
//            + ------- +    + ---------- +    + ------- +
//            | Library |    |  Database  |    |   APIs  |
//            + ------- +    + ---------- +    + ------- +
package main
