-- assume that eid is some existing entity id

function example_factory_add()
    req = scene.components.fileLoadRequest.add(eid)
end

function example_factory_get()
    req, found = scene.components.fileLoadRequest.get(eid)
end

function example_factory_remove()
    scene.components.fileLoadRequest.remove(eid)
end

function example_component_usage()
    req = scene.components.fileLoadRequest.add(eid)

    path = req.Path()
    req.path("http://example.com/image.jpg")
end