-- assume that eid is some existing entity id

function example_factory_add()
    text = scene.components.text.add(eid)
end

function example_factory_get()
    text, found = scene.components.text.get(eid)
end

function example_factory_remove()
    scene.components.text.remove(eid)
end

function example_component_usage()
    text = components.text.add(eid)

    str = text.string()
    text.string("foo")
end