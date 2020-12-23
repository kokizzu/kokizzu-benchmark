//////////////////////////////
// scomb.v
fn new_gap(gap1 int) int {
	mut gap := int(f64(gap1) / 1.3)
	if gap == 9 || gap == 10 {
		return 11
	}
	if gap < 1 {
		return 1
	}
	return gap
}

fn comb_sort(mut a []string) {
	xlen := a.len
	mut gap := xlen
	mut swapped := false
	for {
		swapped = false
		gap = new_gap(gap)
		for i := 0; i < xlen - gap; i++ {
			if a[i] > a[i + gap] {
				swapped = true
				a[i], a[i + gap] = a[i + gap], a[i]
			}
		}
		if !((gap > 1) || swapped) {
			break
		}
	}
}

const (
	n = 10000000
)

fn main() {
	mut arr := []string{len: n}
	for z := 0; z < n; z++ {
		arr[z] = '${n - z}'
	}
	comb_sort(mut arr)
	for z := 1; z < n; z++ {
		if arr[z] < arr[z - 1] {
			println('!')
		}
	}
}
