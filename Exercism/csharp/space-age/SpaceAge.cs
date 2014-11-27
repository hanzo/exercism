namespace Exercism.csharp.space_age
{
	internal class SpaceAge
	{
		public SpaceAge(long seconds)
		{
			Seconds = seconds;
		}

		public double Seconds { get; private set; }

		public double OnEarth()
		{
			return Seconds/31557600;
		}

		public double OnMercury()
		{
			return Seconds/7600543.81992;  // 31557600 * 0.2408467 = 7600543.81992
		}

		public double OnVenus()
		{
			return Seconds/19414149.052176;
		}

		public double OnMars()
		{
			return Seconds/59354032.69008;
		}

		public double OnJupiter()
		{
			return Seconds/374355659.124;
		}

		public double OnSaturn()
		{
			return Seconds/929292362.8848;
		}

		public double OnUranus()
		{
			return Seconds/2651370019.3296;
		}

		public double OnNeptune()
		{
			return Seconds/5200418560.032;
		}
	}
}
