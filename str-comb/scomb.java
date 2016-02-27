//////////////////////////////
// scomb.java
public class scomb {
    public static int newGap(int gap){
        gap /= 1.3;
        if(gap == 9 || gap == 10) return 11;
        if(gap < 1) return 1;
        return gap;
    }
    public static void combSort(String a[], int len){
        int gap = len;
        String temp;
        boolean swapped;
        do {
            swapped = false;
            gap = newGap(gap);
            for(int i=0; i < len-gap; ++i) {
                if(a[i].compareTo(a[i+gap])>0) {
                    swapped = true;
                    temp = a[i];
                    a[i] = a[i+gap];
                    a[i+gap] = temp;
                }
            }
        } while(gap > 1 || swapped);
    }
    public static final int N = 10000000;
    public static void main(String[] args) {
        String[] arr = new String[N];
        for(int z=0;z<N;++z) arr[z] = "" + (N-z);    
        combSort(arr,N);
        for(int z=1;z<N;++z) if(arr[z].compareTo(arr[z-1])<0) System.out.print("!");
    }   
}
