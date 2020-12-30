##############################
# comb.cr
def newGap(gap)
  gap /= 1.3
  gap = gap.to_i
  return 11 if gap == 9 || gap == 10
  return 1 if gap < 1
  gap
end

def combSort(a)
  len = a.size
  gap = len
  swapped = false
  loop do
    swapped = false
    gap = newGap gap
    (0...(len - gap)).each do |i|
      if a[i] > a[i + gap]
        swapped = true
        a[i], a[i + gap] = a[i + gap], a[i]
      end
    end
    break if gap <= 1 && !swapped
  end
end

N = 10000000
arr = (1..N).to_a.reverse!
combSort arr
(1...N).each do |z|
  print '!' if arr[z] < arr[z - 1]
end
