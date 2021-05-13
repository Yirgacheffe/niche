import java.util.Map;
import java.util.HashMap;

public class RomaNumeralsToDecimal {

    private static Map<Character, Integer> roman = new HashMap();

    static {
        roman.put('I', 1   );
        roman.put('V', 5   );
        roman.put('X', 10  );
        roman.put('L', 50  );
        roman.put('C', 100 );
        roman.put('D', 500 );
        roman.put('M', 1000);
    }


    public static void toDecimal(String r) {

        int n = r.length();
        int result = 0;

        for (int i = 0; i < n; i++) {

            int cur = roman.get(r.charAt(i));
            if (i != n - 1) {
                int next = roman.get(r.charAt(i + 1));
                if (cur < next ) {
                    result += next - cur;
                    i++;
                } else {
                    result += cur;
                }
            } else {
                result += roman.get(r.charAt(i));
            }

        }

        System.out.println("Result is: " + result);
    }


    public static void main(String[] args) {
        RomaNumeralsToDecimal.toDecimal("MCMIV");
    }

}   //:~