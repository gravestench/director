-- assume that eid is some existing entity id

function example_factory_add()
    opacity = scene.components.opacity.add(eid)
end

function example_factory_get()
    opacity, found = scene.components.opacity.get(eid)
end

function example_factory_remove()
    scene.components.opacity.remove(eid)
end

function example_component_usage()
    interactive = scene.components.interactive.add(eid)

    -- set up a callback on mouse click
    -- in rectangle at (10,10), with dimensions 40x40
    interactive.setMouse(constants.input.MouseButtonLeft)
    interactive.hitbox(10, 10, 40, 40)

    -- notice that it is a string for the func name being called
    interactive.callback("testCallback")
    -- see function declaration below
end

function testCallback()
    print("hello from lua")
end