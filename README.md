# About
`Director` is an abstraction for "scenes" written on top of raylib-go and akara.

raylib-go is a golang binding to raylib, an opengl graphic framework written in c++.

akara is a golang ECS framework

A "scene" is something that has access to various object creation factories, as well as some basic entities
that get created by default for the scene (such as a camera). Under the hood, everything is managed using ECS via akara.

ECS provides decoupling between the various systems inside of `director`, as well as any systems defined by the end-user.

# Concepts
All entities are represented with an `Entity ID`, which is just a uint64. Entities in an 
ECS framework are "composed" of various components. Some of these components have been
provided for you in `pkg/components`, and are available through the base `Scene` type.
 

## How are Scenes rendered?
TODO

## How do scenes know what to render?
TODO

## How do scenes know where to render something?
TODO

# Examples
see `cmd/examples` 