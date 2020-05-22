# webapp
The mercurius web application is a single page application 
using the wasm compilation target using [forms](https://github.com/golangee/forms).
 
This is a separate go module and it should be discussed, if
a nested module is a clean way to represent that fact. Actually
one could mix both code bases (the *presentation* layer) into
each service package, but that has the following drawbacks:
* There is no easy way to perform the conditional compiling. 
Annotating each file with +build wasm and +build !wasm seems
to be very verbose.
* IDEs are not prepared to work with different targets at the
same time.
* Compiling only all files by explicitly giving them to 
*go build* does not work, because the go tool only accepts
files from the same package.
* Compiling everything together results in a very large and 
bloated *wasm* file, so this is not an option either. Perhaps
this may get better with an improved linker, but not today.
* A developer can no longer see and decide in a simple way
which files need to go into the frontend and which into the backend 
where backend code makes its way into the frontend build.
This mashup results in very big and bloated builds (and sometimes
insecure), as you can often see in Java GWT projects.

Therefore, the final decision is to create a separate clean 
frontend-only module.