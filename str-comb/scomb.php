<? //////////////////////////////
   // scomb.php
function newGap($gap){
    $gap /= 1.3;
    $gap = intval($gap);
    if($gap == 9 || $gap == 10) return 11;
    if($gap < 1) return 1;
    return $gap;
}
function combSort(&$a) {
    $len = sizeof($a);
    $gap = $len;
    $temp = 0.0;
    $swapped = false;
    do {
        $swapped = false;
        $gap = newGap($gap);
        for($i=0; $i < $len-$gap; ++$i) {
            if($a[$i] > $a[$i+$gap]) {
                $swapped = true;
                $temp = $a[$i];
                $a[$i] = $a[$i+$gap];
                $a[$i+$gap] = $temp;
            }
        }
    } while($gap > 1 || $swapped);
}
const N = 10000000;
$arr = array_fill(0,N,'');
for($z=0;$z<N;++$z) $arr[$z] = ''.(N-$z);
combSort($arr);
for($z=1;$z<N;++$z) if($arr[$z]<$arr[$z-1]) echo "!";
