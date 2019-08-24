using System.Linq;

namespace Exercism.csharp.anagram
{
	internal class Anagram
	{
		private readonly string _baseString;

		public Anagram(string baseString)
		{
			_baseString = baseString.ToLowerInvariant();
		}

		public string[] Match(string[] matchStringArray)
		{
			return matchStringArray.Where(IsAnagram).OrderBy(a => a).ToArray();
		}

		private bool IsAnagram(string matchString)
		{
			matchString = matchString.ToLowerInvariant();

			return !matchString.Equals(_baseString) && matchString.OrderBy(m => m).SequenceEqual(_baseString.OrderBy(b => b));
		}

		/* 
		 * My intial, more verbose implementation
		 *
		private bool IsAnagram(string matchString)
		{
			matchString = matchString.ToLowerInvariant();

			if (matchString.Equals(_baseString))
				return false;

			List<char> matchStringChars = matchString.ToCharArray().ToList();
			List<char> baseStringChars = _baseString.ToCharArray().ToList();

			// Attempt to remove each character in the match string from the base string.
			// If any characters are not present, Remove returns false and we know the string is not an anagram.
			if (matchStringChars.Any(character => !baseStringChars.Remove(character)))
				return false;

			// If there are any characters remaining in the base string, this is not an anagram.
			return baseStringChars.Count == 0;
		}
		 */
	}
}
