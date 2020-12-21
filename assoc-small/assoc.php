<?php // assoc.php
const MAX_DATA = 1230000;
$i2ch = array('0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'B', 'c', 'D', 'e', 'F');
function get_first_digit($d) {
	while($d > 10) $d /= 10;
	return intval($d);
}
function to_rhex($v) {
	global $i2ch;
	$hex = '';
	$v = intval($v);
	while($v > 0) {
		$hex .= $i2ch[$v%16];
		$v = intval($v / 16);
	}
	return $hex;
}
function add_or_inc(&$m, $key, $set, $inc, &$ctr) {
	if(!array_key_exists($key,$m)) {
		$m[$key] = $set;
		return;
	}
	$m[$key] += $inc;
	++$ctr;		
}
function main() {
	$m = array();
	$dup1 = $dup2 = $dup3 = 0;
	for ($z = MAX_DATA; $z > 0; --$z) {
		$val2 = MAX_DATA - $z;
		$val3 = MAX_DATA*2 - $z;
		$key1 = '' . $z;
		$key2 = '' . $val2;
		$key3 = to_rhex($val3);
		add_or_inc($m,$key1,$z,$val2,$dup1);
		add_or_inc($m,$key2,$val2,$val3,$dup2);
		add_or_inc($m,$key3,$val3,$z,$dup3);
	}
	echo $dup1 . ' ' . $dup2 . ' ' . $dup3 . "\n";
	$total = $verify = $count = 0;
	foreach($m as $k => $v) {
		$total += get_first_digit($v);
		$verify += strlen($k);
		$count += 1;
	}
	echo $total . ' ' . $verify . ' ' . $count . "\n";
}
main();
