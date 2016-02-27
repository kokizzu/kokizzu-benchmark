// hash_map.java
import java.util.*;
class hash_map {
	static class MSD extends HashMap<String,Double> {
		public boolean add_or_inc(String key,double set, double inc) {
			if(!super.containsKey(key)) {
				super.put(key,set);
				return false;
			}
			super.put(key,super.get(key) + inc);
			return true;
		}
	}
	static final int MAX_DATA = 12000000;
	static final char[] i2ch = new char[]{'0','1','2','3','4','5','6','7','8','9','a','B','c','D','e','F'};
	static int get_first_digit(double d) {
		while(d > 10) d /= 10;
		return (int) d;
	}
	static String to_rhex(int v) {
		String hex = "";
		while(v>0) {
			hex += i2ch[v%16];
			v /= 16;
		}
		return hex;
	}
	public static void main(String[] args) {
		MSD m = new MSD();
		int dup1 = 0, dup2 = 0, dup3 = 0;
		for(int z=MAX_DATA;z>0;--z) {
			int val2 = MAX_DATA-z;
			int val3 = MAX_DATA*2-z;
			String key1 = "" + z;
			String key2 = "" + val2;
			String key3 = to_rhex(val3);
			if(m.add_or_inc(key1,z,val2)) ++dup1;
			if(m.add_or_inc(key2,val2,val3)) ++dup2;
			if(m.add_or_inc(key3,val3,z)) ++dup3;
		}
		System.out.println(dup1+" "+dup2+" "+dup3);
		int total = 0, verify = 0, count = 0;
		for (Map.Entry<String,Double> entry : m.entrySet()) {
			total += get_first_digit(entry.getValue());
			verify += entry.getKey().length();
			count += 1;
		}
		System.out.println(total+" "+verify+" "+count);
	}
}