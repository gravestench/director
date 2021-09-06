math.randomseed(os.time())

shapes = {}

function init()
    for i = 100,0,-1
    do
        shapes[i] = randomShape()
    end
end

function update()
    for _, e in ipairs(shapes) do
        updatePosition(e:id())
    end
end

function updatePosition(eid)
    trs, found = components.transform.get(eid)
    if not found then
        return
    end

    tx, ty, tz = trs.translation()
    tx, ty, tz = tx + 1, ty + 1, tz

    if tx > 1124 then
        tx = -150
    end

    if ty > 868 then
        ty = -150
    end

    trs.translation(tx, ty, tz)
end

function coinFlip()
    v = math.random() * 2
    return v > 1
end

function rcc()
    v = math.random()
    return string.format("%02x", v * 255)
end

function randomShape()
    x = math.random(0, 1024)
    y = math.random(0, 768)
    w = math.random(0, 150/2) * 2
    h = math.random(0, 150/2) * 2

    fill = "#" .. rcc() .. rcc() .. rcc()
    stroke = "#" .. rcc() .. rcc() .. rcc()

    if coinFlip() then
        e = rectangle.new(x, y, w, h, fill, stroke)
    else
        e = circle.new(x, y, w/2, fill, stroke)
    end

    return e
end