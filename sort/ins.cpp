// ins.cpp
#include<cstdlib>
#include<cstdio>
#include<algorithm>
using namespace std;
void insSort(int a[], size_t length) {
    for(size_t z=1;z<length;++z) {
     int copy = a[z];     
     size_t y = z;
        for(;y>0 && a[y-1]>copy;--y) a[y] = a[y-1];
        a[y] = copy;
    }    
}
const int N = 200000;
int main() {
    int* arr = new int[N];
    for(int z=0;z<N;++z) arr[z] = rand();
    insSort(arr,N);
    for(int z=1;z<N;++z) if(arr[z]<arr[z-1]) printf("!");
    delete[] arr;
}
