//////////////////////////////
// scomb.c
#include<stdio.h>
#include<stdlib.h>
#include<stdbool.h>
#include<string.h>
int newGap(int gap){
    gap /= 1.3;
    if(gap == 9 || gap == 10) return 11;
    if(gap < 1) return 1;
    return gap;
}
void combSort(char* a[], int len){
    int gap = len;
    bool swapped;
    do {
        swapped = false;
        gap = newGap(gap);
        for(int i=0; i < len-gap; ++i) {
            if(strcmp(a[i],a[i+gap])>0) {
                swapped = true;
                char *tmp = a[i];
                a[i] = a[i+gap];
                a[i+gap] = tmp;
            }
        }
    } while(gap > 1 || swapped);
}
const int N = 10000000;
int main() {
    char** arr = malloc(sizeof(char*)*N);
    for(int z=0;z<N;++z) {
        int len = snprintf(NULL,0,"%d",N-z);
        arr[z] = malloc(len+1);
        snprintf(arr[z], len+1, "%d", N-z);    
    }
    combSort(arr,N);
    for(int z=1;z<N;++z) if(strcmp(arr[z],arr[z-1])<0) printf("!");
    for(int z=0;z<N;++z) free(arr[z]);
    free(arr);
}
