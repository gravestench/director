require("util")

math.randomseed(os.time())

KEKW = "https://cdn.betterttv.net/emote/5e9c6c187e090362f8b0b9e8/3x"
shapes = {}
elapsed = 0

function init()
    for i = 0,100,1
    do
        obj = randObject()
        shapes[i] = obj
    end
end

function update(timeDelta)
    -- noop
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
        obj = scene.add.image(KEKW, x, y)
    end

    return obj
end