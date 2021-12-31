function main()
    last = 3
    res = [last]
    while true
        last += 2
        prime = true
        for v in res
            if v * v > last; break end
            if last % v == 0
                prime = false
                break
            end
        end
        if prime
            push!(res, last)
            if length(res) % 100_000 == 0; println(last) end
            if last > 9_999_999; break end
        end
    end
end
main()
