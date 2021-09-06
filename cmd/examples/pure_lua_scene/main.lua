math.randomseed(os.time())

function coinflip()
    v = math.random() * 2
    return v > 1
end

-- random color component
function rcc()
    v = math.random()
    return string.format("%02x", v * 255)
end

-- random rectangle
function rrect()
    x = math.random(0, 1024)
    y = math.random(0, 768)
    w = math.random(0, 150)
    h = math.random(0, 150)

    fill = "#" .. rcc() .. rcc() .. rcc()
    stroke = "#" .. rcc() .. rcc() .. rcc()
    if coinflip() then
        e = rectangle.new(x, y, w, h, fill, stroke)
    else
        e = circle.new(x, y, w/2, fill, stroke)
    end

    print("created rectangle with EID: " .. e:id())
end

for _ = 100,0,-1
do
    rrect()
end