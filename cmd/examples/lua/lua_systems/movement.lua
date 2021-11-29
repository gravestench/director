local shouldUpdatePosition = true
local shouldUpdateRotation = false
local shouldUpdateOrigin = false

local second = 1E9 -- 1 billion nanoseconds

local movingEntities

function init()
    sub = director.subscriptions.new()
    sub.require(components.Transform)

    movingEntities = sub.build()

    director.addSubscription(movingEntities)
end

function update()
    for _, eid in ipairs(movingEntities.getEntities()) do
        updateEntity(eid)
    end
end

function updateEntity(eid)
    trs, found = scene.components.transform.get(eid)
    if not found then
        return
    end

    origin, found = scene.components.origin.get(eid)
    if not found then
        return
    end

    if shouldUpdatePosition then
        tx, ty, tz = trs.translation()
        tx, ty, tz = tx + 1, ty + 1, tz

        rw, rh = scene.sys.renderer.window.size()

        if tx > rw + 150 then
            tx = -150
        end

        if ty > rh + 150 then
            ty = -150
        end

        trs.translation(tx, ty, tz)
    end

    if shouldUpdateRotation then
        rx, ry, rz = trs.rotation()
        ry = ry + 1

        trs.rotation(rx, ry, rz)
    end

    if shouldUpdateOrigin then
        n = elapsed / second

        ox, oy, oz = origin.xyz()
        ox = math.cos(n)
        oy = math.sin(n)

        origin.xyz(ox, oy, oz)
    end
end