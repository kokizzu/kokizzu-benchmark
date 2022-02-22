//////////////////////////////
// comb.cs
class Comb {
    static int newGap(int gap){
        gap = (int)(gap / 1.3);
        if(gap == 9 || gap == 10) return 11;
        if(gap < 1) return 1;
        return gap;
    }
    static void combSort(double[] a, int len){
        int gap = len;
        double temp;
        bool swapped;
        do {
            swapped = false;
            gap = newGap(gap);
            for(int i=0; i < len-gap; ++i) {
                if(a[i] > a[i+gap]) {
                    swapped = true;
                    temp = a[i];
                    a[i] = a[i+gap];
                    a[i+gap] = temp;
                }
            }
        } while(gap > 1 || swapped);
    }
    static int N = 10000000;
    static void Main() {
        double[] arr = new double[N];
        for(int z=0;z<N;++z) arr[z] = N-z;    
        combSort(arr,N);
        for(int z=1;z<N;++z) if(arr[z]<arr[z-1]) System.Console.Write("!");
    }
}
