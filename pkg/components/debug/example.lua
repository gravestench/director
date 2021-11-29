-- assume that eid is some existing entity id

function example_factory_add()
    dbg = scene.components.debug.add(eid)
end

function example_factory_get()
    dbg, found = scene.components.debug.get(eid)
end

function example_factory_remove()
    scene.components.debug.remove(eid)
end

function example_component_usage()
    -- this is a tag component, so has no fields or functionality
    --
    -- this component is used merely to flag something for debug
    -- systems/scenes can check if an entity has this debug flag

    _, found = scene.components.debug.get(eid)
    if found then
       -- do debug stuff here
    end
end