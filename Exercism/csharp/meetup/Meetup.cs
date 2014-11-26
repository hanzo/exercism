using System;

namespace Exercism.csharp.meetup
{
	class Meetup
	{
		private readonly int _month;
		private readonly int _year;

		public Meetup(int month, int year)
		{
			_month = month;
			_year = year;
		}

		public DateTime Day(DayOfWeek dayOfWeek, Schedule schedule)
		{
			int startDay = GetStartDay(schedule);

			var startDate = new DateTime(_year, _month, startDay);

			// Determine the number of days that need to be added to the start search date to reach the target day of the week.
			int startDayInt = (int) startDate.DayOfWeek;
			int dayInt = (int)dayOfWeek;
			if (dayInt < startDayInt)
				dayInt += 7;
			int dayOffset = dayInt - startDayInt;

			return startDate.AddDays(dayOffset);
		}

		// For the given schedule value (e.g. First), return the day of the month to begin searching.
		private int GetStartDay(Schedule schedule)
		{
			switch (schedule)
			{
				case Schedule.First:
					return 1;
				case Schedule.Second:
					return 8;
				case Schedule.Third:
					return 15;
				case Schedule.Fourth:
					return 22;
				case Schedule.Teenth:
					return 13;
				case Schedule.Last:
					// Search the last 7 days of the month
					int daysInMonth = DateTime.DaysInMonth(_year, _month);
					return daysInMonth - 6;
				default:
					return 1;
			}
		}
	}

	enum Schedule
	{
		First,
		Second,
		Third,
		Fourth,
		Last,
		Teenth
	}
}
