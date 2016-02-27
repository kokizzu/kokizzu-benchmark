using System.Collections.Generic;
class prime {
	public static void Main(string[] args) {
		List<int> res = new List<int>();
		int last = 3;
		res.Add(last);
		while(true) {
			last = last + 2;
			bool prime = true;
			foreach(int v in res) {
				if(v*v>last) break;
				if(last%v == 0) {
					prime = false;
					break;
				}
			}
			if(prime) {
				res.Add(last);
				if(res.Count%100000 == 0) System.Console.WriteLine(last);
				if(last>9999999) break;
			}
		}
	}
}