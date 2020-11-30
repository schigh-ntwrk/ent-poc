# pet

This is the appointment domain.

At the top level is the `Service` struct.  
The application-wide interface contract for an `pet` service is defined 
under `internal/services`

Nested in this package is the http transport layer for `pet`.  If we were 
to use Twirp, gRPC, or some other transport medium, those implementations would 
also be housed here.  If we were to add an event handling package for 
`pet`, it would also go here.  Basically anything specific to 
`pet` _implementation_ goes here.
