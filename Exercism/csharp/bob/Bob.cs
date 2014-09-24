﻿using System;
using System.Linq;

namespace Exercises.csharp.bob
{
	internal class Bob
	{
		public string Hey(string input)
		{
			if (String.IsNullOrWhiteSpace(input))
				return "Fine. Be that way!";

			// The input is 'yelling' if it contains letters and every letter is capitalized
			if (input.Any(char.IsLetter) && input.All(c => !char.IsLetter(c) || char.IsUpper(c)))
				return "Whoa, chill out!";

			if (input.EndsWith("?"))
				return "Sure.";

			return "Whatever.";
		}
	}
}
