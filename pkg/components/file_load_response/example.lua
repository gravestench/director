-- assume that eid is some existing entity id

function example_factory_add()
    res = scene.components.fileLoadResponse.add(eid)
end

function example_factory_get()
    res, found = scene.components.fileLoadResponse.get(eid)
end

function example_factory_remove()
    scene.components.fileLoadResponse.remove(eid)
end

function example_component_usage()
    res, found = scene.components.fileLoadResponse.get(eid)
    if ~found then return end

    bytes = res.data() -- raw byte stream
    for idx, byte in ipairs(bytes) do --[[ something... ]] end

    scene.components.fileLoadResponse.remove(eid)
end