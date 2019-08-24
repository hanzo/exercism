using System;
using System.Linq;

namespace Exercism.csharp.phone_number
{
	internal class PhoneNumber
	{
		private readonly string _areaCode;
		private readonly string _subAreaCode;
		private readonly string _localNumber;

		public PhoneNumber(string number)
		{
			string sanitizedNumber = SanitizeInput(number);
			_areaCode = sanitizedNumber.Substring(0, 3);
			_subAreaCode = sanitizedNumber.Substring(3, 3);
			_localNumber = sanitizedNumber.Substring(6, 4);
		}

		private string SanitizeInput(string input)
		{
			string number = "0000000000";

			if (String.IsNullOrEmpty(input))
				return number;

			var numericChars = new string(input.Where(char.IsNumber).ToArray());

			if (numericChars.Length == 10)
				number = numericChars;

			if (numericChars.Length == 11 && numericChars[0] == '1')
				number = numericChars.Substring(1);

			return number;
		}

		public string Number
		{
			get { return _areaCode + _subAreaCode + _localNumber; }
		}

		public string AreaCode
		{
			get { return _areaCode; }
		}

		public override string ToString()
		{
			return String.Format("({0}) {1}-{2}", _areaCode, _subAreaCode, _localNumber);
		}
	}
}
