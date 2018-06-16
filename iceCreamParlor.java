//ALDO FUSTER TURPIN

import java.io.*;

import java.math.*;
import java.security.*;
import java.text.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.regex.*;
import java.util.Map;
import java.util.stream.*;;

public class Solution {

    // Complete the icecreamParlor function below.
	//m amount of money
	//arr the cost of each flavour
    static int[] icecreamParlor(int m, int[] arr) {
    	
    	//key->the price of flavour
    	//value->the index in list of prices(in arr)
    	Map<Integer, Integer> complementaryMap = new HashMap<>();
    	int[] result = new int[2];
    	boolean found = false;
    	int currentIndex = 0;
    	
    	while (currentIndex < arr.length && !found) {
    		int element = arr[currentIndex];
    		int the_complementary = m - element;
	    		if (complementaryMap.containsKey(the_complementary)) {
    				int indexOfComplementary = complementaryMap.get(the_complementary);
    				
					/////////////////////////////////////////////////////
					//System.out.println("indexOfComplementary:" + indexOfComplementary);
					//System.out.println("currentIndex:" + currentIndex);
					/////////////////////////////////////////////////////
    				
    				arr[0] = indexOfComplementary + 1;
    				arr[1] = currentIndex + 1;
    				found = true;
	    		}
	    		else {
	    			complementaryMap.put(element, currentIndex);
	    			++currentIndex;
	    		}
    	}
    	/////////////////////////////////////////////////////
    	//System.out.println(Arrays.asList(complementaryMap));
    	System.out.print(arr[0]);
    	System.out.print(" ");
    	System.out.print(arr[1]);
    	System.out.print("\n");
    	/////////////////////////////////////////////////////
    	return result;
    }

    private static final Scanner scanner = new Scanner(System.in);

    public static void main(String[] args) throws IOException {
    	
        int t = scanner.nextInt();
        scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

        for (int tItr = 0; tItr < t; tItr++) {
            int m = scanner.nextInt();
            scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

            int n = scanner.nextInt();
            scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

            int[] arr = new int[n];

            String[] arrItems = scanner.nextLine().split(" ");
            scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

            for (int i = 0; i < n; i++) {
                int arrItem = Integer.parseInt(arrItems[i]);
                arr[i] = arrItem;
            }
     
            icecreamParlor(m, arr);
        }
        scanner.close();
    }
}
