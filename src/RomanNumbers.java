import java.util.HashMap;
import java.util.Map;

public class RomanNumbers {

	private static Map<Integer, String> dictionary = new HashMap<>();
	static {
		dictionary.put(1, "I");
		dictionary.put(2, "II");
		dictionary.put(3, "III");
		dictionary.put(4, "IV");
	};

	public String convertToRoman(int number) {

		return dictionary.get(number);
	}
}
