fn main() {
  let mut last : i32 = 3;
  let mut res = vec![last];
  loop {
    last += 2;
    let mut prime = true;
    for v in res.iter() {
      if *v * *v > last {
        break;
      }
      if (last % *v) == 0  {
        prime = false;
        break;
      }  
    }
    if prime {
    res.push(last);
      if res.len() % 100000 == 0 {
        println!("{}",last);
      }
      if last > 9999999 {
        break;
      }
    } 
  }  
}