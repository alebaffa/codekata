import java.util.HashMap;
import java.util.Map;

public class RomanNumbers {

	private static Map<Integer, String> dictionary = new HashMap<>();
	static {
		dictionary.put(1, "I");
		dictionary.put(4, "IV");
		dictionary.put(5, "V");
		dictionary.put(6, "VI");
		dictionary.put(7, "VII");
	};

	public String convertToRoman(int number) {

		if(dictionary.containsKey(number))
			return dictionary.get(number);

		return dictionary.get(1) + convertToRoman(number - 1);
	}
}
