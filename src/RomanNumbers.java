public class RomanNumbers {

	private static String[] dictionary = {"I", "II", "III"};
	
	public String convertToRoman(int number) {

		return dictionary[number - 1];
	}
}
