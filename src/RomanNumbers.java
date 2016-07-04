import java.util.HashMap;
import java.util.Map;

public class RomanNumbers {

	private static Map<Integer, String> dictionary = new HashMap<>();
	static {
		dictionary.put(1, "I");
		dictionary.put(4, "IV");
		dictionary.put(5, "V");
	};

	public String convertToRoman(int number) {

		if(dictionary.containsKey(number))
			return dictionary.get(number);

		String result = "";

		if(number > 5){
			result += "V";
			return result + convertToRoman(number - 5);
		}

		return dictionary.get(1) + convertToRoman(number - 1);
	}
}
