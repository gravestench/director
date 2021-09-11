math.randomseed(os.time())

theSquare = 0

function init()
    e = makeObject()
    setupInput(e)
end

function update()

end

function randColorComponent()
    v = math.random()
    return string.format("%02x", v * 255)
end

function randColor()
    r, g, b = randColorComponent(), randColorComponent(), randColorComponent()
    return "#" .. r .. g .. b
end

function makeObject()
    x, y = 200, 200
    w, h = 200, 200

    fill = "#FF0000"
    stroke = randColor()

    theSquare = rectangle.new(x, y, w, h, fill, stroke)
    label.new("Press A", 300, 500, 32, "", fill)

    return theSquare
end

function setupInput(e)
    v = components.interactive.add(e)

    v.setKey(constants.input.KeyA)
    v.callback("testCallback")
end

function testCallback()
    trs, found = components.transform.get(theSquare:id())
    if not found then
        return
    end

    x, y, z = trs.translation()
    trs.translation(x + 20, y, z)
end