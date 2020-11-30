# bootstrap

This package exists to initialize all resources and inject all dependencies.  
We use the google wire dependency injector, partly because it makes DI a little 
easier, but also because it forces us to be consistent across all our dependencies 
regarding instantiation.

https://github.com/google/wire
