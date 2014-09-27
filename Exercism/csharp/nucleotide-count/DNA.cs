using System;
using System.Collections.Generic;
using System.Linq;

namespace Exercism.csharp.nucleotide_count
{
	internal class DNA
	{
		private readonly Dictionary<char, int> _nucleotideCounts;

		public DNA(string sequence)
		{
			_nucleotideCounts = new Dictionary<char, int>
			{
				{'A', 0},
				{'T', 0},
				{'C', 0},
				{'G', 0},
			};

			CountNucleotides(sequence);
		}

		public int Count(char nucleotide)
		{
			if (!_nucleotideCounts.ContainsKey(nucleotide))
				throw new InvalidNucleotideException();

			return _nucleotideCounts[nucleotide];
		}

		public Dictionary<char, int> NucleotideCounts
		{
			get { return new Dictionary<char, int>(_nucleotideCounts); }
		}

		private void CountNucleotides(string sequence)
		{
			foreach (char currentChar in sequence.Where(c => _nucleotideCounts.ContainsKey(c)))
				_nucleotideCounts[currentChar] ++;
		}
	}

	internal class InvalidNucleotideException : ApplicationException
	{
	}
}
