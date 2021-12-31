##############################
# scomb.jl
function newGap(gap)
    gap = floor(Int, gap / 1.3)
    if gap == 9 || gap == 10; return 11 end
    if gap < 1; return 1 end
    gap
end
function combSort!(a, len)
    gap = len
    while true
        swapped = false
        gap = newGap(gap)
        for i in 1:(len-gap)
            j = i + gap
            if a[i] > a[j]
                swapped = true
                a[i], a[j] = a[j], a[i]
            end
        end
        if gap <= 1 && !swapped; break end
    end
end
const N = 10_000_000
function main()
    arr = [ string(N-i+1) for i in 1:N ]
    combSort!(arr, N)
    for i in 2:N
        if arr[i] < arr[i-1]
            print("!")
        end
    end
end

main()
