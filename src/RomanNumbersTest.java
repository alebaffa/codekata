import org.junit.*;

import static org.hamcrest.CoreMatchers.is;
import static org.hamcrest.MatcherAssert.assertThat;
import static org.hamcrest.core.IsEqual.equalTo;

public class RomanNumbersTest {

	@Test
	public void convertArabicToRomanNumbers(){
		RomanNumbers converter = new RomanNumbers();
		assertThat(converter.convertToRoman(1), is(equalTo("I")));
	}


}