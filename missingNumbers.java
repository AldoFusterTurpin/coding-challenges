// ALDO FUSTER TURPIN
import java.io.*;
import java.util.*;


public class Solution {

    private static final Scanner scanner = new Scanner(System.in);

    public static void main(String[] args) throws IOException {
    	//key->the number
    	//value->the frequency(number of appearances) of the number
    	
    	//Map of arr:
    	Map<Integer, Integer> frequenciesMap = new HashMap<>();
    	
    	//Map of brr:
    	Map<Integer, Integer> frequenciesMap2 = new HashMap<>();
    	
    	SortedSet<Integer> result = new TreeSet<>();
    	
    	//read n
        scanner.nextInt();//nextInt() DOES NOT CONSUME THE NEW LINE 
        
        // read newline
        scanner.nextLine(); 
        
        String input = scanner.nextLine();
        String[] numbers = input.split(" ");

        for(int i = 0; i < numbers.length; ++i) {
        	int frequency = 1;
            int value = Integer.parseInt(numbers[i]);
            if (frequenciesMap.containsKey(value)) frequency = frequenciesMap.get(value) + 1;
            frequenciesMap.put(value, frequency);
    	}

        //read m
        scanner.nextInt();
        
        // read newline
        scanner.nextLine(); 
        
        input = scanner.nextLine();
        numbers = null;
        numbers = input.split(" ");
        
        for (int i = 0; i < numbers.length; i++) {
        	int frequency = 1;
            int value = Integer.parseInt(numbers[i]);
            if (frequenciesMap2.containsKey(value)) frequency = frequenciesMap2.get(value) + 1;
            frequenciesMap2.put(value, frequency);
        }
    	
    	Iterator<Map.Entry<Integer, Integer>> entries = frequenciesMap.entrySet().iterator();
    	while (entries.hasNext()) {
    	    Map.Entry<Integer, Integer> entry = entries.next();
    	    int k = entry.getKey();
    	    int v = entry.getValue();
    	    
    	    if (frequenciesMap2.containsKey(k)) {
    			if (frequenciesMap2.get(k) != v) {
    				result.add(k);
    			}
    			entries.remove();
    			frequenciesMap2.remove(k);
    		}
    		else {
    			result.add(k);
    		}
    	}
    	
    	entries = frequenciesMap2.entrySet().iterator();
    	while (entries.hasNext()) {
    		Map.Entry<Integer, Integer> entry = entries.next();
    		int k = entry.getKey();
    		result.add(k);
    	}
    	
    	Iterator<Integer> it = result.iterator();
    	while(it.hasNext()) {
    		System.out.print(it.next());
    		System.out.print(" ");
    	}
    	
        scanner.close();
    }
}
