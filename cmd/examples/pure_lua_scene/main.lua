math.randomseed(os.time())

shapes = {}
elapsed = 0

function init()
    for i = 0,100,1
    do
        e = randObject()
        shapes[i] = e
    end
end

function update(timeDelta)
    elapsed = elapsed + timeDelta
    for _, entity in ipairs(shapes) do
        updatePosition(entity)
        updateRotation(entity)
        updateOrigin(entity)
    end
end

function updatePosition(eid)
    trs, found = scene.components.transform.get(eid)
    if not found then
        return
    end

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

function updateRotation(eid)
    trs, found = scene.components.transform.get(eid)
    if not found then
        return
    end


    rx, ry, rz = trs.rotation()
    ry = ry + 1

    trs.rotation(rx, ry, rz)
end

local second = 1000000000

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

function coinFlip()
    v = math.random() * 2
    return v > 1
end

function randColorComponent()
    v = math.random()
    return string.format("%02x", v * 255)
end

function randColor()
    r, g, b = randColorComponent(), randColorComponent(), randColorComponent()
    return "#" .. r .. g .. b
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
        e = scene.add.rectangle(x, y, w, h, fill, stroke)
    elseif randNumber > 1 then
        e = scene.add.circle(x, y, w/2, fill, stroke)
    else
        KEKW = "https://cdn.betterttv.net/emote/5e9c6c187e090362f8b0b9e8/3x"
        e = scene.add.image(KEKW, x, y)
    end

    return e
end