import java.util.HashMap;
import java.util.Map;

public class RomanNumbers {

	private static Map<Integer, String> dictionary = new HashMap<>();
	static {
		dictionary.put(1, "I");
		dictionary.put(4, "IV");
	};

	public String convertToRoman(int number) {

		if(dictionary.containsKey(number))
			return dictionary.get(number);

		return dictionary.get(1) + convertToRoman(number - 1);
	}
}
