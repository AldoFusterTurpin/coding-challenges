//ALDO FUSTER TURPIN

import java.io.*;
import java.math.*;
import java.text.*;
import java.util.*;
import java.util.regex.*;
import java.lang.Math;

public class Solution {
    
    private static int jumpsToPStartingFromBeginning(int n, int p) {
        if (p%2 == 0) {
            return p / 2;
        }
        return (p - 1) / 2;
    }
    
    private static int jumpsToPStartingFromEnd(int n, int p) {
        // same page
        if (p == n || (p%2 == 0 && n == p+1)) {
            return 0;
        }

        if (n%2 == 0 && p%2 != 0) {
            return (n + 1 - p) / 2;
        }
        return (n - p) / 2;
    }
    
    /*
     * Complete the pageCount function below.
     */
    private static int pageCount(int n, int p) {
        return Math.min(jumpsToPStartingFromBeginning(n,p), 
                        jumpsToPStartingFromEnd(n,p));
    }

    private static final Scanner scanner = new Scanner(System.in);

    public static void main(String[] args) throws IOException {
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        int n = scanner.nextInt();
        scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])*");

        int p = scanner.nextInt();
        scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])*");

        int result = pageCount(n, p);

        bufferedWriter.write(String.valueOf(result));
        bufferedWriter.newLine();

        bufferedWriter.close();

        scanner.close();
    }
}