using System;
using System.Collections.Generic;
using System.Linq;

namespace Exercism.csharp.scrabble_score
{
	internal class Scrabble
	{
		private readonly string _inputWord;

		private static readonly List<Tuple<int, char[]>> CharacterGroups = new List<Tuple<int, char[]>>
		{
			Tuple.Create(1, new[] {'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T'}),
			Tuple.Create(2, new[] {'D', 'G'}),
			Tuple.Create(3, new[] {'B', 'C', 'M', 'P'}),
			Tuple.Create(4, new[] {'F', 'H', 'V', 'W', 'Y'}),
			Tuple.Create(5, new[] {'K'}),
			Tuple.Create(8, new[] {'J', 'X'}),
			Tuple.Create(10, new[] {'Q', 'Z'})
		};

		public Scrabble(string inputWord)
		{
			_inputWord = inputWord;
		}

		public int Score()
		{
			return Score(_inputWord);
		}

		public static int Score(string inputWord)
		{
			if (String.IsNullOrWhiteSpace(inputWord))
				return 0;

			inputWord = inputWord.ToUpper();

			// For each group of characters and its associated score, count the occurrences of those characters in the given word and
			// multiply this count by the letter score. Sum these scores together and return the result.
			return CharacterGroups.Sum(characterGroup => characterGroup.Item1 * inputWord.Count(c => characterGroup.Item2.Contains(c)));
		}
	}
}
