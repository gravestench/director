# About
`Director` is an abstraction for "scenes" written on top of raylib-go and akara.

raylib-go is a golang binding to raylib, an opengl graphic framework written in c++.

akara is a golang ECS framework

A "scene" is something that has access to various object creation factories, as well as some basic entities
that get created by default for the scene (such as a camera). Under the hood, everything is managed using ECS via akara.

ECS provides decoupling between the various systems inside of `director`, as well as any systems defined by the end-user.

# Examples
see `cmd/examples` 