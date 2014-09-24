using System.Collections.Generic;

namespace Exercism.csharp.etl
{
	internal class ETL
	{
		public static Dictionary<string, int> Transform(Dictionary<int, IList<string>> oldDataDict)
		{
			var newDataDict = new Dictionary<string, int>();

			foreach (var oldData in oldDataDict)
			{
				foreach (var letter in oldData.Value)
				{
					newDataDict.Add(letter.ToLower(), oldData.Key);
				}
			}

			return newDataDict;
		}
	}
}
