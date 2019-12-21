using System;
using System.Collections.Concurrent;
using System.Collections.Generic;
using System.Text;

namespace DiningPhilosophers
{
    public class ConcurrentHashSet<T>
    {
        private ConcurrentDictionary<T, object> _values = new ConcurrentDictionary<T, object>();

        public int Length => _values.Count;

        public void Add(T value)
        {
            _values.TryAdd(value, null);
        }

        public bool Contains(T value)
        {
            return _values.ContainsKey(value);
        }

        public void Remove(T value)
        {
            var outVal = new object();
            _values.Remove(value, out outVal);
        }
    }
}
