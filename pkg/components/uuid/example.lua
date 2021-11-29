-- assume that eid is some existing entity id

function example_factory_add()
    uuid = scene.components.uuid.add(eid)
end

function example_factory_get()
    uuid, found = scene.components.uuid.get(eid)
end

function example_factory_remove()
    scene.components.uuid.remove(eid)
end

function example_component_usage()
    uuid = components.uuid.add(eid)

    str = uuid.string() -- there is only a getter, no setter
    uuid.string("foo") -- does nothing
end