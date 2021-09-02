math.randomseed(os.time())

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
    rect_eid = rectangle.new(x, y, w, h, fill, stroke)

    print("created rectangle with EID: " .. rect_eid:value())
end

for _ = 100,0,-1
do
    rrect()
end