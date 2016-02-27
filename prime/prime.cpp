#include <iostream>
#include <vector>
using namespace std;
int main() {
	vector<int> res;
	res.push_back(3);
	int last = 3;
	while(true) {
		last += 2;
		bool prime = true;
		for(int v : res) {
			if(v*v > last) break;
			if(last % v == 0) {
				prime = false;
				break;
			}
		}
		if(prime) {
			res.push_back(last);
			if(res.size() % 100000 == 0) cout << last << endl;
			if(last>9999999) break;
		}
	}
}