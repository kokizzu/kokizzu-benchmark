// merge.cpp
#include<vector>
#include<string>
#include<algorithm>
#include<cstdio>
using namespace std;
void merge_sort( int* beg, int* end) {
    if (end - beg > 1)  {
        int* mid = beg + (end - beg) / 2;
        merge_sort(beg, mid);
        merge_sort(mid, end);
        inplace_merge(beg, mid, end);
    }
}
const int N = 10000000;
int main() {
    int* arr = new int[N];
    for(int z=0;z<N;++z) arr[z] = rand();
    merge_sort(arr,arr+N);
    for(int z=1;z<N;++z) if(arr[z]<arr[z-1]) printf("!");
    delete[] arr;
}

