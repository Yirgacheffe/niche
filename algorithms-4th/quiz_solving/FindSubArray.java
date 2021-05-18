import java.util.HashMap;
import java.util.Map;

public class FindSubArray {

    public static void findByHash(int[] a, int sum) {

        Map<Integer, Integer> k = new HashMap();
        int cSum =  0;
        
        for (int i = 0; i < a.length; i++) {
            cSum += a[i];
            int e = cSum - sum;

            if (k.get(e) != null) {
                int start = k.get(e) + 1;
                System.out.println("found: " + start + " to " + i);
                return;
            } else {
                k.put(cSum, i);
            }
        }

        System.out.println("not found!");
        return;

    }


    public static void main(String[] args) {
        findByHash(new int[]{1, 4, 20, 3, 10, 5}, 33);
    }

}