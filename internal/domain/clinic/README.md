# clinic

This is the appointment domain.

At the top level is the `Service` struct.  
The application-wide interface contract for an `clinic` service is defined 
under `internal/services`

Nested in this package is the http transport layer for `clinic`.  If we were 
to use Twirp, gRPC, or some other transport medium, those implementations would 
also be housed here.  If we were to add an event handling package for 
`clinic`, it would also go here.  Basically anything specific to 
`clinic` _implementation_ goes here.
