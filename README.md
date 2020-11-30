# Pet Hospital Proof of Concept

## wat

This is a proof of concept demonstrating a few things:
- usage of the Facebook ent library (https://entgo.io/) as an abstraction over the data layer
- demonstrate scalable service patterns and boundaries

## patterns

### contracts

The application has an internal representation and an external representation of its schema.  For this demo project, those representations are the same

### domains

Each application domain 

### central "repository"

In this demo, we use `ent` as _an abstraction over the database_, not as a means 
of defining service responsibility.  Each domain is responsible for its own 
subsection of the ent client.

For example, if we have two domains `foo` and `bar` that represent `foo` and `bar` 
application objects, we can assign each domain's service a client that is strictly 
scoped to that domain.  The `foo` domain service will own a `FooClient` that was 
injected at application bootstrap.  The `FooClient` is limited to functionality 
specific to the `foo` domain, including graph traversal to other top-level objects.

This enforces a service-level separation of concerns while also maintaining 
consistency and uniformity in the data model.  The `internal/ent/schema` directory 
becomes the _source of truth_ for the application data model.
```
.
└── service
    ├── api
    │   └── public_api_contract
    └── internal
        ├── bootstrap
        │   └── application_bootstrap_functionality
        ├── domain
        │   ├── bar
        │   │   └── bar_service_implemented_here
        │   └── foo
        │       └── foo_service_implemented_here
        ├── ent
        │   ├── schema
        │   └── we_only_care_about_the_schema_folder_here
        └── services
            ├── application_wide_service_interfaces_defined_here
            ├── bar_service_interface_definition
            ├── foo
            │   ├── event
            │   │   └── event_handling_for_foo_defined_here
            │   └── transport
            │       ├── grpc
            │       ├── httpx
            │       ├── transport_for_foo_defined_here
            │       └── twirp
            └── foo_service_interface_definition
```
---
links

- https://entgo.io/
- https://www.ardanlabs.com/blog/2017/02/package-oriented-design.html

