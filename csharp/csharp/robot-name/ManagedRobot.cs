using System;
using System.Collections.Generic;
using NUnit.Framework;

namespace Exercism.csharp.robot_name
{
	class ManagedRobot
	{
		public string Name { get; private set; }

		private static readonly Random _random = new Random();

		public ManagedRobot()
		{
			Name = NameRobot();
			RobotManager.UsedNames.Add(Name);
		}

		public void Reset()
		{
			Name = NameRobot();
			RobotManager.UsedNames.Add(Name);
		}

		// TODO: remove name from manager when robot is disposed

		private string NameRobot()
		{
			string name = GetRandomName();

			// TODO: prevent infinite loop that will occur when available names are exhausted
			while (RobotManager.UsedNames.Contains(name))
			{
				name = GetRandomName();
			}

			return name;
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

	static class RobotManager
	{
		public static List<string> UsedNames = new List<string>();
	}
}
