-- assume that eid is some existing entity id

function example_factory_add()
    fill = scene.components.fill.add(eid)
end

function example_factory_get()
    fill, found = scene.components.fill.get(eid)
end

function example_factory_remove()
    scene.components.fill.remove(eid)
end

function example_component_usage()
    fill, found = scene.components.fill.get(eid)
    if ~found then return end

    fill.rgba(255, 255, 255, 255)
    r, g, b, a = fill.rgba()
end