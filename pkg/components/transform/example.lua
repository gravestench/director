-- assume that eid is some existing entity id

function example_factory_add()
    opacity = scene.components.opacity.add(eid)
end

function example_factory_get()
    opacity, found = scene.components.opacity.get(eid)
end

function example_factory_remove()
    scene.components.opacity.remove(eid)
end

function example_component_usage()
    updateTranslation(eid)
    updateRotation(eid)
    updateScale(eid)
end

function updateTranslation(eid)
    trs, found = scene.components.transform.get(eid)
    if not found then
        return
    end

    tx, ty, tz = trs.translation()
    tx = tx + 1
    ty = ty + 1

    trs.translation(tx, ty, tz)
end

function updateRotation(eid)
    trs, found = scene.components.transform.get(eid)
    if not found then
        return
    end

    rx, ry, rz = trs.rotation()
    ry = ry + 1

    trs.rotation(rx, ry, rz)
end

n = 0

function updateOrigin(eid)
    origin, found = scene.components.origin.get(eid)
    if not found then
        return
    end

    ox, oy, oz = origin.xyz()

    n = n + 0.001
    ox = math.cos(n)
    oy = math.sin(n)

    origin.xyz(ox, oy, oz)
end