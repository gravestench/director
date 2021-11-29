-- assume that eid is some existing entity id

function example_factory_add()
    origin = scene.components.origin.add(eid)
end

function example_factory_get()
    origin, found = scene.components.origin.get(eid)
end

function example_factory_remove()
    scene.components.origin.remove(eid)
end

function example_component_usage()
    origin = scene.components.origin.add(eid)

    ox, oy = origin.xy()

    -- setting the origin point
    origin.xy(0, 0) -- top-left
    origin.xy(1, 0) -- top-right
    origin.xy(0.5, 0.5) -- center
    origin.xy(1, 1) -- bottom-right
    origin.xy(1, 0) -- top-right
end