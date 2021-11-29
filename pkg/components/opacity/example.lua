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
    opacity = scene.components.opacity.add(eid)

    val = opacity.value()
    opacity.value(0.5) -- set the opacity to 50%
end