-- assume that eid is some existing entity id

function example_factory_add()
    color = scene.components.color.add(eid)
end

function example_factory_get()
    color, found = scene.components.color.get(eid)
end

function example_factory_remove()
    scene.components.color.remove(eid)
end

function example_component_usage()
    color = scene.components.color.add(eid)

    color.rgba(255, 255, 255, 255)
    r, g, b, a = color.rgba()
end