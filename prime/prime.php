<? 
function main() {
	$res = array(3);
	$last = 3;
	while(true) {
		$last += 2;
		$prime = true;
		for($z=0;$z<sizeof($res);++$z) {
			$v = $res[$z];
			if($v*$v > $last) break;
			if($last%$v == 0) {
				$prime = false;
				break;
			}
		}
		if($prime) {
			$res[] = $last;
			if(sizeof($res)%100000 == 0) echo $last."\n";
			if($last > 9999999) break;
		}
	}
}
main();