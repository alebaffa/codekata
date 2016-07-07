/**
 * Created by alebaffa on 7/7/16.
 */
public class Number implements Comparable<Number> {

	private int arabic;
	private String roman;

	public Number(int arabicIn, String romanIn) {
		arabic = arabicIn;
		roman = romanIn;
	}

	public int getArabic() {
		return this.arabic;
	}

	public String getRoman() {
		return this.roman;
	}


	@Override
	public int compareTo(Number number) {

		if(this.getArabic() > number.getArabic())
			return 1;

		else if(this.getArabic() < number.getArabic())
			return -1;

		else return 0;
	}
}
