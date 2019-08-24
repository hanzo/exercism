using System;

namespace Exercism.csharp.gigasecond
{
	internal class Gigasecond
	{
		private DateTime _birthday;
		private readonly double SecondsPerGigasecond = Math.Pow(10, 9);

		public Gigasecond(DateTime birthday)
		{
			_birthday = birthday;
		}

		public DateTime Date()
		{
			return _birthday.AddSeconds(SecondsPerGigasecond).Date;
		}
	}
}
