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
	}


}