import org.junit.*;

import static org.hamcrest.CoreMatchers.is;
import static org.hamcrest.MatcherAssert.assertThat;
import static org.hamcrest.core.IsEqual.equalTo;

public class RomanNumbersTest {

	@Test
	public void convertArabicToRomanNumbers(){
		RomanNumbers converter = new RomanNumbers();
		assertThat(converter.convertToRoman(1), is(equalTo("I")));
		assertThat(converter.convertToRoman(2), is(equalTo("II")));
		assertThat(converter.convertToRoman(3), is(equalTo("III")));
		assertThat(converter.convertToRoman(4), is(equalTo("IV")));
		assertThat(converter.convertToRoman(5), is(equalTo("V")));
		assertThat(converter.convertToRoman(6), is(equalTo("VI")));
		assertThat(converter.convertToRoman(7), is(equalTo("VII")));
		assertThat(converter.convertToRoman(9), is(equalTo("IX")));
		assertThat(converter.convertToRoman(10), is(equalTo("X")));
		assertThat(converter.convertToRoman(11), is(equalTo("XI")));
		assertThat(converter.convertToRoman(40), is(equalTo("XL")));
		assertThat(converter.convertToRoman(50), is(equalTo("L")));
		assertThat(converter.convertToRoman(55), is(equalTo("LV")));
		assertThat(converter.convertToRoman(60), is(equalTo("LX")));
	}


}