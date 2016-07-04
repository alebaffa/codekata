import java.util.*;

public class RomanNumbers {

	private static Map<Integer, String> dictionary = new LinkedHashMap<>();

	static {
		dictionary.put(50, "L");
		dictionary.put(40, "XL");
		dictionary.put(10, "X");
		dictionary.put(9, "IX");
		dictionary.put(5, "V");
		dictionary.put(4, "IV");
		dictionary.put(1, "I");
	};

	public String convertToRoman(int number) {

		String result = "";

		Iterator iterator = dictionary.entrySet().iterator();

		while (iterator.hasNext()) {
			Map.Entry pair = (Map.Entry) iterator.next();
			Integer arabicNumber = (Integer) pair.getKey();
			String romanNumber = (String) pair.getValue();

			while (number >= arabicNumber) {
				result += romanNumber;
				number -= arabicNumber;
			}
		}

		return result;
	}
}
