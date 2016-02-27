void main() {
	var res = [3];
	var last = 3;
	while(true) {
		last += 2;
		var prime = true;
		for(var z=0;z<res.length;++z) {
			var v = res[z];
			if(v*v > last) break;
			if(last%v == 0) {
				prime = false;
				break;
			}
		}
		if(prime) {
			res.add(last);
			if(res.length%100000 == 0) print(last);
			if(last > 9999999) break;
		}
	}	
}