// map.v

const (
	max_data = 1230000
)

fn get_first_digit(x f64) int {
	mut d := x
	for d > 10 {
		d /= 10
	}
	return int(d)
}
fn to_rhex(x int) string {
	i2ch := [`0`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `a`, `B`, `c`, `D`, `e`, `F`]
	mut hex := ''
	mut v := x
	for v > 0 {
		hex = '${hex}${i2ch[v%16]}'
		v /= 16
	}
	return hex
}
fn set_or_inc(mut m map[string]f64, key string, set int, inc int, mut ctr &int) {
	if !(key in m) {
		m[key] = f64(set)
	} else {
		m[key] += f64(inc)
		unsafe {
			*ctr = *ctr + 1
		}
	}
}
fn main() {
	mut m := map[string]f64{}
	mut dup1 := 0
	mut dup2 := 0
	mut dup3 := 0
	for z := max_data; z > 0; z-- {
		val2 := max_data - z
		val3 := max_data*2 - z
		key1 := '${z}'
		key2 := '${val2}'
		key3 := to_rhex(val3)
		set_or_inc(mut m, key1, z, val2, mut &dup1)
		set_or_inc(mut m, key2, val2, val3, mut &dup2)
		set_or_inc(mut m, key3, val3, z, mut &dup3)
	}
	println('$dup1, $dup2, $dup3')
	mut total := 0
	mut verify := 0
	mut count := int(0)
	for k, v in m {
		total += get_first_digit(v)
		verify += k.len
		count = count + 1
	}
	println('$total, $verify, $count')
}
