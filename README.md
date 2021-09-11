[![Go Report Card](https://goreportcard.com/badge/github.com/Gravestench/Director)](https://goreportcard.com/report/github.com/Gravestench/Director)
[![Twitch Status](https://img.shields.io/twitch/status/gravestench_?style=social)](https://www.twitch.tv/gravestench_)

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

## What is an entity?
An entity is just an identifier, a `Uint64`.

## What is a component?
A component is a struct that either wraps another struct or contains primitive data types. An entity may only have one 
instance of a specific component type. Components are generally used to describe features that an
entity possess.

## What is an object?
This is just how I refer to an entity that is composed of several required components (like an image, or a label).
It's just a term I use, I don't know if it has much significance.

## What is a scene?
A scene is a struct that contains all of the component factories, as well as references to 
supporting systems, object factories, and entity lifecycle methods. 
Each scene is equipped to easily create objects and render them to the screen, 
without actually having to worry about rendering.

## How are Scenes rendered?
Scenes have at least one "viewport," but can have more.
A viewport contains a reference to only one camera. 
Both the camera and the viewport have their own textures. 
The scene will render from  the camera's perspective, onto the camera's texture.
The scene will then copy the camera's texture onto the viewport, stretching it to fille the entire viewport.  

Viewports can be treated like any visible entity (they can be transformed, scaled, rotated, have their opacity set, etc.)

## How do scenes know what to render?
Each scene has its own set of object factories, such as the object factory for creating images.
When any object factory is used to create an entity, the entity ID is added to the scene's "render list."

# Examples
see `cmd/examples` 