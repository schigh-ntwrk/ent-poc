# httpx

This package defines the application-wide http interface.  Each domain is 
responsible for defining its own routes and handlers.  This top-level http 
package is modular and unaware of the underlying domain implementations by design.
