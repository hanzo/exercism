using System.Collections.Generic;
using System.Linq;

namespace Exercism.csharp.word_count
{
	internal class Phrase
	{
		private readonly string _phrase;

		public Phrase(string phrase)
		{
			_phrase = phrase.ToLowerInvariant();
		}

		public Dictionary<string, int> WordCount()
		{
			var wordCounts = new Dictionary<string, int>();

			string[] words = _phrase.Split(new[] {',', ' '});

			foreach (string word in words)
			{
				// Skip chunks that are only punctuation or symbols
				if (word.All(c => char.IsPunctuation(c) || char.IsSymbol(c)))
					continue;

				// Find the first and last alphanumeric characters in the string and trim off any preceeding or trailing characters
				int firstLetterPos = word.IndexOf(word.First(c => char.IsLetter(c) || char.IsNumber(c)));
				int lastLetterPos = word.LastIndexOf(word.Last(c => char.IsLetter(c) || char.IsNumber(c)));
				string trimmedWord = word.Substring(firstLetterPos, lastLetterPos - firstLetterPos + 1);

				if (wordCounts.ContainsKey(trimmedWord))
					wordCounts[trimmedWord]++;
				else
					wordCounts[trimmedWord] = 1;
			}

			return wordCounts;
		}
	}
}
