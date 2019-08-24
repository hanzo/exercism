using System;

namespace Exercism.csharp.triangle
{
	class Triangle
	{
		private readonly decimal _lenA;
		private readonly decimal _lenB;
		private readonly decimal _lenC;
		private readonly TriangleKind _kind;

		public Triangle(decimal lenA, decimal lenB, decimal lenC)
		{
			_lenA = lenA;
			_lenB = lenB;
			_lenC = lenC;
			_kind = ComputeKind();
		}

		public TriangleKind Kind()
		{
			return _kind;
		}

		private TriangleKind ComputeKind()
		{
			// Ensure that the given lengths represent a valid triangle. The triangle inequality states that 
			// the sum of the lengths of any two sides of a triangle always exceeds the length of the third side.
			if (   (_lenA + _lenB <= _lenC)
			    || (_lenA + _lenC <= _lenB)
			    || (_lenB + _lenC <= _lenA)
				)
			{
				throw new TriangleException();
			}	

			if (_lenA == _lenB && _lenA == _lenC)
				return TriangleKind.Equilateral;

			if (_lenA == _lenB || _lenA == _lenC || _lenB == _lenC)
				return TriangleKind.Isosceles;
				
			return TriangleKind.Scalene;
		}
	}

	internal enum TriangleKind
	{
		Equilateral,
		Isosceles,
		Scalene
	}

	internal class TriangleException : Exception
	{
		
	}
}
