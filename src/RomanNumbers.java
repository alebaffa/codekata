import java.util.HashMap;
import java.util.Map;

public class RomanNumbers {

	private static Map<Integer, String> dictionary = new HashMap<>();
	static {
		dictionary.put(1, "I");
		dictionary.put(4, "IV");
		dictionary.put(5, "V");
		dictionary.put(9, "IX");
		dictionary.put(10, "X");
		dictionary.put(40, "XL");
		dictionary.put(50, "L");
	};

	public String convertToRoman(int number) {

		if(dictionary.containsKey(number))
			return dictionary.get(number);

		String result = "";

		if(number > 50){
			result += "L";
			return result + convertToRoman(number - 50);
		}

		if(number > 10){
			result += "X";
			return result + convertToRoman(number - 10);
		}

		if(number > 5){
			result += "V";
			return result + convertToRoman(number - 5);
		}

		return dictionary.get(1) + convertToRoman(number - 1);
	}
}
