//////////////////////////////
// comb.rs
fn new_gap(gap: usize) -> usize {
    let ngap = ((gap as f64) / 1.3) as usize;
    if ngap == 9 || ngap == 10 { return 11 }
    if ngap < 1 { return 1 }
    return ngap
}
fn comb_sort(a: &mut [f64]) {
    let xlen = a.len();
    let mut gap = xlen;
    let mut swapped : bool;
    let mut temp : f64;    
    loop {
        swapped = false;
        gap = new_gap(gap);
        for i in 0..(xlen-gap) {
            if a[i] > a[i+gap] {
                swapped = true;
                temp = a[i];
                a[i] = a[i+gap];
                a[i+gap] = temp;
            }
        }
        if !(gap > 1 || swapped) { break }
    }
}
const N : usize = 10000000;
fn main() {
    let mut arr: Vec<f64> = std::iter::repeat(0.0).take(N).collect();
    for z in 0..(N) { arr[z] = (N - z) as f64; }
    comb_sort(arr.as_mut_slice());
    for z in 1..(N) { if arr[z] < arr[z-1] { print!("!") } }
}

