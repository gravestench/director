-- assume that eid is some existing entity id

function example_factory_add()
    stroke = scene.components.stroke.add(eid)
end

function example_factory_get()
    stroke, found = scene.components.stroke.get(eid)
end

function example_factory_remove()
    scene.components.stroke.remove(eid)
end

function example_component_usage()
    stroke = components.stroke.add(eid)

    -- getting/setting the color
    stroke.rgba(255, 255, 255, 255)
    r, g, b, a = stroke.rgba()

    -- getting/setting the line width
    stroke.width(10) -- 10 pixel line width
    width = stroke.width()
end