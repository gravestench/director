-- assume that eid is some existing entity id

function example_factory_add()
    hasChildren = scene.components.hasChildren.add(eid)
end

function example_factory_get()
    hasChildren, found = scene.components.hasChildren.get(eid)
end

function example_factory_remove()
    scene.components.hasChildren.remove(eid)
end

function example_component_usage()
    parent = scene.newEntity()
    reference = components.hasChildren.add(parent)

    child1 = newEntity()
    child2 = newEntity()
    child3 = newEntity()

    reference.children({ child1, child2, child3})
end