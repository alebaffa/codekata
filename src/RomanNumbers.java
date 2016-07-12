import java.util.*;

public class RomanNumbers {

	private static List<Number> dictionary = new ArrayList<>();

	static {
		dictionary.add(new Number(50, "L"));
		dictionary.add(new Number(40, "XL"));
		dictionary.add(new Number(10, "X"));
		dictionary.add(new Number(9, "IX"));
		dictionary.add(new Number(5, "V"));
		dictionary.add(new Number(4, "IV"));
		dictionary.add(new Number(1, "I"));
	}


	public String convertToRoman(int arabicInput) {

		StringBuffer result = new StringBuffer();
		dictionary.sort(Comparator.comparingInt(Number::getArabic).reversed());

		for(Number item : dictionary)
			while (arabicInput >= item.getArabic()) {
				result.append(item.getRoman());
				arabicInput -= item.getArabic();
			}

		return result.toString();
	}
}
