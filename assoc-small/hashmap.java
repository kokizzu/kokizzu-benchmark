// hashmap.java
import java.util.HashMap;
import java.util.Map;
public class hashmap {
	public static final int MAX_DATA = 1230000;
	public static final char[] i2ch = {'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'B', 'c', 'D', 'e', 'F'};
	public static int get_first_digit(double d) {
		while(d > 10) {
			d /= 10;
		}
		return (int)(d);
	}
	public static String to_rhex(int v) {
		StringBuilder hex = new StringBuilder();
		while(v > 0) {
			hex.append(i2ch[v%16]);
			v /= 16;
		}
		return hex.toString();
	}
	public static int set_or_inc(HashMap<String,Double> m, String key,int set,int inc,int ctr) {
		if(!m.containsKey(key)) {
			m.put(key, (double)(set));
		} else {
			m.put(key, m.get(key) + inc);
			ctr += 1;
		}
		return ctr;
	}
	public static void main(String[] args) {
		HashMap<String,Double> m = new HashMap<>();
		int dup1 = 0, dup2 = 0, dup3 = 0;
		for(int z = MAX_DATA; z > 0; z--) {
			int val2 = MAX_DATA - z;
			int val3 = MAX_DATA*2 - z;
			String key1 = String.valueOf(z);
			String key2 = String.valueOf(val2);
			String key3 = to_rhex(val3);
			dup1 = set_or_inc(m, key1, z, val2, dup1);
			dup2 = set_or_inc(m, key2, val2, val3, dup2);
			dup3 = set_or_inc(m, key3, val3, z, dup3);
		}
		System.out.format("%d %d %d\n",dup1, dup2, dup3);
		int total = 0, verify = 0, count = 0;
		for(Map.Entry e : m.entrySet()) {
			total += get_first_digit((double)e.getValue());
			verify += e.getKey().toString().length();
			count += 1;
		}
		System.out.format("%d %d %d\n",total, verify, count);
	}
}