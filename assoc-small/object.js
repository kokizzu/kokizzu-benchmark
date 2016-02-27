// object.js
if(typeof print == 'undefined') print = function(){};
if(typeof console == 'undefined') console = {log:print};
var MAX_DATA = 120000;
var i2ch = ['0','1','2','3','4','5','6','7','8','9','a','B','c','D','e','F'];
function get_first_digit(d) {
	while(d > 10) d /= 10;
	return d|0;
}
function to_rhex(v) {
	var hex = '';
	v = v|0;
	while(v>0) {
		hex += i2ch[v%16];
		v = (v/16)|0;
	}
	return hex;
}
function add_or_inc(m,key,set,inc) {
	if(m[key] === undefined) {
		m[key] = set;
		return false;
	}
	m[key] += inc;
	return true;
}
(function() {
	var m = {};
	var dup1 = 0, dup2 = 0, dup3 = 0;
	for(var z=MAX_DATA;z>0;--z) {
		var val2 = MAX_DATA-z;
		var val3 = MAX_DATA*2-z;
		var key1 = "" + z;
		var key2 = "" + val2;
		var key3 = to_rhex(val3);
		if(add_or_inc(m,key1,z,val2)) ++dup1;
		if(add_or_inc(m,key2,val2,val3)) ++dup2;
		if(add_or_inc(m,key3,val3,z)) ++dup3;
	}
	console.log(dup1,dup2,dup3);
	var total = 0, verify = 0, count = 0;
	for (var key in m) {
		total += get_first_digit(m[key]);
		verify += key.length;
		count += 1;
	}
	console.log(total,verify,count);
})()