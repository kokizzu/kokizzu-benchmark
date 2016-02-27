// map.cpp
#include <map>
#include <string>
#include <iostream>
#include <cstdlib>
using namespace std;
typedef map<string,double> MSD;
const int MAX_DATA = 12000000;
const char i2ch[] = {'0','1','2','3','4','5','6','7','8','9','a','B','c','D','e','F'};
int get_first_digit(double d) {
	while(d > 10) d /= 10;
	return d;
}
string to_rhex(int v) {
	char hex[32];	
	int start = 0;
	while(v>0) {
		hex[start] = i2ch[v%16];
		v /= 16;
		++start;
	}
	hex[start] = 0;
	return hex;
}
void add_or_inc(MSD &m, const string& key,double set, double inc, int& ctr) {
	MSD::iterator it = m.find(key);
	if(it == m.end()) {
		m[key] = set;
		return;
	}
	m[key] += inc;
	++ctr;
}
int main() {
	MSD m;	
	int dup1 = 0, dup2 = 0, dup3 = 0;
	for(int z=MAX_DATA;z>0;--z) {
		int val2 = MAX_DATA-z;
		int val3 = MAX_DATA*2-z;
		string key1 = to_string(z);
		string key2 = to_string(val2);
		string key3 = to_rhex(val3);
		add_or_inc(m,key1,z,val2,dup1);
		add_or_inc(m,key2,val2,val3,dup2);
		add_or_inc(m,key3,val3,z,dup3);
	}
	cout << dup1 << ' ' << dup2 << ' ' << dup3 << endl;
	int total = 0, verify = 0, count = 0;
	for(auto const &it : m) {
		total += get_first_digit(it.second);
		verify += it.first.size();
		count += 1;
	}	
	cout << total << ' ' << verify << ' ' << count << endl;
}