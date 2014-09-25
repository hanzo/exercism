using System.Collections.Generic;

namespace Exercism.csharp.grade_school
{
	internal class School
	{
		private readonly Dictionary<int, List<string>> _roster;

		public Dictionary<int, List<string>> Roster
		{
			get { return new Dictionary<int, List<string>>(_roster); }
		}

		public School()
		{
			_roster = new Dictionary<int, List<string>>();
		}

		public void Add(string name, int grade)
		{
			if (_roster.ContainsKey(grade))
				_roster[grade].Add(name);
			else
				_roster[grade] = new List<string> {name};

			_roster[grade].Sort();
		}

		public List<string> Grade(int grade)
		{
			return _roster.ContainsKey(grade) ? _roster[grade] : new List<string>();
		}
	}
}
