-- assume that eid is some existing entity id

function example_factory_add()
    renderOrder = scene.components.renderOrder.add(eid)
end

function example_factory_get()
    renderOrder, found = scene.components.renderOrder.get(eid)
end

function example_factory_remove()
    scene.components.renderOrder.remove(eid)
end

function example_component_usage()
    renderOrder = components.renderOrder.add(eid)

    -- the highest index is rendered on top of everything else
    index = renderOrder.value()
    renderOrder.value(100) -- set the layer index to 100
end