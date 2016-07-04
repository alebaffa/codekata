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

		while(number >= 50){
			result += "L";
			number -= 50;
		}

		while(number >= 10){
			result += "X";
			number -= 10;
		}

		while(number >= 5){
			result += "V";
			number -= 5;
		}

		while(number >= 1){
			result += "I";
			number -= 1;
		}

		return result;
	}
}
