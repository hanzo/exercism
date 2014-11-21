using System;

namespace Exercism.csharp.robot_name
{
	class Robot
	{
		public string Name { get; private set; }

		private static readonly Random _random = new Random();

		public Robot()
		{
			Name = GetRandomName();
		}

		public void Reset()
		{
			Name = GetRandomName();
		}

		private string GetRandomName()
		{
			string name = string.Empty;

			for (int x = 0; x <= 4; x++)
			{
				if (x <= 1)
					name += (char)('A' + _random.Next(26));
				else
					name += _random.Next(10);
			}

			return name;
		}
	}
}
