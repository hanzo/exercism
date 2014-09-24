using System;

namespace Exercism.csharp.hamming
{
	internal class Hamming
	{
		public static int Compute(string strand1, string strand2)
		{
			var diffCount = 0;

			var minStrandLength = Math.Min(strand1.Length, strand2.Length);

			for (var x = 0; x < minStrandLength; x++)
			{
				if (strand1[x] != strand2[x])
					diffCount++;
			}

			return diffCount;
		}
	}
}
