// judy.cpp
#include "deps/judySArray.h"
#include <string>
#include <iostream>
#include <cstdlib>
#include <cstring>
#include <vector>
using namespace std;
typedef judySArray<double> MSD;
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
	const char* cstr = key.c_str();
	double it = m.find(cstr);
	if(!it) {
		m.insert(cstr,set);
		return;
	}
	m.insert(cstr,it+inc);
	++ctr;
}
#include "deps/cedar.h"
class MSD2  {
public:
	vector<double> data;
	typedef cedar::da<int> CI;
	CI da;
	bool exists(const string& key,double &old) {
		int idx = -1;
		bool found = da.exactMatchExists(key.c_str(),key.size(),&idx);
		if(found) old = data[idx];
		return found;
	}
	void insert(const string& key,double val) {
		da.update(key.c_str(),key.size(),data.size());
		data.push_back(val);
	}
	void update(const string& key,double val) {
		int idx = -1;
		bool found = da.exactMatchExists(key.c_str(),key.size(),&idx);
		if(found) {
			data[idx] = val;
			return;
		}
		insert(key,val);
	}
};
void add_or_inc(MSD2 &m, const string& key,double set, double inc, int& ctr) {
	double old;
	if(!m.exists(key,old)) {
		m.insert(key,set);
		return;
	}
	m.update(key,old+inc);
	++ctr;
}
int main() {
	MSD m(64);	
	MSD2 m2;
	int dup1 = 0, dup2 = 0, dup3 = 0;
	int vup1 = 0, vup2 = 0, vup3 = 0;
	for(int z=MAX_DATA;z>0;--z) {
		int val2 = MAX_DATA-z;
		int val3 = MAX_DATA*2-z;
		string key1 = to_string(z);
		string key2 = to_string(val2);
		string key3 = to_rhex(val3);
		add_or_inc(m,key1,z,val2,dup1);
		add_or_inc(m,key2,val2,val3,dup2);
		add_or_inc(m,key3,val3,z,dup3);
		add_or_inc(m2,key1,z,val2,vup1);
		add_or_inc(m2,key2,val2,val3,vup2);
		add_or_inc(m2,key3,val3,z,vup3);
	}
	cout << dup1 << ' ' << dup2 << ' ' << dup3 << endl;
	cout << vup1 << ' ' << vup2 << ' ' << vup3 << endl;
	int xotal = 0, xerify = 0, xount = 0;
	union { int i; int x; } b;
	size_t from = 0, p = 0;
	char key[256] = {0};
	for (b.i = m2.da.begin(from, p); b.i != MSD2::CI::CEDAR_NO_PATH; b.i = m2.da.next(from, p)) {
		double it2 = m2.data[b.x];
		xotal += get_first_digit(it2); 
		m2.da.suffix(key,p,from);
		xerify += strlen(key);
		xount += 1;
		double it = m.find(key);
		if(it != it2) {
			cout << "mismatch value for " << key << " : " << it2 << " vs " << it << endl;
		}
	}
	int total = 0, verify = 0, count = 0;
	for(auto &it = m.begin();m.success(); m.next()) {
		total += get_first_digit(it.value);	
		verify += strlen((const char *) it.key);
		count += 1;
	}
	cout << total << ' ' << verify << ' ' << count << endl;
	cout << xotal << ' ' << xerify << ' ' << xount << endl;
}