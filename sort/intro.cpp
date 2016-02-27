// intro.cpp
#include<cstdlib>
#include<cstdio>
#include<algorithm>
using namespace std;
const int N = 10000000;
int main() {
    int* arr = new int[N];
    for(int z=0;z<N;++z) arr[z] = rand();
    sort(arr,arr+N);
    for(int z=1;z<N;++z) if(arr[z]<arr[z-1]) printf("!");
    delete[] arr;
}
