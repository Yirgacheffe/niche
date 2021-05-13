import java.util.Arrays;

public class MergeTwoSortedArrays {

    private static int[] a = new int[]{1, 5, 9, 10, 15, 20};
    private static int[] b = new int[]{2, 3, 8, 30};

    public static void merge(int m, int n) {

        for (int i = n - 1; i >= 0; i--) {

            int j = 0;
            int last = a[m-1];

            for (j = m-2; j >= 0 && a[j] > b[i];j--) {
                a[j+1] = a[j];
            }

            if (j != m -2 || last > b[i]) {
                a[j+ 1] = b[i];
                b[i] = last;
            }
        }

        System.out.println( Arrays.toString(a) );      // ......
        System.out.println( Arrays.toString(b) );      // ......

    }

    public static void main(String[] args) {
        MergeTwoSortedArrays.merge(a.length, b.length);
    }

}   //:~
