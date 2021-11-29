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
An object is just what I call entities created using the scene's object factories.
All objects that are created will themselves be entities, so they are just ID's.
However, **there are a set of components that all objects will have when
created within a scene**. These components are:
* `SceneGraphNode` - A node for representing the parent-child hierarchy of the scene. The default parent node is the root node of the scene graph. 
* `Transform` - contains the position, rotation, and scale of an entity
* `Origin` - represents the origin point of the entity, using normalized values
stored in a Vector3.
The default origin point is the center of the entity, with values `(0.5, 0.5, 0.5)`
* `Opacity` - The opacity of the entity, as a normalized value between `0.0` and `1.0`
* `RenderOrder` - the render order of the entity, default is `0`. Higher numbers are rendered later (on top).
* `UUID` - contains a unique identifier as a string

## What is a Scene?
A scene is a struct that contains all of the component factories, as well as references to 
supporting systems, object factories, and entity lifecycle methods. 
Each scene is equipped to easily create objects and render them to the screen, 
without actually having to worry about rendering.

## What is a System?
In Director, there is a generic `System` that scenes are derived from. This is
the non-graphical aspect of a scene, and contains most everything a scene needs
except for the rendering functionality.

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

# Lua Scene/System API
Scenes and Systems can be created with lua scripts. The only requirement 
is that the script contain an `init` and `update` function. This is true for both Scenes and Systems.

Before script execution, a scene or system will initialize the lua state
machine with global tables that are constants, other systems, the
scene object factories, etc.

## `scene`
The `scene` table is the current scene, exposed as a lua table.
**This table is only available inside of lua scripts that are loaded as Scenes.**


This table contains bindings to the object factories, the director instance, 
component factories, and director systems (like the input, rendering, or 
events systems).

### `scene.add`
The scene's object factory. This table contains factory functions for all of the
various object types which are bundled inside of director.

All object factory functions will yield an entity ID number for the entity
that is created. This entity ID can be used for adding additional 
components, retrieving existing components, or even for deleting
the entity.

#### `scene.add.rectangle` example:
```lua
x, y = 0, 0
w, h = 10, 10 -- width and height
fill = "#FF00FF" -- magenta
stroke = "#00FF00" -- green

eid = scene.add.rectangle(x, y, w, h, fill, stroke)
```


#### `scene.add.image` example:
```lua
x, y = 0, 0

eidA = scene.add.image("path/to/file.jpg", x, y)
eidB = scene.add.image("http://example.com/example.jpg", x, y)
```

#### `scene.add.label` example:
```lua
x, y = 0, 0
size = 12 -- pixels
font = "Mono" -- font name
color = "#FF00FF" -- magenta
str = "This is a label."

eid = scene.add.label(str, x, y, size, font, color)
```

### `scene.components`
The scene component factories. These are used for creating, retrieving, and 
deleting component instances using a given entity ID.

**All component factories have the following functions**:
* `add(eid)` - adds and yields a component instance for the given entity ID.
* `remove(eid)` - deletes a component instance for the given entity ID.
* `get(eid)` - yields a component instance which can be nil **AND** a bool for whether a component was found.

Here are the components that are available for use:
* `scene.components.animation`
* `scene.components.camera`
* `scene.components.color`
* `scene.components.debug`
* `scene.components.fileLoadRequest`
* `scene.components.fileLoadResponse`
* `scene.components.fill`
* `scene.components.font`
* `scene.components.hasChildren`
* `scene.components.opacity`
* `scene.components.origin`
* `scene.components.renderOrder`
* `scene.components.size`
* `scene.components.stroke`
* `scene.components.text`
* `scene.components.transform`
* `scene.components.uuid`


# Examples
see `cmd/examples`