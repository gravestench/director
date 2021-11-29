math.randomseed(os.time())

theSquare = 0

function init()
    fill = "#FF0000"
    scene.add.label("Press A", 300, 500, 32, "", fill)

    setupInput(makeSquare())
end

function update()
    -- nothing to do...
end

function randColorComponent()
    v = math.random()
    return string.format("%02x", v * 255)
end

function randColor()
    r, g, b = randColorComponent(), randColorComponent(), randColorComponent()
    return "#" .. r .. g .. b
end

function makeSquare()
    rw, rh = scene.sys.renderer.window.size()
    x, y = rw/2, rh/2
    w, h = 200, 200

    fill = "#FF0000"
    stroke = randColor()

    theSquare = scene.add.rectangle(x, y, w, h, fill, stroke)

    return theSquare
end

function setupInput(e)
    v = scene.components.interactive.add(e)

    v.setKey(constants.input.KeyA)
    v.callback("testCallback")
end

function testCallback()
    trs, found = scene.components.transform.get(theSquare)

    x, y, z = trs.translation()
    trs.translation(x + 20, y, z)
end