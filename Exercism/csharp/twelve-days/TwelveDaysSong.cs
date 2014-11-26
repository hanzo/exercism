using System.Text;

namespace Exercism.csharp.twelve_days
{
	internal class TwelveDaysSong
	{
		private readonly string[] _daysOfChristmas =
		{
			"zeroeth",  // placeholder for index 0
			"first",
			"second",
			"third",
			"fourth",
			"fifth",
			"sixth",
			"seventh",
			"eighth",
			"ninth",
			"tenth",
			"eleventh",
			"twelfth"
		};

		private readonly string[] _gifts =
		{
			"an off-by-one error",  // placeholder for index 0
			"a Partridge in a Pear Tree.",
			"two Turtle Doves, and ",
			"three French Hens, ",
			"four Calling Birds, ",
			"five Gold Rings, ",
			"six Geese-a-Laying, ",
			"seven Swans-a-Swimming, ",
			"eight Maids-a-Milking, ",
			"nine Ladies Dancing, ",
			"ten Lords-a-Leaping, ",
			"eleven Pipers Piping, ",
			"twelve Drummers Drumming, "
		};

		private string GetGiftList(int verseNum)
		{
			var sb = new StringBuilder();

			for (int giftNum = verseNum; giftNum > 0; giftNum--)
				sb.Append(_gifts[giftNum]);

			return sb.ToString();
		}

		public string Verse(int verseNumber)
		{
			if (verseNumber < 1 || verseNumber > 12)
				return "Verse number must be between 1 and 12";

			return new StringBuilder("On the ")
				.Append(_daysOfChristmas[verseNumber])
				.Append(" day of Christmas my true love gave to me, ")
				.Append(GetGiftList(verseNumber))
				.Append("\n")
				.ToString();
		}

		public string Verses(int startVerse, int endVerse)
		{
			var sb = new StringBuilder();

			for (int verse = startVerse; verse <= endVerse; verse++)
				sb.Append(Verse(verse)).Append("\n");

			return sb.ToString();
		}

		public string Sing()
		{
			return Verses(1, 12);
		}
	}
}
