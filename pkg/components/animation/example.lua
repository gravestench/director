-- assume that eid is some existing entity id

function example_factory_add()
    anim = scene.components.animation.add(eid)
end

function example_factory_get()
    anim, found = scene.components.animation.get(eid)
end

function example_factory_remove()
    scene.components.animation.remove(eid)
end

function example_component_usage()
    scene.components.animation.add(eid)

    anim, found = scene.components.animation.get(eid)
    if ~found then return end

    anim.frame(2) -- sets frame to 2
    currentFrame = anim.frame()
end