//Aldo Fuster Turpin

import java.io.*;
import java.math.*;
import java.security.*;
import java.text.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.function.*;
import java.util.regex.*;
import java.util.stream.*;
import static java.util.stream.Collectors.joining;
import static java.util.stream.Collectors.toList;

class Result {

    /*
     * Complete the 'countingValleys' function below.
     *
     * The function is expected to return an INTEGER.
     * The function accepts following parameters:
     *  1. INTEGER steps
     *  2. STRING path
     */
    public static int countingValleys(int steps, String path) {
        int nUp, nDown, nValleys;
        nUp = nDown = nValleys = 0;
        
        char previous, actual;
        
        for (int i = 1; i < steps; i += 2) {
            previous = path.charAt(i-1);
            actual = path.charAt(i);
            
            if (previous == 'U') {
                ++nUp;
            } else {
                ++nDown;
            }
            
            if (actual == 'U') {
                ++nUp;
            } else {
                ++nDown;
            }
            
            if (nUp == nDown && nUp != 0) {
                if (actual == 'U') {
                    nValleys++;
                }
                nUp = nDown = 0;
            }
        }   
        return nValleys;
    }

}

public class Solution {
    public static void main(String[] args) throws IOException {
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        int steps = Integer.parseInt(bufferedReader.readLine().trim());

        String path = bufferedReader.readLine();

        int result = Result.countingValleys(steps, path);

        bufferedWriter.write(String.valueOf(result));
        bufferedWriter.newLine();

        bufferedReader.close();
        bufferedWriter.close();
    }
}
