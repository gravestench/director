require("util")

math.randomseed(os.time())

local second = 1000000000
local shapes = {}
local maxShapes = 50
local elapsed = 0

rw, rh = scene.sys.renderer.window.size()

function init()
    for i = 1,maxShapes,1
    do
        obj = randObject()
        shapes[i] = obj
    end
end

function update(timeDelta)
    elapsed = elapsed + timeDelta
    for _, entity in ipairs(shapes) do
        updatePositionRotation(entity)
    end
end

function updatePositionRotation(eid)
    trs, found = scene.components.transform.get(eid)
    if not found then
        running = false
        return
    end

    tx, ty, tz = trs.translation.xyz()
    tx, ty, tz = tx + 1, ty + 1, tz


    if (tx + w) > (rw + 200) then
        tx = -w
    end

    if (ty + h) > (rh + 200) then
        ty = -h
    end

    trs.translation.xyz(tx, ty, tz)

    rx, ry, rz = trs.rotation.xyz()
    ry = ry + 1

    trs.rotation.xyz(rx, ry, rz)
end

function updateOrigin(eid)
    origin, found = scene.components.origin.get(eid)
    if not found then
        return
    end

    n = elapsed / second

    ox, oy, oz = origin.xyz()
    ox = math.cos(n)
    oy = math.sin(n)

    origin.xyz(ox, oy, oz)
end

function randObject()
    x = math.random(0, 1024)
    y = math.random(0, 768)
    w = math.random(0, 150/2) * 2
    h = math.random(0, 150/2) * 2

    fill = randColor()
    stroke = randColor()

    randNumber = math.random() * 3

    if randNumber > 2 then
        obj = scene.add.rectangle(x, y, w, h, fill, stroke)
    elseif randNumber > 1 then
        obj = scene.add.circle(x, y, w/2, fill, stroke)
    else
        KEKW = "https://cdn.betterttv.net/emote/5e9c6c187e090362f8b0b9e8/3x"
        obj = scene.add.image(KEKW, x, y)
    end

    return obj
end