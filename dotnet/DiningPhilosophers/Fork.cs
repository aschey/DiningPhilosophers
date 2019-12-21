using System;
using System.Collections.Generic;
using System.Text;
using System.Threading;

namespace DiningPhilosophers
{
    class Fork
    {
        public bool InUse { get; set; }

        private readonly object _lockObj = new object();

        public void Take()
        {
            lock (_lockObj)
            {
                if (InUse)
                {
                    throw new Exception("Taking fork that's in use");
                }
                InUse = true;
            }
        }

        public void Release()
        {
            InUse = false;
        }
    }
}
