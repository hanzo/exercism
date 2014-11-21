using System.Collections.Generic;
using System.Linq;

namespace Exercism.csharp.sum_of_multiples
{
	class SumOfMultiples
	{
		private readonly List<int> _baseNumbers;

		public SumOfMultiples()
		{
			_baseNumbers = new List<int> { 3, 5 };
		}

		public SumOfMultiples(List<int> baseNumbers)
		{
			_baseNumbers = baseNumbers;
		}

		public int To(int toNumber)
		{
			return Enumerable.Range(0, toNumber).Where(x => _baseNumbers.Any(b => x % b == 0)).Sum();
		}
	}
}
