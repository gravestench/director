-- assume that eid is some existing entity id

function example_factory_add()
    cam = scene.components.camera.add(eid)
end

function example_factory_get()
    cam, found = scene.components.camera.get(eid)
end

function example_factory_remove()
    scene.components.camera.remove(eid)
end

function example_component_usage()
    cam, found = scene.components.camera.get(eid)
    if ~found then return end

    cam.zoom(1.5)
    zoom = cam.zoom() -- should be 1.5

    cam.rotate(90)
    rot = cam.rotate() -- should be 90
end