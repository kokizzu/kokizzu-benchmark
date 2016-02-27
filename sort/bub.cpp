// bub.cpp
#include<cstdlib>
#include<cstdio>
#include<algorithm>
using namespace std;
void bubSort(int a[], size_t length) {
    for(size_t z=length-1;z>0;--z) {
        for(size_t y=0;y<z;++y) if(a[y]>a[y+1]) swap(a[y],a[y+1]);
    }    
}
const int N = 200000;
int main() {
    int* arr = new int[N];
    for(int z=0;z<N;++z) arr[z] = rand();
    bubSort(arr,N);
    for(int z=1;z<N;++z) if(arr[z]<arr[z-1]) printf("!");
    delete[] arr;
}
