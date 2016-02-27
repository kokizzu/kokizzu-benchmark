// bin-ins.cpp
#include<cstdlib>
#include<cstdio>
#include<algorithm>
using namespace std;
void binInsSort(int data[], int size) {
    int index = -1,i = 1,j;
    for(;i<size;i++) {
        int temp = data[i];
        int high = i,low = 0,mid;
        while(low <= high) {
            mid = (low + high) /2;
            if(temp < data[mid]) {
                high = mid - 1;
                index = mid;
            } else if (temp > data[mid]) {
                low = mid + 1;      
            } else if(data[mid] == temp) {
                index = mid;
                break; 
            } 
        }
        for(j = i;j > index;j--) data[j] = data[j-1];
        data[j] = temp;
    }   
}
const int N = 200000;
int main() {
    int* arr = new int[N];
    for(int z=0;z<N;++z) arr[z] = rand();
    binInsSort(arr,N);
    for(int z=1;z<N;++z) if(arr[z]<arr[z-1]) printf("!");
    delete[] arr;
}
