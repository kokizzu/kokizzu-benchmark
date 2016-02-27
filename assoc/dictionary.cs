// dictionary.cs
using System;
using System.Collections.Generic;
class dictionary {
	class MSD: Dictionary<string,double> {
		public bool add_or_inc(string key, double set, double inc) {
			if(!base.ContainsKey(key)) {
				base.Add(key,set);
				return false;
			}
			base[key] += inc;
			return true;
		}
	}
	static int MAX_DATA = 12000000;
	static char[] i2ch = new char[]{'0','1','2','3','4','5','6','7','8','9','a','B','c','D','e','F'};
	static int get_first_digit(double d) {
		while(d > 10) d /= 10;
		return (int) d;
	}
	static string to_rhex(int v) {
		string hex = "";
		while(v>0) {
			hex += i2ch[v%16];
			v /= 16;
		}
		return hex;
	}
	public static void Main(string[] args) {
		MSD m = new MSD();
		int dup1 = 0, dup2 = 0, dup3 = 0;
		for(int z=MAX_DATA;z>0;--z) {
			int val2 = MAX_DATA-z;
			int val3 = MAX_DATA*2-z;
			string key1 = "" + z;
			string key2 = "" + val2;
			string key3 = to_rhex(val3);
			if(m.add_or_inc(key1,z,val2)) ++dup1;
			if(m.add_or_inc(key2,val2,val3)) ++dup2;
			if(m.add_or_inc(key3,val3,z)) ++dup3;
		}
		Console.WriteLine(dup1 + " " + dup2 + " " + dup3);
		int total = 0;
		int verify = 0;
		int count = 0;
		foreach(KeyValuePair<string,double> entry in m) {
			total += get_first_digit(entry.Value);
			verify += entry.Key.Length;
			count += 1;
		}
		Console.WriteLine(total + " " + verify + " " + count);
	}
}