// map.dart
const MAX_DATA = 1230000;
final i2ch = ['0','1','2','3','4','5','6','7','8','9','a','B','c','D','e','F'];
int get_first_digit(d) {
	while(d > 10) d /= 10;
	return d.toInt();
}
string to_rhex(v) {
	var hex = '';
	v = v.toInt();
	while(v>0) {
		hex += i2ch[v%16].toString();
		v = (v/16).toInt();
	}
	return hex;
}
bool add_or_inc(m,key,set,inc) {
	if(m[key] == null) {
		m[key] = set;
		return false;
	}
	m[key] += inc;
	return true;
}
void main() {
	var m = {};
	int dup1 = 0, dup2 = 0, dup3 = 0;
	for(var z=MAX_DATA;z>0;--z) {
		int val2 = MAX_DATA-z;
		int val3 = MAX_DATA*2-z;
		string key1 = z.toString();
		string key2 = val2.toString();
		string key3 = to_rhex(val3);
		if(add_or_inc(m,key1,z,val2)) ++dup1;
		if(add_or_inc(m,key2,val2,val3)) ++dup2;
		if(add_or_inc(m,key3,val3,z)) ++dup3;
	}
	print('$dup1 $dup2 $dup3');
	int total = 0, verify = 0, count = 0;
	void iterate(key,val) {
		total += get_first_digit(val);
		verify += key.length;
		count += 1;	
	}
	m.forEach(iterate);
	print('$total $verify $count');
}