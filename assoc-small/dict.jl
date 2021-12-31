# dict.jl
const MAX_DATA = 1_230_000
const I2CH = [
    '0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
    'a', 'B', 'c', 'D', 'e', 'F'
]
function get_first_digit(d)
    while d > 10
        d = div(d, 10)
    end
    d
end
function to_rhex(v)
    hex = ""
    while v > 0
        hex *= I2CH[v%16 + 1]
        v = div(v, 16)
    end
    hex
end
function add_or_inc!(m, key, set, inc)
    if !(key in keys(m))
        m[key] = set
        false
    else
        m[key] += inc
        true
    end
end
function main()
    m = Dict{String, Int}()
    sizehint!(m, MAX_DATA*2)

    dup1, dup2, dup3 = 0, 0, 0
    for z in reverse(1:MAX_DATA)
        val2 = MAX_DATA - z
        val3 = MAX_DATA*2 - z
        key1 = string(z)
        key2 = string(val2)
        key3 = to_rhex(val3)
        if add_or_inc!(m, key1,    z, val2); dup1 += 1 end
        if add_or_inc!(m, key2, val2, val3); dup2 += 1 end
        if add_or_inc!(m, key3, val3,    z); dup3 += 1 end
    end
    println("$dup1 $dup2 $dup3")

    total, verify, count = 0, 0, 0
    for (k, v) in m
        total  += get_first_digit(v)
        verify += length(k)
        count  += 1
    end
    println("$total $verify $count")
end
main()
