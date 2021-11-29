-- assume that eid is some existing entity id

function example_factory_add()
    font = scene.components.font.add(eid)
end

function example_factory_get()
    font, found = scene.components.font.get(eid)
end

function example_factory_remove()
    scene.components.font.remove(eid)
end

function example_component_usage()
    font = components.font.add(eid)

    face = font.face()
    font.face("Comic Sans") -- lul

    pxSize = font.size()
    font.size(24) -- pixels
end