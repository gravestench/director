-- assume that eid is some existing entity id

function example_factory_add()
    size = scene.components.size.add(eid)
end

function example_factory_get()
    size, found = scene.components.size.get(eid)
end

function example_factory_remove()
    scene.components.size.remove(eid)
end

function example_component_usage()
    size = scene.components.size.add(eid)

    w, h = size.size() -- get the size (width x height)
    size.size(10, 10) -- set the the size, in pixels
end