/*
Director is a graphical scene abstraction, implemented using the `akara` ECS framework and
`raylib-go`, the golang port of `raylib`. A scene is an extension of a system (from the ECS design pattern),
but with access to a bunch of object factories and generic rendering facilities built in.

Inside of the root of this repository, the top-level API for director is declared. The files that contain these
declarations are all prefixed with `api_`. The API declarations are just aliases to things inside
of the `pkg` directory, and are intended to reduce the cognitive load on the end-user. If these
declarations were not provided here, then an end-user would need to know about the internals of director,
and where everything is defined inside of `pkg`.

The real boon is that we can coherently separate
things inside of pkg, as well also coherently organize them for use by the end-user here.
The downside to declaring the API this way is that it is tightly coupled to `pkg`, and changes inside of
`pkg` will likely require changes to the API, but this is not at all surprising.
*/
package director
