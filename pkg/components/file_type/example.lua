-- assume that eid is some existing entity id

function example_factory_add()
    ft = scene.components.fileType.add(eid)
end

function example_factory_get()
    ft, found = scene.components.fileType.get(eid)
end

function example_factory_remove()
    scene.components.fileType.remove(eid)
end

function example_component_usage()
    ft, found = scene.components.fileType.get(eid)
    if ~found then return end

    if ft.type() == "image/png" then
        -- do something with png's or whatever
    end
end